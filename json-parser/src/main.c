#include <stdio.h>
#include <stdlib.h>

#include "../include/parser.h"

int main(int argc, char* argv[]) {
    if (argc != 2) {
        fprintf(stderr, "Usage: %s <file.json>\n", argv[0]);
        return 1;
    }

    FILE* f = fopen(argv[1], "r");
    if (!f) {
        fprintf(stderr, "Error: cannot open %s\n", argv[1]);
        return 1;
    }

    fseek(f, 0, SEEK_END);
    long fsize = ftell(f);
    rewind(f);

    char* buf = malloc(fsize + 1);
    if (!buf) {
        fprintf(stderr, "Error: out of memory\n");
        fclose(f);
        return 1;
    }

    fread(buf, 1, fsize, f);
    buf[fsize] = '\0';
    fclose(f);

    Parser p = parser_new(buf);
    int valid = parse_json(&p);
    free(buf);

    if (valid) {
        printf("Valid JSON\n");
        return 0;
    }

    printf("Invalid JSON\n");
    return 1;
}
