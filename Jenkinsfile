pipeline{
	agent any
	stages{
		stage('Branch Check Out'){

            steps{
		    sh '''
		    echo ${env.JOB_NAME}
		    echo env.BRANCH_NAME
		    echo env.GIT_BRANCH
		    '''
            	}
		}
	}
}
