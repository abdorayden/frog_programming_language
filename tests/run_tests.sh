#!/bin/bash

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Path to the frog interpreter
go build ..
FROG_INTERPRETER="../frog_programming_language"

# Check if the interpreter exists and is executable
if [ ! -x "$FROG_INTERPRETER" ]; then
    echo -e "${RED}Error: Frog interpreter not found or not executable at $FROG_INTERPRETER${NC}"
    exit 1
fi

# Counters
FAILED_TESTS=0
PASSED_TESTS=0
TOTAL_TESTS=0

# Iterate over all .frg files in the current directory that have a .expected file
for test_file in *.frg; do
    if [ -f "$test_file" ]; then
        expected_file="${test_file}.expected"
        if [ -f "$expected_file" ]; then
            TOTAL_TESTS=$((TOTAL_TESTS + 1))
            echo -e "${BLUE}Running test:${NC} $test_file"
            
            # Run the interpreter and capture stdout
            output=$(./$FROG_INTERPRETER "$test_file" 2> /dev/null)
            
            # Read the expected output
            expected_output=$(cat "$expected_file")
            
            # Compare the output with the expected output
            if [ "$output" = "$expected_output" ]; then
                echo -e "  ${GREEN}[PASS]${NC}"
                PASSED_TESTS=$((PASSED_TESTS + 1))
            else
                echo -e "  ${RED}[FAIL]${NC}"
                echo "    Expected:"
                echo -e "      ${GREEN}$expected_output${NC}"
                echo "    Got:"
                echo -e "      ${RED}$output${NC}"
                FAILED_TESTS=$((FAILED_TESTS + 1))
            fi
        fi
    fi
done

# Special case for error testing
# ERROR_TEST="error.frg"
# if [ -f "$ERROR_TEST" ]; then
#     TOTAL_TESTS=$((TOTAL_TESTS + 1))
#     echo -e "${BLUE}Running test:${NC} $ERROR_TEST (expecting error)"
#     # Run the interpreter and capture only stderr
#     error_output=$(./$FROG_INTERPRETER "$ERROR_TEST" 2>&1 1>/dev/null)
#     if [ -n "$error_output" ]; then
#         echo -e "  ${GREEN}[PASS]${NC}"
#         PASSED_TESTS=$((PASSED_TESTS + 1))
#     else
#         echo -e "  ${RED}[FAIL]${NC}"
#         echo "    Expected an error message, but got none."
#         FAILED_TESTS=$((FAILED_TESTS + 1))
#     fi
# fi

# Final summary
echo
echo "--------------------"
echo "Test Summary"
echo "--------------------"
echo -e "Total tests:   $TOTAL_TESTS"
echo -e "${GREEN}Passed tests:${NC}  $PASSED_TESTS"
echo -e "${RED}Failed tests:${NC}  $FAILED_TESTS"
echo "--------------------"

if [ $FAILED_TESTS -eq 0 ]; then
    echo -e "${GREEN}All tests passed!${NC}"
    exit 0
else
    echo -e "${RED}$FAILED_TESTS tests failed.${NC}"
    exit 1
fi
