pipeline{
	agent any
	stages{
		stage('Branch Check Out'){
			when {
        			branch "fix-*"
      				}
            steps{
		    sh 'echo Funciona"'
            	}
		}
	}
}
