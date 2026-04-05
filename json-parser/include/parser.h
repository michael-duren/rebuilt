#ifndef PARSER
#define PARSER

#include "lexer.h"
#include "token.h"

typedef struct {
    Lexer lexer;
    Token current;
} Parser;

Parser parser_new(const char* input);
bool parse_json(Parser* p);

typedef struct {
    bool success;
    bool finished;
} ParseResult;

#endif  // PARSER
//
