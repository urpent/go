REPORT_DIR ?= .

.PHONY: test
test:
	make staticcheck
	make gosec
	make unit-test
	go mod tidy


.PHONY: unit-test
unit-test:
	go install gotest.tools/gotestsum@latest
	gotestsum --jsonfile ${REPORT_DIR}/go_report.json -- -coverprofile=${REPORT_DIR}/go_coverage.out  -race -v ./...
	go tool cover -html go_coverage.out -o go_coverage.html

.PHONY: staticcheck
staticcheck:
	go fmt ./...
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...


.PHONY: gosec
gosec:
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	gosec -fmt=text ./...