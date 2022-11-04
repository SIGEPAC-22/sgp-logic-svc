pipeline{
    agent none
    enviroment{
        namebranch = "${env.BRANCH_NAME}"
    }
    stages{
      stage('Docker Build'){
        agent{
          label('dev && qa')
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
            if (name_b == "master") {
              sh '''
              echo ${name_b}
              '''
            }
          }
        }
      }
      stage('SonarQube Analysis'){

      }
      stage('RUN DB DEV'){

      }
      stage('Deploy to DEV'){

      }
      stage('Cucumber Tests'){

      }
      stage('RUN DB QA'){

      }
      stage('Deploy to qa'){

      }
      stage('QA Approval'){

      }
      stage('RUN DB UAT'){

      }
      stage('Deploy to UAT'){

      }
      stage('Cucumber Tests'){

      }
      stage('Wait to deploy in prd'){

      }
      stage('RUN DB PRD'){

      }
      stage('Deploy to prd'){

      }
    }
}

