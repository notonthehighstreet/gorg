NAME=gorg

install:
	cd cmd/gorgcli && go build -o ${GOPATH}/bin/${NAME}
