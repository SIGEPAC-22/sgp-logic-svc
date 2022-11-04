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
          sh 'echo DB DEV'
        }
      }
      stage('Deploy to DEV'){
        steps{
          sh 'echo Deploy DEV'
        }
      }
      stage('Cucumber Tests DEV'){
        steps{
          sh 'echo Cucumber Tests DEV'
        }
      }
      stage('RUN DB QA'){
        steps{
          sh 'echo DB QA'
        }
      }
      stage('Deploy to qa'){
        steps{
          sh 'echo Deploy QA'
        }
      }
      stage('QA Approval'){
        steps{
          sh 'echo Aprobacion QA'
        }
      }
      stage('RUN DB UAT'){
        steps{
          sh 'echo DB UAT'
        }
      }
      stage('Deploy to UAT'){
        steps{
          sh 'echo Deploy UAT'
        }
      }
      stage('Cucumber Tests UAT'){
        steps{
          sh 'echo Cucumber Tests UAT'
        }
      }
      stage('Wait to deploy in prd'){
        steps{
          sh 'echo Aprobacion PRD'
        }
      }
      stage('RUN DB PRD'){
        steps{
          sh 'echo Deploy DB'
        }
      }
      stage('Deploy to prd'){
        steps{
          sh 'echo Deploy PRD'
        }
      }
    }
}

