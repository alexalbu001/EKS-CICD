pipeline {
    agent {
  label 'jenkins/docker-builder'
    }
    environment {
        DOCKER_IMAGE = "zoeycide/flask-hello-world"
        DOCKER_TAG = "latest"
    }
    stages {
        stage('Checkout Code') {
            steps {
                checkout scm
            }
        }
        stage('Build & Push Docker Image') {
            steps {
                script {
                    docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}")
                    docker.withRegistry('https://registry.hub.docker.com', 'dockerhub-credentials') {
                        docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
                    }
                }
            }
        }
        stage('Deploy to Kubernetes') {
            steps {
                script {
                    withCredentials([file(credentialsId: 'kubeconfig', variable: 'KUBECONFIG')]) {
                        // KUBECONFIG environment variable points to the loaded kubeconfig file
                        sh "kubectl set image deployment/flask-hello-world-deployment flask-hello-world=${DOCKER_IMAGE}:${DOCKER_TAG} --record"
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
