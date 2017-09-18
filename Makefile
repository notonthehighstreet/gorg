NAME=gorg

test:
	cd pkg && go test -v -cover ./...
	cd cmd/gorgcli/command && go test -v -cover ./...

install:
	cd cmd/gorgcli && go build -o ${GOPATH}/bin/${NAME}
