pipeline {
  agent none
  environment {
    name_b = "${env.BRANCH_NAME}"
  }
  stages {
    stage('Branch Check Out') {
      agent{
        label 'boot'
      }
      steps {
        script {
          if (name_b == "master") {
            sh '''
            echo ${name_b}
            '''
          }
        }
      }
    }
  }
}
