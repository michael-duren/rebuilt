#include "../include/lexer.h"

#include <_stdio.h>
#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static char read_next(Lexer* l) {
    if (l->len < l->read_pos) {
        return NULL;
    }

    char next = l->input[l->read_pos];
    l->read_pos++;
    l->pos++;
    return next;
}

static char peek(Lexer* l) {
    if (l->len < l->read_pos) {
        return NULL;
    }

    return l->input[l->read_pos];
}

Lexer lexer_new(const char* input) {
    Lexer l = {
        .input = input,
        .read_pos = 0,
        .pos = 0,
        .len = (int)strlen(input),
    };
    read_next(&l);
    return l;
}

static void skip_whitespace(Lexer* l) {
    while (l->pos < l->len && isspace((unsigned char)l->input[l->pos])) {
        l->pos++;
    }
}

static void skip_comment(Lexer* l) {
    while (l->pos < l->len && l->input[l->pos] != '\n' &&
           l->input[l->pos] != EOF) {
        l->pos++;
    }
}

static Token read_string(Lexer* l) {
    char* value = malloc(1024);
    int len = 0;
}

static bool is_letter(char ch) {
    return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_';
}

static bool is_digit(char ch) { return '0' <= ch && ch <= '9'; }

Token lexer_next(Lexer* l) {
    Token t;

    skip_whitespace(l);

    if (l->pos >= l->len) {
        t.type = TOKEN_EOF;
        return t;
    }

    char c = l->input[l->pos];

    switch (c) {
        case '{':
            t.type = TOKEN_LBRACE;
            l->pos++;
            break;
        case '}':
            t.type = TOKEN_RBRACE;
            l->pos++;
            break;
        case '"':
            // TODO: READ STRING
            // t.type = TOKEN_DELIMITER;
            break;
        default:
            t.type = TOKEN_ERROR;
            l->pos++;
            break;
    }

    return t;
}
