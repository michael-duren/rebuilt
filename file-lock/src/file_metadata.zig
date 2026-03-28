const std = @import("std");
const fs = std.fs;

pub const FileMetadata = struct {
    /// metadata from std handles format from different OS's
    md: fs.File.Stat,
    path: []const u8,
    /// optional content hash
    checksum: ?[]const u8,
    allocator: std.mem.Allocator,

    pub fn init(allocator: std.mem.Allocator, path: []const u8, hash_content: bool) !FileMetadata {
        const abs_path = try fs.realpathAlloc(allocator, path);
        errdefer allocator.free(abs_path);

        const file = try fs.openFileAbsolute(abs_path, .{});
        defer file.close();

        const stat = try file.stat();
        var metadata = FileMetadata{
            .path = abs_path,
            .md = stat,
            .checksum = null,
            .allocator = allocator,
        };
        if (hash_content) {
            metadata.checksum = try computeFileHash(file, allocator);
        }

        return metadata;
    }

    fn computeFileHash(f: fs.File, allocator: std.mem.Allocator) ![]const u8 {
        var hasher = std.crypto.hash.sha2.Sha256.init(.{});

        // use 4kb buf
        var buf: [4096]u8 = undefined;

        try f.seekTo(0);

        while (true) {
            const bytes_read = try f.read(&buf);
            if (bytes_read == 0) break;

            hasher.update(buf[0..bytes_read]);
        }

        var hash: [32]u8 = undefined;
        hasher.final(&hash);

        const hex_hash = try allocator.alloc(u8, 64);
        _ = std.fmt.bufPrint(hex_hash, "{x}", .{hash}) catch unreachable;
        return hex_hash;
    }

    pub fn deinit(self: *const FileMetadata) void {
        self.allocator.free(self.path);
        if (self.checksum) |checksum| {
            self.allocator.free(checksum);
        }
    }

    pub fn clone(fmd: FileMetadata) !FileMetadata {
        var new_metadata = FileMetadata{
            .path = try fmd.allocator.dupe(u8, fmd.path),
            .md = fmd.md,
            .checksum = null,
            .allocator = fmd.allocator,
        };

        if (fmd.checksum) |checksum| {
            new_metadata.checksum = try fmd.allocator.dupe(u8, checksum);
        }

        return new_metadata;
    }
};
