#!/bin/bash
# Automated test runner for the JSON parser

PASS=0
FAIL=0

run_test() {
  local file="$1"
  local expect_exit="$2" # 0 = valid, 1 = invalid
  local label="$3"

  ./json_parser "$file" >/dev/null 2>&1
  local actual_exit=$?

  if [ "$actual_exit" -eq "$expect_exit" ]; then
    echo "  PASS: $label"
    PASS=$((PASS + 1))
  else
    echo "  FAIL: $label (expected exit $expect_exit, got $actual_exit)"
    FAIL=$((FAIL + 1))
  fi
}

echo "=== Step 1 Tests ==="
run_test "tests/step1/valid.json" 0 "valid.json   -> valid"
run_test "tests/step1/invalid.json" 1 "invalid.json -> invalid"

echo ""
echo "Results: $PASS passed, $FAIL failed"

if [ "$FAIL" -ne 0 ]; then
  exit 1
fi
