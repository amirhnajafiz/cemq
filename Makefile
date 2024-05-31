install:
	go build -o cemq
	chmod +x ./cemq

test:
	./cemq help
