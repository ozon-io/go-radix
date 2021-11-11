test:
	go test ./... -coverprofile cover.prof && go tool cover -func cover.prof

test-static:
	staticcheck ./... | staticcheck_color.sh ; exit $${PIPESTATUS[0]}

test-cover:
	go test ./... -coverprofile cover.prof && go tool cover -html=cover.prof -o cover.html
