const std = @import("std");
const testing = std.testing;
const FileMetadata = @import("file_metadata.zig").FileMetadata;

pub const FileIndex = struct {
    /// Files and their metadata
    files: std.StringHashMap(FileMetadata),
    /// Files by their ids
    inodes: std.AutoHashMap(u64, []const u8),
    /// Allocator
    allocator: std.mem.Allocator,
    /// Set of paths
    path_storage: std.StringHashMap(void),

    pub fn init(allocator: std.mem.Allocator) FileIndex {
        return FileIndex{
            .files = std.StringHashMap(FileMetadata).init(allocator),
            .inodes = std.AutoHashMap(u64, []const u8).init(allocator),
            .allocator = allocator,
            .path_storage = std.StringHashMap(void).init(allocator),
        };
    }

    pub fn deinit(self: *FileIndex) void {
        // free all metadata objects
        var it = self.files.iterator();
        while (it.next()) |entry| {
            entry.value_ptr.deinit();
        }

        var path_it = self.path_storage.keyIterator();
        while (path_it.next()) |path_ptr| {
            self.allocator.free(path_ptr.*);
        }

        self.files.deinit();
        self.inodes.deinit();
        self.path_storage.deinit();
    }

    pub fn addFile(self: *FileIndex, metadata: FileMetadata) !void {
        const cloned = try metadata.clone();
        errdefer cloned.deinit();

        const path_copy = try self.allocator.dupe(u8, cloned.path);
        errdefer self.allocator.free(path_copy);

        try self.path_storage.put(path_copy, {});
        try self.files.put(path_copy, cloned);
        if (cloned.md.inode != 0) {
            try self.inodes.put(cloned.metadata.inode, path_copy);
        }
    }

    pub fn contains(self: *const FileIndex, path: []const u8) bool {
        return self.files.contains(path);
    }

    pub fn get(self: *const FileIndex, path: []const u8) ?*const FileMetadata {
        return self.files.getPtr(path);
    }

    pub fn findByInode(self: *const FileIndex, inode: u64) ?[]const u8 {
        return self.inodes.get(inode);
    }
};

test "successfully creates file index" {
    const allocator = testing.allocator;

    const fileIndex = FileIndex.init(allocator);
    try testing.expect(@TypeOf(fileIndex) == FileIndex);

    fileIndex.deinit();
}
