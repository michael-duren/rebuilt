#ifndef TOKENS
#define TOKENS

typedef enum {
    TOKEN_LBRACE,
    TOKEN_RBRACE,
    TOKEN_STRING,
    TOKEN_NUMBER,
    TOKEN_BOOL,

    TOKEN_COLON,
    TOKEN_COMMA,

    TOKEN_EOF,
    TOKEN_ERROR,
} TokenType;

typedef struct {
    TokenType type;
} Token;

#endif  // !TOKENS
