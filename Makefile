
install:
	go install

get:
	luke request https://jsonplaceholder.typicode.com/posts/1

post:
	luke request -X POST https://jsonplaceholder.typicode.com/posts -d '{}'

debug:
	go build -gcflags "all=-N -l" -o luke-debug
	dlv exec ./luke-debug -- request -X POST "https://jsonplaceholder.typicode.com/posts" -d '{}' -s

