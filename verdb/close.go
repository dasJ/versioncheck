package verdb

import (
	"errors"
)

// #include <sys/file.h>
import "C"

// Close closes the verdb file handle and release the lock.
func (db *Verdb) Close() (err error) {
	rc := C.flock(C.int(db.dbFile.Fd()), C.LOCK_UN)
	if rc < 0 {
		err = errors.New("File lock release failed with error code " + string(rc))
	}
	err = db.dbFile.Close()
	return
}
