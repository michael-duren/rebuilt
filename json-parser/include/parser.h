#ifndef PARSER
#define PARSER

#include "lexer.h"
#include "token.h"

typedef struct {
    Lexer lexer;
    Token current;
} Parser;

Parser parser_new(const char* input);
int parse_json(Parser* p);

#endif  // PARSER
