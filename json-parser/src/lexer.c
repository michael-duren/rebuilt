#include "../include/lexer.h"

#include <ctype.h>
#include <string.h>

Lexer lexer_new(const char* input) {
    Lexer l;
    l.input = input;
    l.pos = 0;
    l.len = (int)strlen(input);
    return l;
}

static void skip_whitespace(Lexer* l) {
    while (l->pos < l->len && isspace((unsigned char)l->input[l->pos])) {
        l->pos++;
    }
}

Token lexer_next(Lexer* l) {
    Token t;

    skip_whitespace(l);

    if (l->pos >= l->len) {
        t.type = TOKEN_EOF;
        return t;
    }

    char c = l->input[l->pos];
    l->pos++;

    switch (c) {
        case '{':
            t.type = TOKEN_LBRACE;
            break;
        case '}':
            t.type = TOKEN_RBRACE;
            break;
        case '"':
            t.type = TOKEN_DELIMIETER;
            break;
        default:
            t.type = TOKEN_ERROR;
            break;
    }

    return t;
}
