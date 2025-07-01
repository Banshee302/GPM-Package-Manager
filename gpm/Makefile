# === GPM Makefile with Feature Selection Support ===

# Config
FEATURES_FILE := gpm.features
GPM_BINARY := gpm
AUTOBUILD_BINARY := autobuild/autobuild
GPM_API := gpmapi/src/main.go
LOG_FILE := build.log

# Read feature values
USE_SSH := $(shell grep -i '^USE_SSH=' $(FEATURES_FILE) | cut -d '=' -f2)
USE_HTTPS := $(shell grep -i '^USE_HTTPS=' $(FEATURES_FILE) | cut -d '=' -f2)
ENABLE_AUTOBUILD := $(shell grep -i '^ENABLE_AUTOBUILD=' $(FEATURES_FILE) | cut -d '=' -f2)
ENABLE_API := $(shell grep -i '^ENABLE_API=' $(FEATURES_FILE) | cut -d '=' -f2)

# Sanity check
ifeq ("$(wildcard $(FEATURES_FILE))","")
$(error Missing gpm.features. Run ./menuconfig.sh first.)
endif

# Default build target
all: clean deps build install

# Dependency setup
deps:
	@echo "[1] Checking dependencies..." | tee -a $(LOG_FILE)
	go mod tidy | tee -a $(LOG_FILE)

# Build targets with conditional features
build:
	@echo "[2] Building GPM core..." | tee -a $(LOG_FILE)
	go build -o $(GPM_BINARY) | tee -a $(LOG_FILE)

ifneq ($(ENABLE_AUTOBUILD),false)
	@echo "[3] Building Autobuild..." | tee -a $(LOG_FILE)
	go build -o $(AUTOBUILD_BINARY) autobuild/autobuild.go | tee -a $(LOG_FILE)
endif

ifneq ($(ENABLE_API),false)
	@echo "[4] Building GPM-API..." | tee -a $(LOG_FILE)
	go build -o GpmAPI $(GPM_API) | tee -a $(LOG_FILE)
endif

# Install binaries
install:
	@echo "[5] Installing GPM binaries..." | tee -a $(LOG_FILE)
	@sudo mv -f $(GPM_BINARY) /usr/local/bin/
ifneq ($(ENABLE_AUTOBUILD),false)
	@sudo mv -f $(AUTOBUILD_BINARY) /usr/local/bin/
endif
	@echo "Install complete."

# Clean build artifacts
clean:
	@echo "[CLEAN] Removing old builds..." | tee -a $(LOG_FILE)
	@rm -rf $(GPM_BINARY) $(AUTOBUILD_BINARY) GpmAPI $(LOG_FILE)

# Run GPM
run:
	@echo "[RUNNING GPM]" | tee -a $(LOG_FILE)
	@./$(GPM_BINARY) | tee -a $(LOG_FILE)

# Show feature status
status:
	@echo "=== GPM Build Features ==="
	@echo " USE_SSH          = $(USE_SSH)"
	@echo " USE_HTTPS        = $(USE_HTTPS)"
	@echo " ENABLE_AUTOBUILD = $(ENABLE_AUTOBUILD)"
	@echo " ENABLE_API       = $(ENABLE_API)"
	@echo "=========================="
