pipeline{
	agent any
	stages{
		stage('Branch Check Out'){

            steps{
		    sh '''
    git rev-parse --abbrev-ref HEAD > GIT_BRANCH'
    git_branch = readFile('GIT_BRANCH').trim()
    echo git_branch
   '''
            	}
		}
	}
}
