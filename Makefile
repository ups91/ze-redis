
build:
	go build -o ./build/ze-redis ./cmd/

docker:
	rm -rf ./build/* ;\
	mkdir -p  ./build/linux ;\
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./build/linux/ze-redis.so ./cmd/ ;\
	#chown 65534:65534 ./build/linux/ze-redis.so ;\
	#chmod 666 ./build/linux/ze-redis.so ;\
	docker build -f ./Dockerfile -t ups91/ze-redis-test:latest --no-cache --force-rm ./build/linux/ ;\
	docker push  ups91/ze-redis-test:latest
	