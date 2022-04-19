node(label: "worker-hotel") {
     checkout scm
     // Install the desired Go version
     def root = tool name: 'golang', type: 'go'
     def envs = ""
     def key = ""
     def port
     // Export environment variables pointing to the directory where Go was installed
     withEnv(["GOROOT=${root}",
             "PATH+GO=${root}/bin",
             "ENV_FILE_PATH=${env.WORKSPACE}/.env",
             "GOPATH=/var/lib/jenkins/go",
             "GOENV=devci",
             "LOGS_PATH=${env.WORKSPACE}/logs",
             "WORKDIR=${env.WORKSPACE}",
             "GOAPP=TIX-HOTEL-AUTOCOMPLETE-BE"]) {
           stage ('Prepare code') {
               sh "mkdir -p ${env.WORKSPACE}/logs"
               sh "mkdir -p ${env.WORKSPACE}/sonarqube-report"
               sh "rm -rf ${env.WORKSPACE}/sonarqube-report/*"
           }
           stage('Unit Test') {
               def scannerHome = tool 'sonar-scanner';
               withSonarQubeEnv('Sonar-Tiket') {
                   sh "${envs} go mod download"
                   sh "${envs} go clean -testcache"
                   sh "${envs} go test ./... -failfast -cover -coverprofile=sonarqube-report/coverage.out -json > sonarqube-report/coverage.json && /usr/bin/git checkout ."
                   withCredentials([usernamePassword(credentialsId: 'jenkinsci-sonar-scanner', passwordVariable: 'pass', usernameVariable: 'user')]) {
                     sh "${scannerHome}/bin/sonar-scanner -Dsonar.login=${user} -Dsonar.password=${pass} -Dproject.settings=sonar.properties"
                   }
               }
           }
           stage("Quality Gate") {
               timeout(time: 1, unit: 'MINUTES') { // Just in case something goes wrong, pipeline will be killed after a timeout
                   sleep(20)
                   def qg = waitForQualityGate()
                   if (qg.status != 'OK') {
                       error "Pipeline aborted due to quality gate failure: ${qg.status}"
                   }
                   sh "docker-compose down --remove-orphans"
               }
           }
           }
 }