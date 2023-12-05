# Makefile for generating Protocol Buffers code

# Variables
PROTO_SRC_DIR := ./proto
PROTO_OUT_DIR := ./generated
PROTOC := protoc

# Find all .proto files
PROTO_FILES := $(shell find $(PROTO_SRC_DIR) -name '*.proto')

# Generate targets for each .proto file
PROTO_TARGETS := $(patsubst $(PROTO_SRC_DIR)/%.proto,$(PROTO_OUT_DIR)/%.pb,$(PROTO_FILES))

# Rule to generate Protocol Buffers code
$(PROTO_OUT_DIR)/%.pb: $(PROTO_SRC_DIR)/%.proto
	@echo "Generating $@ from $<"
	$(PROTOC) --proto_path=$(PROTO_SRC_DIR) --proto_path=$(PROTO_OUT_DIR) --$(PROTOC_LANG)_out=$(PROTO_OUT_DIR) $<

# Build all Protocol Buffers targets
all: $(PROTO_TARGETS)

# Clean generated files
clean:
	@echo "Cleaning generated files"
	rm -f $(PROTO_TARGETS)

# PHONY targets
.PHONY: all clean
