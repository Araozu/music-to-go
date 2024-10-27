pipeline {
	agent any
	stages {
		stage('Generate templ files') {
			agent {
				docker {
					image 'ghcr.io/a-h/templ:latest'
					reuseNode true
				}
			}
			steps {
				sh 'go mod tidy'
				sh 'templ generate'
				sh 'go build main.go'
			}
		}
		stage('Build go binary') {
			agent {
				docker {
					image 'golang:latest'
					reuseNode true
				}
			}
			steps {
				sh 'go mod tidy'
				sh 'go build main.go'
			}
		}
		stage('Build static assets') {
			agent {
				docker {
					image 'node:22'
					reuseNode true
				}
			}
			steps {
				sh 'npm i -g pnpm'
				sh 'pnpm i'
				sh 'pnpm build'
			}
		}
		stage('Populate env') {
			steps {
				sh 'rm .env || true'
				sh 'echo "PORT=8007" >> .env'
				sh 'echo "DB_HOST=music-to-go-db" >> .env'
				sh 'echo "DB_PORT=5432" >> .env'
				sh 'echo "DB_DATABASE=music" >> .env'
				sh 'echo "DB_USERNAME=root" >> .env'
				sh 'echo "DB_PASSWORD=root" >> .env'
				sh 'echo "DB_SCHEMA=public" >> .env'
			}
		}
		stage('Profit') {
			steps {
				dir('docker') {
					sh 'docker-compose down || true'
					sh 'docker-compose up -d --build'
				}
			}
		}
	}
}
