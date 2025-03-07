

completion:
	go install
	luke completion bash | sudo tee /etc/bash_completion.d/luke > /dev/null
# source /etc/bash_completion.d/luke

name:
	find cmd -name "*.go" -print | xargs sed -i 's/RootCmd/rootCmd/g'