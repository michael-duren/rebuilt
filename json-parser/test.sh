#!/bin/bash
# Automated test runner for the JSON parser

PASS=0
FAIL=0

VERBOSE=0

while [[ $# -gt 0 ]]; do
    case "$1" in
    -v | --verbose)
        VERBOSE=1
        shift
        ;;
    *)
        echo "Unknown option: $1"
        exit 1
        ;;
    esac
done

run_test() {
    local file="$1"
    local expect_exit="$2" # 0 = valid, 1 = invalid
    local label="$3"

    if [[ $VERBOSE -eq 1 ]]; then
        ./bin/json-parser "$file"
    else
        ./bin/json-parser "$file" >/dev/null 2>&1
    fi

    local actual_exit=$?
    if [ "$actual_exit" -eq "$expect_exit" ]; then
        echo "  PASS: $label"
        PASS=$((PASS + 1))
    else
        echo "  FAIL: $label (expected exit $expect_exit, got $actual_exit)"
        FAIL=$((FAIL + 1))
    fi
}

if [[ ! -f "./bin/json-parser" ]]; then
    echo "binary not found"
fi

echo "=== Step 1 Tests ==="
run_test "tests/step1/valid.json" 0 "valid.json   -> valid"
run_test "tests/step1/invalid.json" 1 "invalid.json -> invalid"
run_test "tests/step2/valid.json" 0 "valid.json   -> valid"
run_test "tests/step2/valid2.json" 0 "valid.json   -> valid"
run_test "tests/step2/invalid.json" 0 "valid.json   -> valid"
run_test "tests/step2/invalid2.json" 0 "valid.json   -> valid"

echo ""
echo "Results: $PASS passed, $FAIL failed"

if [ "$FAIL" -ne 0 ]; then
    exit 1
fi
