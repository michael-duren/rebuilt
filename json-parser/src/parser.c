#include "../include/parser.h"

#include <stdio.h>

Parser parser_new(const char* input) {
    Parser p;
    p.lexer = lexer_new(input);
    p.current = lexer_next(&p.lexer);
    return p;
}

bool parse_internal(Parser* p, bool inDelimeter) {
    printf("token is: %d\n\n", p->current.type);
    if (inDelimeter) {
        return false;
    }
    return true;
}

bool parse_json(Parser* p) {
    if (p->current.type != TOKEN_LBRACE) {
        return false;
    }

    int result = parse_internal(p, false);
    if (!result) {
        return result;
    }

    p->current = lexer_next(&p->lexer);

    if (p->current.type != TOKEN_RBRACE) {
        return false;
    }

    p->current = lexer_next(&p->lexer);

    if (p->current.type != TOKEN_EOF) {
        return false;
    }

    return true;
}
