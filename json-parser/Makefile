CC = gcc
CFLAGS = -Wall -Wextra -Wpedantic -std=c23 -Iinclude
LDFLAGS =

SRC_DIR = src
OBJ_DIR = obj
BIN_DIR = bin
INC_DIR = include
UNITY_DIR = unity
TEST_DIR = tests/unit
TEST_OBJ_DIR = obj/test

TARGET = $(BIN_DIR)/json-parser
TEST_TARGET = $(BIN_DIR)/test_runner

SRCS = $(wildcard $(SRC_DIR)/*.c)
OBJS = $(SRCS:$(SRC_DIR)/%.c=$(OBJ_DIR)/%.o)

# Unity and test files
UNITY_SRC = $(UNITY_DIR)/unity.c
UNITY_OBJ = $(TEST_OBJ_DIR)/unity.o

# Source files without main.c for testing
TEST_SRCS = $(filter-out $(SRC_DIR)/main.c, $(SRCS))
TEST_OBJS = $(TEST_SRCS:$(SRC_DIR)/%.c=$(TEST_OBJ_DIR)/%.o)

# Unit test files
UNIT_TEST_SRCS = $(wildcard $(TEST_DIR)/*.c)
UNIT_TEST_OBJS = $(UNIT_TEST_SRCS:$(TEST_DIR)/%.c=$(TEST_OBJ_DIR)/%.o)

.PHONY: all clean run test unit-test

all: $(TARGET)

$(TARGET): $(OBJS) | $(BIN_DIR)
	$(CC) $(LDFLAGS) $^ -o $@

$(OBJ_DIR)/%.o: $(SRC_DIR)/%.c | $(OBJ_DIR)
	$(CC) $(CFLAGS) -c $< -o $@

$(BIN_DIR) $(OBJ_DIR) $(TEST_OBJ_DIR):
	mkdir -p $@

# Unity test framework compilation
$(UNITY_OBJ): $(UNITY_SRC) | $(TEST_OBJ_DIR)
	$(CC) $(CFLAGS) -I$(UNITY_DIR) -c $< -o $@

# Test object files (compiled without main.c)
$(TEST_OBJ_DIR)/%.o: $(SRC_DIR)/%.c | $(TEST_OBJ_DIR)
	$(CC) $(CFLAGS) -I$(UNITY_DIR) -c $< -o $@

# Unit test compilation
$(TEST_OBJ_DIR)/%.o: $(TEST_DIR)/%.c | $(TEST_OBJ_DIR)
	$(CC) $(CFLAGS) -I$(UNITY_DIR) -c $< -o $@

# Link test runner
$(TEST_TARGET): $(UNITY_OBJ) $(TEST_OBJS) $(UNIT_TEST_OBJS) | $(BIN_DIR)
	$(CC) $(LDFLAGS) $^ -o $@

clean:
	rm -rf $(OBJ_DIR) $(BIN_DIR)

run: $(TARGET)
	./$(TARGET) ./tests/step1/valid.json

test: all
	./test.sh

debug-test: all
	./test.sh -v

unit-test: $(TEST_TARGET)
	./$(TEST_TARGET)

.PHONY: bear
bear:
	bear -- $(MAKE) clean all
