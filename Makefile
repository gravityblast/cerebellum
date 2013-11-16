ENV=TRAFFIC_CONFIG_FILE=traffic.conf.sample
GO_CMD=go
GO_TEST= $(ENV) TRAFFIC_ENV=test $(GO_CMD) test ./...
GO_BUILD=$(GO_CMD) build -v
RUN=$(ENV) ./cerebellum

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
