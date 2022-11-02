pipeline{
	agent any
	environment{
        	branch_name = "${env.BRANCH_NAME}"
	}
	stages{
		stage('Branch "${branch_name}" Check Out'){
            steps{
                sh 'echo "${env.BRANCH_NAME}"'
            	}
		}
	}
}
