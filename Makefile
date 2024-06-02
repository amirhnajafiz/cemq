.PHONY : installer compile install test

installer:
	chmod +x ./installer/install.sh
	./installer/install.sh

compile:
	chmod +x ./installer/dependencies.sh
	./installer/dependencies.sh || { echo "Could not compile CEMQ."; exit 1; }
	go build -o cemq
	chmod +x ./cemq

install: compile installer

test:
	cemq help
