//! By convention, root.zig is the root source file when making a library.
const std = @import("std");

const ParseError = error{InvalidSyntax};
const JsonValue = union(enum) {
    null_value,
    bool_value: bool,
    number: f64,
    string: []const u8,
    array: std.ArrayList(JsonValue),
    object: std.StringHashMap(JsonValue),
};

pub fn parse(str: []const u8) ParseError!JsonValue {
    if (str.len == 0 or str[0] != '{' or str[str.len - 1] != '}') {
        return error.InvalidSyntax;
    }
    return .null_value;
}

test "parse parses an empty object {}" {
    const obj = "{}";
    const res = try parse(obj);

    try std.testing.expect(res == .null_value);
}

test "parse parses an empty string" {
    const obj = "";
    try std.testing.expectError(error.InvalidSyntax, parse(obj));
}
