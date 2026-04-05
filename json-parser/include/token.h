#ifndef TOKENS
#define TOKENS

typedef enum {
    TOKEN_LBRACE,
    TOKEN_RBRACE,
    TOKEN_DELIMITER,

    TOKEN_COLON,

    TOKEN_COMMA,

    TOKEN_KEY,
    TOKEN_VALUE,

    TOKEN_EOF,
    TOKEN_ERROR,
} TokenType;

typedef struct {
    TokenType type;
} Token;

#endif  // !TOKENS
