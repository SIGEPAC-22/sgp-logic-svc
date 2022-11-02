pipeline{
	agent any
	stages{
		stage('Branch Check Out'){

            steps{
		    sh '''
		    ${env.BRANCH_NAME}
		    env.GIT_BRANCH
		    '''
            	}
		}
	}
}
