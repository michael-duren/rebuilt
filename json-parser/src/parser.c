#include "../include/parser.h"

Parser parser_new(const char* input) {
    Parser p;
    p.lexer = lexer_new(input);
    p.current = lexer_next(&p.lexer);
    return p;
}

int parse_json(Parser* p) {
    if (p->current.type != TOKEN_LBRACE) {
        return 0;
    }

    p->current = lexer_next(&p->lexer);

    if (p->current.type != TOKEN_RBRACE) {
        return 0;
    }

    p->current = lexer_next(&p->lexer);

    if (p->current.type != TOKEN_EOF) {
        return 0;
    }

    return 1;
}
