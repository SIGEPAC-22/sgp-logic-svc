pipeline{
    agent none
    environment{
        namebranch = "${env.BRANCH_NAME}"
    }
    stages{
      stage('Docker Build'){
        agent{
            label 'dev'
            label 'qa'
          }
        }
        when{
          anyOf{
            branch 'sgp'
            branch 'sprint-*'
            branch 'master'
          }
        }
        steps{
          script{
            if (namebranch == "master") {
              sh '''
              echo ${namebranch}
              '''
            }
          }
        }
      }
      stage('SonarQube Analysis'){
        steps{
          sh 'echo SonarQube'
        }
      }
      stage('RUN DB DEV'){
        steps{
          sh 'echo SonarQube'
        }
      }
      stage('Deploy to DEV'){
        steps{
          sh 'echo SonarQube'
        }
      }
      stage('Cucumber Tests DEV'){
        steps{
          sh 'echo SonarQube'
        }
      }
      stage('RUN DB QA'){
        steps{
          sh 'echo SonarQube'
        }
      }
      stage('Deploy to qa'){
        steps{
          sh 'echo SonarQube'
        }
      }
      stage('QA Approval'){
        steps{
          sh 'echo SonarQube'
        }
      }
      stage('RUN DB UAT'){
        steps{
          sh 'echo SonarQube'
        }
      }
      stage('Deploy to UAT'){
        steps{
          sh 'echo SonarQube'
        }
      }
      stage('Cucumber Tests UAT'){
        steps{
          sh 'echo SonarQube'
        }
      }
      stage('Wait to deploy in prd'){
        steps{
          sh 'echo SonarQube'
        }
      }
      stage('RUN DB PRD'){
        steps{
          sh 'echo SonarQube'
        }
      }
      stage('Deploy to prd'){
        steps{
          sh 'echo SonarQube'
        }
      }
    }
}

