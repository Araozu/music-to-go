pipeline {
agent {
	docker {
		image 'golang:1.23'
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
		stage('???') {
			steps {
				sh 'echo "TODO: call docker compose up"'
			}
		}
		stage('Profit') {
			steps {
				sh 'echo "done"'
			}
		}
	}
}
