#ifndef LEXER
#define LEXER

#include "token.h"

typedef struct {
    const char* input;
    int pos;
    int read_pos;
    int len;
} Lexer;

Lexer lexer_new(const char* input);
Token lexer_next(Lexer* l);

#endif  // LEXER
