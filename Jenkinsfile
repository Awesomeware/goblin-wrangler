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

        stage('Node') {
            steps {
                container('node') {
                    sh 'cd entertainer && yarn install && yarn build'
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
                container('docker') {
                    sh "sed -i s#GOBLINWRANGLER_NAMESPACE#${env.BRANCH_NAME.toLowerCase()}#g artificer/goblin-wrangler/namespace.yaml"
                    sh "sed -i s#GOBLINWRANGLER_HOST#${env.BRANCH_NAME == 'main' ? '' : env.BRANCH_NAME.toLowerCase() + '.'}goblinwrangler.com#g artificer/goblin-wrangler/ingresses.yaml"
                    sh "sed -i s#KINGPIN_IMAGE#registry.digitalocean.com/goblin-wrangler/${env.BRANCH_NAME.toLowerCase()}/kingpin:latest#g artificer/goblin-wrangler/deployments.yaml"
                    sh "kubectl apply -f artificer/goblin-wrangler/namespace.yaml"
                    sh "kubectl apply -k artificer/goblin-wrangler/ -n ${env.BRANCH_NAME.toLowerCase()}"
                    sh "mv entertainer/build entertainer/html && kubectl cp entertainer/html ${env.BRANCH_NAME.toLowerCase()}/\$(kubectl get pod -o name -l app=entertainer -n ${env.BRANCH_NAME.toLowerCase()} | cut -d '/' -f 2):/usr/share/nginx && mv entertainer/html entertainer/build"
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
