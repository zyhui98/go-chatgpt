export tag=v1.0
root:
	export ROOT=github.com/cncamp/golang

build:
	echo "building go-chatgpt binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building go-chatgpt container"
	docker build -t zyhui98/go-chatgpt:${tag} .

push: release
	echo "pushing zyhui98/go-chatgpt"
	docker push zyhui98/go-chatgpt:${tag}
