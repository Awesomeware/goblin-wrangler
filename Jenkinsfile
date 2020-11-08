pipeline {
    agent {
        kubernetes {
            defaultContainer 'jnlp'
            yamlFile 'agent.yaml'
        }
    }

    stages {
        stage('Python') {
            steps {
                container('python') {
                    sh 'curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python'
                    sh 'cd kingpin && $HOME/.poetry/bin/poetry install'
                    sh 'cd kingpin && $HOME/.poetry/bin/poetry run pylint --output-format=parseable kingpin > pylint.report || true'
                    sh 'cd kingpin && $HOME/.poetry/bin/poetry run pycodestyle kingpin > pep8.report || true'
                    sh 'cd kingpin && $HOME/.poetry/bin/poetry run pytest --junitxml=junit.xml --cov-report xml:coverage.xml --cov=kingpin'  
                }
            }

            post {
                always {
                    recordIssues tools: [pyLint(pattern: "kingpin/pylint.report"), pep8(pattern: "kingpin/pep8.report")], healthy: 1
                    junit testResults: 'kingpin/junit.xml'
                    cobertura coberturaReportFile: 'kingpin/coverage.xml'
                }
            }
        }

        stage('Docker') {
            steps {
                container('docker') {
                    sh "docker build -t registry.digitalocean.com/goblin-wrangler/${env.BRANCH_NAME.toLowerCase()}/kingpin:latest kingpin/ && docker push registry.digitalocean.com/goblin-wrangler/${env.BRANCH_NAME.toLowerCase()}/kingpin:latest"
                }
            }
        }

        stage('Deploy') {
            steps {
                container('busybox') {
                    sh "sed -i s#KINGPIN_NAMESPACE#${env.BRANCH_NAME.toLowerCase()}#g artificer/kingpin/namespace.yaml"
                    sh "sed -i s#KINGPIN_HOST#${env.BRANCH_NAME == 'main' ? '' : env.BRANCH_NAME.toLowerCase() + '.'}goblinwrangler.com#g artificer/kingpin/kingpin.yaml"
                    sh "sed -i s#KINGPIN_IMAGE#registry.digitalocean.com/goblin-wrangler/${env.BRANCH_NAME.toLowerCase()}/kingpin:latest#g artificer/kingpin/kingpin.yaml"
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
