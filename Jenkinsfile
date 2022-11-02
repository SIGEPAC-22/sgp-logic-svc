pipeline{
	agent any
	environment{
        	branch_name = "${env.GIT_BRANCH}"
		name_final = "${env.JOB_NAME}"
	}
	stages{
		stage('Branch "${branch_name}" Check Out'){
            steps{
		    sh 'echo "${branch_name}"'
            	}
		}
	}
}
