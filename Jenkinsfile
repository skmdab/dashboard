pipeline {
    agent any

    options {
        buildDiscarder(logRotator(artifactDaysToKeepStr: '', artifactNumToKeepStr: '', daysToKeepStr: '', numToKeepStr: '10'))
    }

    stages {
        stage('Checkout the code') {
            steps {
                git branch: 'testing', url: 'https://github.com/skmdab/dashboard.git'
            }
        }

        stage('Installing dependencies') {
            steps {
                script {
                    // Use npx to run npm commands with the correct Node.js version
                    sh '/usr/bin/npm install'
                }
            }
        }

        stage('Translation fix') {
            steps {
                script {
                    sh 'npm run fix:i18n'
                }
            }
        }

        stage('Building Dashboard') {
            steps {
                script {
                    sh 'make build'
                }
            }
        }

        stage('Building Docker image') {
            steps {
                script {
                    sh 'docker build --build-arg BUILDPLATFORM=linux/amd64 -f dist/amd64/Dockerfile -t mfzkubeos/kubeos-dashboard:latest dist/amd64/'
                }
            }
        }

        stage('Pushing image to docker hub') {
            steps {
                withCredentials([string(credentialsId: 'docker_creds', variable: 'docker_password')]) {
                    sh "docker login -u mfzkubeos -p ${docker_password}"
                    sh "docker push mfzkubeos/kubeos-dashboard"
                    }
                }
            }
        }
    }




