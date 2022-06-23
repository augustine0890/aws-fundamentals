.PHONY: clean build package

clean:
	rm -rf ./hello package.zip

build:
	go get github.com/aws/aws-lambda-go/lambda
	GOOS=linux GOARCH=amd64 go build -o hello .

package:
	zip -r package.zip hello