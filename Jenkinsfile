pipeline{
    agent any
    stages{
        stage("Env Build Number"){
            steps{
                echo "${env.BRANCH_NAME}"
		echo "${env.BRANCH_IS_PRIMARY}"
		echo "${env.TAG_NAME}"
		echo "${env.BUILD_DISPLAY_NAME}"
		echo "${env.JOB_NAME}"
		echo "${env.JOB_BASE_NAME}"
		echo "${env.BUILD_TAG}"
		echo "${env.NODE_NAME}"
		echo "${env.NODE_LABELS}"
		echo "${env.WORKSPACE}"
		echo "${env.GIT_BRANCH}"
		echo "${env.GIT_LOCAL_BRANCH}"
		echo "${env.GIT_CHECKOUT_DIR}"
		echo "${env.GIT_AUTHOR_NAME}"
	    }
        }
    }
}
