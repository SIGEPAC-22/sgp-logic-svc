pipeline{
    agent any
    stages{
        stage("Env Build Number"){
            steps{
                echo "${env.BRANCH_NAME}"
		echo "${env.BRANCH_IS_PRIMARY}"
	    }
        }
    }
}
