#ifndef TOKENS
#define TOKENS

typedef enum {
    TOKEN_LBRACE,
    TOKEN_RBRACE,
    TOKEN_LDELIMITER,
    TOKEN_RDELIMITER,

    TOKEN_EOF,
    TOKEN_ERROR,
} TokenType;

typedef struct {
    TokenType type;
} Token;

#endif  // !TOKENS
