clean:
	rm -fr ./service-image/build

compile:
	go build -o ./service-image/build/service main.go 
	cp config.json ./service-image/build/