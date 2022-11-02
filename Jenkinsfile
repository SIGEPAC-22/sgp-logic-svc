pipeline{
	agent any
	stages{
		stage('Branch Check Out'){

            steps{
		    sh '''
		    echo env.BRANCH_NAME
		    echo env.GIT_BRANCH
		    '''
            	}
		}
	}
}
