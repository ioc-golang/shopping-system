tidy:
	go mod tidy -compat=1.17

build-all-binary: tidy
	export GOPROXY="https://goproxy.cn"
	GOARCH=amd64 GOOS=linux go build -o ./cmd/advertisement/app ./cmd/advertisement
	GOARCH=amd64 GOOS=linux go build -o ./cmd/festival/app ./cmd/festival
	GOARCH=amd64 GOOS=linux go build -o ./cmd/product/app ./cmd/product
	GOARCH=amd64 GOOS=linux go build -o ./cmd/shopping-ui/app ./cmd/shopping-ui

clear-all-binary:
	rm ./cmd/advertisement/app
	rm ./cmd/festival/app
	rm ./cmd/product/app
	rm ./cmd/shopping-ui/app


build-all-images: build-all-binary
	docker buildx build \
		--build-arg appname=festival \
		--platform linux/amd64 \
		-t laurencelizhixin/ioc-shopping-system-festival:latest \
		. --push
	docker buildx build \
		--build-arg appname=advertisement \
		--platform linux/amd64 \
		-t laurencelizhixin/ioc-shopping-system-advertisement:latest \
		. --push
	docker buildx build \
		--build-arg appname=product \
		--platform linux/amd64 \
		-t laurencelizhixin/ioc-shopping-system-product:latest \
		. --push
	docker buildx build \
		--build-arg appname=shopping-ui \
		--platform linux/amd64 \
		-t laurencelizhixin/ioc-shopping-system-shopping-ui:latest \
		. --push
	make clear-all-binary


deploy-to-k8s:
	kubectl apply -f ./deploy/k8s/namespace
	kubectl apply -f ./deploy/k8s/shopping-system

remove-from-k8s:
	kubectl delete -f ./deploy/k8s/shopping-system
	kubectl delete -f ./deploy/k8s/namespace

deploy-with-docker:
	docker-compose -f ./deploy/docker-compose/docker-compose.yaml up -d

remove-from-docker:
	docker-compose -f ./deploy/docker-compose/docker-compose.yaml down

iocli-update:
	iocli gen