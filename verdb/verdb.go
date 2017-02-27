package verdb

import "os"

// currentDbVersion is the version we expect the DB to be in.
// We also write it to new databases.
const currentDbVersion = "1.0"

// VerdbEntry is a entry of the Verdb.
// It maps a name of a package to its version.
type VerdbEntry struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Verdb is a database containing all versions of all packages.
// It is saved as JSON and versioned for future additions/changes.
type Verdb struct {
	Version string       `json:"version"`
	dbFile  *os.File     `json:"-"`
	Entries []VerdbEntry `json:"entries"`
}

// Returns the stored version of an upstream or "" if none was recorded.
func (db Verdb) VersionOf(name string) string {
	for _, entry := range db.Entries {
		if entry.Name == name {
			return entry.Version
		}
	}
	return ""
}

// Saves the version of an upstream, creating it if needed.
func (db *Verdb) UpdateVersion(name, version string) {
	num := -1
	for i, entry := range db.Entries {
		if entry.Name == name {
			num = i
		}
	}
	if num != -1 {
		db.Entries[num].Version = version
	} else {
		db.Entries = append(db.Entries, VerdbEntry{name, version})
	}
}
