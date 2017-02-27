package verdb

import (
	"encoding/json"
	"errors"
	"github.com/dasJ/versioncheck/config"
	"os"
)

// #include <sys/file.h>
import "C"

// Read reads the databse from file or creates the file if it does not exist yet.
// The path of the database is found the the VersioncheckConfig
func Read(cfg config.VersioncheckConfig) (ret Verdb, err error) {
	_, err = os.Stat(cfg.DbLocation)
	if os.IsNotExist(err) {
		file, err := os.Create(cfg.DbLocation)
		if err != nil {
			panic(err)
		}
		file.Close()
	}
	dbFile, err := os.OpenFile(cfg.DbLocation, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	// Lock database
	fd := dbFile.Fd()
	rc := C.flock(C.int(fd), C.LOCK_EX)
	if rc < 0 {
		err = errors.New("Can not lock database")
		return
	}
	// Check DB size
	stat, err := dbFile.Stat()
	if err != nil {
		C.flock(C.int(fd), C.LOCK_UN)
		return
	}
	if stat.Size() == 0 {
		return
	}
	// Parse database
	jsonParser := json.NewDecoder(dbFile)
	err = jsonParser.Decode(&ret)
	if err != nil {
		C.flock(C.int(fd), C.LOCK_UN)
		return
	}
	if ret.Version != currentDbVersion {
		err = errors.New("Invalid database version: " + ret.Version)
	}
	ret.dbFile = dbFile

	return
}
