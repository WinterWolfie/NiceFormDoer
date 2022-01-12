build:
	env GOOS=linux GOARCH=amd64
	go build -o binary/FormDoerLinux

	env GOOS=windows GOARCH=amd64
	go build -o binary/FormDoerWindows