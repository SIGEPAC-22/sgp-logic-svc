pipeline{
	agent any
	environment{
		name_b = "${env.BRANCH_NAME}"
	}
	stages{
		stage('Branch Check Out'){

            steps{
		    script{
			    if (name_b == "master"){
			    			    sh '''
		    echo ${name_b}
		    '''
			    }
		    }

            	}
		}
	}
}
