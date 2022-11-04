pipeline {
  agent any
  environment {
    name_b = "${env.BRANCH_NAME}"
  }
  stages {
    stage('Branch Check Out') {
      steps {
        script {
          if (name_b == "master") {
            agent {
              label 'windows'
            }
            sh '''
            echo ${name_b}
            '''
          }
        }
      }
    }
  }
}
