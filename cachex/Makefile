REPORT_DIR ?= ..

.PHONY: test
test:
	make staticcheck
	make gosec
	make unit-test
	go mod tidy


.PHONY: unit-test
unit-test:
	go install gotest.tools/gotestsum@latest
	gotestsum --jsonfile ${REPORT_DIR}/cachex_report.json -- -coverprofile=${REPORT_DIR}/cachex_coverage.out -race -v ./...


.PHONY: staticcheck
staticcheck:
	go fmt ./...
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...


.PHONY: gosec
gosec:
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	gosec -fmt=text ./...


.PHONY: mock
mock:
	go install github.com/vektra/mockery/v2@latest
	mockery --name=Cacher --filename=cacher_mock.go --structname=CacherMock --inpackage