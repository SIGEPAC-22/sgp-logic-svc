pipeline{
    agent none
    environment{
        namebranch = "${env.BRANCH_NAME}"
        name_final = "${env.JOB_NAME}"
        DB_CREDS=credentials('db-creds')
    }
    stages{
      stage('Docker Build'){
        agent{
            label 'dev'
          }
        when{
          anyOf{
            branch 'sgp*'
            branch 'sprint-*'
            branch 'master'
          }
        }
        steps{
				  script{
					    def result = sh(returnStdout: true, script: 'echo "$(docker ps -q --filter name=${name_final})"').trim()
					    if (result != ""){
						    sh '''
						    docker stop ${name_final}
						    docker rm -vf ${name_final}
						    docker build . -t ${name_final}
						    docker system prune -f
						    '''
					    }else{
						    sh '''
						    docker build . -t ${name_final}
                docker system prune -f
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
				  script{
					  sh '''
		    			docker run --rm flyway/flyway:8.5.1 version
		    			docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:8.5.1 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW migrate
		    			docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:8.5.1 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW validate
		    			docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:8.5.1 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW info
		    			'''
				  }
			  }
      }
      stage('Deploy to DEV'){
        steps{
				  script{
					    sh '''
		    			docker run  -dt -p :90 --name ${name_final} ${name_final}
		    			docker system prune -f
              '''
				  }
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
