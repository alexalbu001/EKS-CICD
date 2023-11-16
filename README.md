EKS CI/CD Project
Overview
This repository contains the infrastructure and pipeline configuration for a Kubernetes-based CI/CD workflow, deployed on AWS EKS. It showcases the deployment of a sample application using Jenkins for continuous integration and delivery, with monitoring set up via Prometheus and Grafana.

Repository Structure
k8s/: Kubernetes configuration files for setting up the application deployment, services, and monitoring tools (Prometheus and Grafana).
Jenkinsfile: Defines the Jenkins pipeline for automated build, test, and deployment of the application to the Kubernetes cluster.
Dockerfile: Configuration for building the Docker container of the application.




Getting Started

Prerequisites

AWS EKS cluster
Jenkins set up with access to the EKS cluster
Docker for building container images


Setting up the Kubernetes Environment

Navigate to the k8s/ directory.
Apply the Kubernetes configurations to set up the application and monitoring tools:
kubectl apply -f deployment.yaml


Configuring Jenkins Pipeline

Set up a Jenkins job using the Jenkinsfile in the root directory.
Configure Jenkins to trigger a build on a new push to the GitHub repository.
Building and Deploying the Application
The Dockerfile in the root directory can be used to build the application container.
The Jenkins pipeline automates the deployment of this container to the Kubernetes cluster.


Monitoring with Prometheus and Grafana

Prometheus and Grafana are configured via the Kubernetes configurations in the k8s/ directory.
Access the Grafana dashboard and Prometheus UI via the NodePort services set up in the Kubernetes cluster.