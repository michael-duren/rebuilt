const std = @import("std");
const json_parser = @import("json_parser");

pub fn main() !void {
    var args = std.process.args();
    _ = args.next();
    // Prints to stderr, ignoring potential errors.
    const res = try json_parser.parse("{}");
    std.debug.print("Json parser worked!", .{res});
}
