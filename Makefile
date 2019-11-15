
mocks:
	moq -out ./task/repo_mocks.go ./task Service


cover:
	go test -coverprofile cover.out ./...
	uncover cover.out
	rm cover.out


test:
	go test -v ./...