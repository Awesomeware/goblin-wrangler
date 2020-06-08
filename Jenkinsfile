pipeline {
    agent {
        kubernetes {
            defaultContainer 'jnlp'
            yamlFile 'agent.yaml'
        }
    }

    stages {
        stage('Build and Push') {
            steps {
                container('docker') {
                    sh "docker build -t registry.goblinwrangler.com/${env.BRANCH_NAME.toLowerCase()}/kingpin:latest kingpin/ && docker push registry.goblinwrangler.com/${env.BRANCH_NAME.toLowerCase()}/kingpin:latest"
                }
            }
        }

        stage('Deploy') {
            steps {
                container('busybox') {
                    sh "sed -i s#KINGPIN_NAMESPACE#${env.BRANCH_NAME.toLowerCase()}#g artificer/kingpin/namespace.yaml"
                    sh "sed -i s#KINGPIN_HOST#${env.BRANCH_NAME.toLowerCase()}.goblinwrangler.com#g artificer/kingpin/kingpin.yaml"
                    sh "sed -i s#KINGPIN_IMAGE#registry.goblinwrangler.com/${env.BRANCH_NAME.toLowerCase()}/kingpin:latest#g artificer/kingpin/kingpin.yaml"
                }
                container('kubectl') {
                    sh "kubectl apply -f artificer/kingpin/namespace.yaml"
                    sh "kubectl apply -k artificer/kingpin/ -n ${env.BRANCH_NAME.toLowerCase()}"
                }
            }

            post {
                success {
                    script {
                        if (env.CHANGE_ID) {
                            pullRequest.createStatus(status: 'success', context: 'continuous-integration/jenkins/deployment', description: 'Pull request is deployed', targetUrl: "https://${env.BRANCH_NAME.toLowerCase()}.goblinwrangler.com")
                        }
                    }
                }
            }
        }
    }
}