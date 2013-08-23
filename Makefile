DB ="user=musicbrainz dbname=musicbrainz host=127.0.0.1 sslmode=disable"

GO_CMD=go
GO_TEST=DB=$(DB) $(GO_CMD) test
GO_BUILD=$(GO_CMD) build -v

all: build
test: RunTests
build: BuildApp

BuildApp:RunTests
	$(GO_BUILD)

RunTests:
	$(GO_TEST)
