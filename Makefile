#APPBIN = expenses-backend
#BACKEND_DIR = $(GOPATH)/src/github.com/snichme/$(APPBIN)

BACKEND_DIR = backend
FRONTEND_DIR = frontend
RELEASE_DIR = release

all: run

build-backend:
	$(MAKE) -C $(BACKEND_DIR) build
	mkdir -p $(RELEASE_DIR)
	cp $(BACKEND_DIR)/main ./release/

run-backend:
	$(MAKE) -C $(BACKEND_DIR) run

build-frontend:
	$(MAKE) -C $(FRONTEND_DIR) build
	mkdir -p $(RELEASE_DIR)
	cp -rf $(FRONTEND_DIR)/resources/public $(RELEASE_DIR)/public
	cp $(FRONTEND_DIR)/resources/index.html $(RELEASE_DIR)/public/

build: build-frontend build-backend

run: build
	./$(RELEASE_DIR)/main

build-docker:
	docker build --rm --tag=johannesboyne/godockersample .

run-docker:
	docker run \
		-e HELLO=world \
		-p 1337:1337 \
		johannesboyne/godockersample

deploy: build run-docker

clean:
	rm -rf $(RELEASE_DIR)
