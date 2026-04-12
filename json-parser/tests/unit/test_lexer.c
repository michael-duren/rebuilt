#include "unity.h"
#include "lexer.h"
#include "token.h"
#include <string.h>

void setUp(void) {
    // Set up code that runs before each test
}

void tearDown(void) {
    // Clean up code that runs after each test
}

void test_lexer_new_initializes_correctly(void) {
    const char* input = "{}";
    Lexer lexer = lexer_new(input);

    TEST_ASSERT_EQUAL_PTR(input, lexer.input);
    TEST_ASSERT_EQUAL_INT(0, lexer.pos);
    TEST_ASSERT_EQUAL_INT(2, lexer.len);
}

void test_lexer_recognizes_left_brace(void) {
    const char* input = "{";
    Lexer lexer = lexer_new(input);

    Token token = lexer_next(&lexer);

    TEST_ASSERT_EQUAL_INT(TOKEN_LBRACE, token.type);
}

void test_lexer_recognizes_right_brace(void) {
    const char* input = "}";
    Lexer lexer = lexer_new(input);

    Token token = lexer_next(&lexer);

    TEST_ASSERT_EQUAL_INT(TOKEN_RBRACE, token.type);
}

void test_lexer_recognizes_delimiter(void) {
    const char* input = "\"";
    Lexer lexer = lexer_new(input);

    Token token = lexer_next(&lexer);

    TEST_ASSERT_EQUAL_INT(TOKEN_DELIMITER, token.type);
}

void test_lexer_returns_error_for_unknown_char(void) {
    const char* input = "@";
    Lexer lexer = lexer_new(input);

    Token token = lexer_next(&lexer);

    TEST_ASSERT_EQUAL_INT(TOKEN_ERROR, token.type);
}

void test_lexer_returns_eof_on_empty_input(void) {
    const char* input = "";
    Lexer lexer = lexer_new(input);

    Token token = lexer_next(&lexer);

    TEST_ASSERT_EQUAL_INT(TOKEN_EOF, token.type);
}

void test_lexer_handles_multiple_tokens(void) {
    const char* input = "{}";
    Lexer lexer = lexer_new(input);

    Token token1 = lexer_next(&lexer);
    TEST_ASSERT_EQUAL_INT(TOKEN_LBRACE, token1.type);

    Token token2 = lexer_next(&lexer);
    TEST_ASSERT_EQUAL_INT(TOKEN_RBRACE, token2.type);

    Token token3 = lexer_next(&lexer);
    TEST_ASSERT_EQUAL_INT(TOKEN_EOF, token3.type);
}

void test_lexer_skips_whitespace(void) {
    const char* input = "  {  }  ";
    Lexer lexer = lexer_new(input);

    Token token1 = lexer_next(&lexer);
    TEST_ASSERT_EQUAL_INT(TOKEN_LBRACE, token1.type);

    Token token2 = lexer_next(&lexer);
    TEST_ASSERT_EQUAL_INT(TOKEN_RBRACE, token2.type);
}
