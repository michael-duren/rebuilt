#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef enum {
    TOKEN_LBRACE,
    TOKEN_RBRACE,
    TOKEN_EOF,
    TOKEN_ERROR,
} TokenType;

typedef struct {
    TokenType type;
} Token;

// walks through raw text one char at a time
// produces tokens.
typedef struct {
    const char* input;
    int pos;
    int len;
} Lexer;

Lexer lexer_new(const char* input) {
    Lexer l;
    l.input = input;
    l.pos = 0;
    l.len = (int)strlen(input);
    return l;
}

static void skip_whitespace(Lexer *l) {
    while (l->pos < l->len && isspace((unsigned char)l->input[l->pos])) {
        l->pos++;
    }
}

Token lexer_next(Lexer *l) {
    Token t;

    skip_whitespace(l);

    if (l->pos >= l->len) {
        t.type = TOKEN_EOF;
        return t;
    }

    char c = l->input[l->pos];
    l->pos++;

    switch (c) {
        case '{': t.type = TOKEN_LBRACE; break;
        case '}': t.type = TOKEN_RBRACE; break;
        default: t.type = TOKEN_ERROR; break;
    }

    return t;
}

typedef struct {
    Lexer lexer;
    Token current;
} Parser;

Parser parser_new(const char *input) {
    Parser p;
    p.lexer = lexer_new(input);
    p.current = lexer_next(&p.lexer);
    return p;
}

int main(int argc, char *argv[]) {
    if (argc != 2) {
        fprintf(stderr, "Usage: %s <file.json>\n", argv[0]);
        return 1;
    }

    FILE *f = fopen(argv[1], "r");
    if (!f) {
        fprintf(stderr, "Error: cannot open %s\n", argv[1]);
        return 1;
    }

    fseek(f, 0, SEEK_END);
    long fsize = ftell(f);
    rewind(f);

    char *buf = malloc(fsize + 1);
    if (!buf) {
        fprintf(stderr, "Error: out of memory\n");
        fclose(f);
        return 1;
    }

    fread(buf, 1, fsize, f);
    buf[fsize] = '\0';
    fclose(f);

    Parser p = parser_new(buf);
}
