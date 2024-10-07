pipeline {
	agent any
	stages {
		stage('Install') {
			agent {
				docker {
					image 'golang:1.23-alpine'
					reuseNode true
				}
			}
			steps {
				sh 'go mod tidy'
				sh 'go build main.go'
			}
		}
		stage('Populate env') {
			steps {
				sh 'rm .env || true'
				sh 'echo "PORT=8007" > .env'
			}
		}
		stage('Profit') {
			steps {
				dir('docker') {
					sh 'docker-compose stop || true'
					sh 'docker-compose up -d --build'
				}
			}
		}
	}
}
