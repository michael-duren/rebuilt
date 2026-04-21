// TODO: implement
/**
 * The Entry class represents a single entry in the hash map, containing a key and
value.
 */
typedef struct {
    char* key;   // Key associated with this entry
    int value;   // Value associated with this key
} Entry;

// TODO: implement
/**
 * The HashMap class represents a hash map data structure that can store entries in an
array-like manner.
 */
typedef struct {
    int size;      // Number of slots (array length) for the hash map
    int count;     // Current number of items in the hash map
    Entry** items; // Pointer to array of pointers to Entries, representing the data structure
} HashMap;

// TODO: implement
/**
 * Creates a new hash map with the specified size.
 * @param size Number of slots (array length) for the hash map
 * @return A pointer to the newly created HashMap
 */
HashMap* create_hashmap(int size);

// TODO: implement
/**
 * Inserts a new entry into the hash map, associating it with the specified key.
 * If an entry with the same key already exists in the map, its value is updated.
 * @param map The HashMap to insert into
 * @param key Key to associate with the new or existing Entry
 * @param value Value to set for the associated key
 */
void insert_hashmap(HashMap* map, char* key, int value);
