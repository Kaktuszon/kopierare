APP_NAME = kopierare

build:
	go build -o bin/${APP_NAME}

run: build
	./bin/${APP_NAME}

archive: build
	tar -cvf bin/${APP_NAME}.tar bin/${APP_NAME}