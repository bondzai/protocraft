# Makefile for generating Protocol Buffers code

PROTO_SRC_DIR := .
PROTO_OUT_DIR := ./generated
PROTOC := protoc
PROTOC_LANG := python

PROTO_FILES := $(shell find $(PROTO_SRC_DIR) -name '*.proto')
PROTO_TARGETS := $(patsubst $(PROTO_SRC_DIR)/%.proto,$(PROTO_OUT_DIR)/%_pb2.py,$(PROTO_FILES))

$(PROTO_OUT_DIR)/%_pb2.py: $(PROTO_SRC_DIR)/%.proto
	@echo "Generating $@ from $<"
	mkdir -p $(PROTO_OUT_DIR)
	$(PROTOC) --proto_path=$(PROTO_SRC_DIR) --proto_path=$(PROTO_OUT_DIR) --$(PROTOC_LANG)_out=$(PROTO_OUT_DIR) $<

all: $(PROTO_TARGETS)

clean:
	@echo "Cleaning generated files"
	rm -rf $(PROTO_OUT_DIR)

.PHONY: all clean
