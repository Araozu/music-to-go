pipeline {
agent {
	docker {
		image 'golang:1.23-alpine'
		}
	}
	stages {
		stage('Install deps') {
			steps {
				sh 'go mod tidy'
			}
		}
		stage('Build binary') {
			steps {
				sh 'go build main.go'
			}
		}
		stage('Profit') {
			steps {
				dir('docker') {
					sh 'docker compose stop || true'
					sh 'docker compose up --build -d'
				}
			}
		}
	}
}
