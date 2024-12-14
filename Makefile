clean:
	rm -fr ./service-image/build

compile:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./service-image/build/service main.go
	cp config.json ./service-image/build/