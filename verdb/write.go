package verdb

import (
	"encoding/json"
)

// Write writes database back to the file where it was read from.
func (db *Verdb) Write() (err error) {
	db.dbFile.Seek(0, 0)
	db.dbFile.Truncate(0)
	encoder := json.NewEncoder(db.dbFile)
	encoder.SetIndent("", "	")
	err = encoder.Encode(db)
	return
}
