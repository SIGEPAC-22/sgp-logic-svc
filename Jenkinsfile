pipeline{
	agent any
	environment{
		name_final = "${env.JOB_NAME}"
	}
	stages{
		stage('Branch Check Out'){

            steps{
		    sh '''
		    echo ${name_final}
		    '''
            	}
		}
	}
}
