test:
	go test ./...

test-static:
	staticcheck ./... | staticcheck_color.sh ; exit $${PIPESTATUS[0]}

