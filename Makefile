.PHONY : installer compile install

installer:
	chmod +x ./installer/install.sh
	./installer/install.sh

compile:
	go build -o cemq
	chmod +x ./cemq

install: compile installer

test:
	./cemq help
