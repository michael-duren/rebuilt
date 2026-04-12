#include "unity.h"

// External test function declarations
extern void test_lexer_new_initializes_correctly(void);
extern void test_lexer_recognizes_left_brace(void);
extern void test_lexer_recognizes_right_brace(void);
extern void test_lexer_recognizes_delimiter(void);
extern void test_lexer_returns_error_for_unknown_char(void);
extern void test_lexer_returns_eof_on_empty_input(void);
extern void test_lexer_handles_multiple_tokens(void);
extern void test_lexer_skips_whitespace(void);

int main(void) {
    UNITY_BEGIN();

    // Lexer tests
    RUN_TEST(test_lexer_new_initializes_correctly);
    RUN_TEST(test_lexer_recognizes_left_brace);
    RUN_TEST(test_lexer_recognizes_right_brace);
    RUN_TEST(test_lexer_recognizes_delimiter);
    RUN_TEST(test_lexer_returns_error_for_unknown_char);
    RUN_TEST(test_lexer_returns_eof_on_empty_input);
    RUN_TEST(test_lexer_handles_multiple_tokens);
    RUN_TEST(test_lexer_skips_whitespace);

    return UNITY_END();
}
