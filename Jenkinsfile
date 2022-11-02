pipeline{
	agent any
	environment{
        	branch_name = "${env.BRANCH_NAME}"
		name_final = "${env.JOB_NAME}"
	}
	stages{
		stage('Branch "${branch_name}" Check Out'){
            steps{
		    sh 'echo "${name_final}"'
            	}
		}
	}
}
