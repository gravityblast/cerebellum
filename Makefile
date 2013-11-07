DB="user=musicbrainz dbname=musicbrainz host=127.0.0.1 sslmode=disable"

GO_CMD=go
GO_TEST=TRAFFIC_ENV=test DB=$(DB) $(GO_CMD) test ./...
GO_BUILD=$(GO_CMD) build -v
RUN=DB=$(DB) ./cerebellum

all: build
test: RunTests
build: BuildApp
run: RunApp

BuildApp:RunTests
	$(GO_BUILD)

RunTests:
	$(GO_TEST)

RunApp:BuildApp
	$(RUN)
