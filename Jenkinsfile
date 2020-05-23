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
    }
}