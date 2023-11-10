pipeline {
    agent none // No default agent since we define our own
    environment {
        DOCKER_IMAGE = "zoeycide/flask-hello-world"
        DOCKER_TAG = "latest"
    }
    stages {
        stage('Build & Push Docker Image') {
            steps {
                script {
                    // Define the pod with a single container for Docker operations
                    podTemplate(
                        name: 'docker-build-pod',
                        label: 'docker-build-pod',
                        containers: [
                            containerTemplate(name: 'docker', image: 'docker', command: 'cat', ttyEnabled: true)
                        ],
                        volumes: [
                            hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock')
                        ]
                    ) {
                        node('docker-build-pod') {
                            container('docker') {
                                checkout scm
                                sh "docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} ."
                                sh "docker login -u $DOCKER_USER -p $DOCKER_PASS"
                                sh "docker push ${DOCKER_IMAGE}:${DOCKER_TAG}"
                            }
                        }
                    }
                }
            }
        }
        stage('Deploy to Kubernetes') {
            steps {
                script {
                    withCredentials([file(credentialsId: 'kubeconfig', variable: 'KUBECONFIG')]) {
                        sh "kubectl set image deployment/flask-hello-world-deployment flask-hello-world-container=${DOCKER_IMAGE}:${DOCKER_TAG} --record"
                    }
                }
            }
        }
    }
    post {
        success {
            echo 'Build succeeded!'
        }
        failure {
            echo 'Build failed!'
        }
    }
}
