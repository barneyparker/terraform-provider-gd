.PHONY: setup clean build install package
	
setup:
	go get ./...
	
clean:
	rm terraform-provider-gd

build:
	go build -o terraform-provider-gd

install:
	cp terraform-provider-gd ~/.terraform/plugins/terraform-provider-gd

package:
	./package.sh darwin amd64
	./package.sh freebsd 386
	./package.sh freebsd amd64
	./package.sh freebsd arm
	./package.sh linux 386
	./package.sh linux amd64
	./package.sh linux arm
	./package.sh openbsd 386
	./package.sh openbsd amd64
	./package.sh openbsd arm
	./package.sh solaris amd64
	./package.sh windows 386
	./package.sh windows amd64
	# Should automatically upload here
	# Then rm *.tgz