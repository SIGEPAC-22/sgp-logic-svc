pipeline{
	agent any
	environment{
		name_final = "${env.JOB_NAME}"
		name_b = "${env.BRANCH_NAME}"
	}
	stages{
		stage('Branch Check Out'){

            steps{
		    sh '''
		    echo ${name_final}
		    echo ${name_b}
		    '''
            	}
		}
	}
}
