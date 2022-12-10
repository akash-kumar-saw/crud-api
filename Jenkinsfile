pipeline {
    agent any

    tools {
        go 'go1.19'
    }

    // Setting up the Environment
    environment {
        GO119MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    
    // Setting up the Triggers
    triggers {
        pollSCM '* * * * *'
    }

    stages {

        // Installing all Dependencies    
        stage('Unit Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }
        }

        // Compiling and Building the Go Application and Docker Image
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build'
                sh 'docker build . -t akashkumarsaw/crud-api'
            }
        }

        // Testing the Go Application
        stage('Functional Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'go vet .'
                    echo 'Running linting'
                    sh 'golint .'
                    echo 'Running test'
                    sh 'cd test && go test -v'
                }
            }
        }

        // Delivering the Docker Image to DockerHub
        stage('Deliver') {
            agent any
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
                sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
                sh 'docker push akashkumarsaw/crud-api'
            }
        }
    } 
}
