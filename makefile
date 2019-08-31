GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# If either of these change, update the Dockerfile
BINARY_NAME=kubedec
TARGET_DIR=target

COVERAGE_FILE=c.out

GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

BUILD_CMD=$(GOBUILD) -o $(TARGET_DIR)/$(BINARY_NAME) -v
TEST_CMD=$(GOTEST) -v ./...

ZIP_ARCHIVE=current.zip

linuxBuild: clean deps
	GOOS=linux GOARCH=amd64 $(BUILD_CMD)
build:
	$(BUILD_CMD)
dist:
	zip -j $(TARGET_DIR)/$(ZIP_ARCHIVE) $(TARGET_DIR)/$(BINARY_NAME) resources/docker/default-play.docker resources/builds/build_app.sh resources/sonar/* resources/k8s/*
test:
	$(TEST_CMD)
coverage:
	$(TEST_CMD) -coverprofile=$(COVERAGE_FILE)
	$(GOCMD) tool cover -html=$(COVERAGE_FILE)
clean:
	$(GOCLEAN)
	rm -rf $(TARGET_DIR)
	rm -f $(COVERAGE_FILE)
run: build
	./$(TARGET_DIR)/$(BINARY_NAME)
