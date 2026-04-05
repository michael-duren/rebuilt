#include "../include/parser.h"

#include <lexer.h>
#include <stdio.h>

#include "token.h"

Parser parser_new(const char* input) {
    Parser p;
    p.lexer = lexer_new(input);
    p.current = lexer_next(&p.lexer);
    return p;
}

static ParseResult parse_internal(Parser* p, bool in_delimiter) {
    while (true) {
        p->current = lexer_next(&p->lexer);
        TokenType type = p->current.type;

        if (in_delimiter) {
            if (type != TOKEN_DELIMITER && type != TOKEN_EOF) {
                continue;
            }
            if (type == TOKEN_EOF) {
                return (ParseResult){.success = false, .finished = true};
            }
            return (ParseResult){.success = true, .finished = true};
        }

        if (type == TOKEN_DELIMITER) {
            in_delimiter = true;
            continue;
        }

        return (ParseResult){.success = false, .finished = true};
    }
}

bool parse_json(Parser* p) {
    if (p->current.type != TOKEN_LBRACE) {
        return false;
    }

    while (true) {
        ParseResult r = parse_internal(p, false);
        if (r.finished) {
            break;
        }
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
