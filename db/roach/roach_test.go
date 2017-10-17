package roach_test

import (
	"database/sql"
	"strconv"
	"testing"

	errors "github.com/tomogoma/go-typed-errors"
	"github.com/tomogoma/seedms/config"
	"github.com/tomogoma/seedms/db/roach"
	testingH "github.com/tomogoma/seedms/testing"
)

var isInit bool

func setup(t *testing.T) config.Database {
	conf := testingH.ReadConfig(t)
	conf.Database.DBName = conf.Database.DBName + "_test"
	if !isInit {
		rdb := getDB(t, conf.Database)
		defer rdb.Close()
		err := dropAllTables(rdb, conf.Database.DBName)
		if err != nil {
			t.Fatalf("Error setting up: drop prev test tables: %v", err)
		}
		isInit = true
	}
	return conf.Database
}

func tearDown(t *testing.T, conf config.Database) {
	rdb := getDB(t, conf)
	defer rdb.Close()
	if err := delAllTables(rdb, conf.DBName); err != nil {
		t.Fatalf("Error tearing down: delete all tables: %v", err)
	}
}

func dbCreated(rdb *sql.DB, dbName string) (bool, error) {
	rslt, err := rdb.Query("SHOW databases")
	if err != nil {
		return false, errors.Newf("list databases: %v", err)
	}
	defer rslt.Close()
	for rslt.Next() {
		var name string
		if err := rslt.Scan(&name); err != nil {
			return false, errors.Newf("dbName from resultset: %v", err)
		}
		if dbName == name {
			return true, nil
		}
	}
	if err := rslt.Err(); err != nil {
		return false, errors.Newf("iterating resultset: %v", err)
	}
	return false, nil
}

func delAllTables(rdb *sql.DB, dbName string) error {
	haveDB, err := dbCreated(rdb, dbName)
	if err != nil {
		return errors.Newf("check test db created: %v", err)
	}
	if !haveDB {
		return nil
	}
	for i := len(roach.AllTableNames) - 1; i >= 0; i-- {
		_, err := rdb.Exec("DELETE FROM " + roach.AllTableNames[i])
		if err != nil {
			return errors.Newf("delete %s: %v", roach.AllTableNames[i], err)
		}
	}
	return nil
}

func dropAllTables(rdb *sql.DB, dbName string) error {
	haveDB, err := dbCreated(rdb, dbName)
	if err != nil {
		return errors.Newf("check test db created: %v", err)
	}
	if !haveDB {
		return nil
	}
	for i := len(roach.AllTableNames) - 1; i >= 0; i-- {
		_, err := rdb.Exec("DROP TABLE IF EXISTS " + roach.AllTableNames[i])
		if err != nil {
			return errors.Newf("drop %s: %v", roach.AllTableNames[i], err)
		}
	}
	return nil
}

func TestNewRoach(t *testing.T) {
	conf := setup(t)
	defer tearDown(t, conf)
	tt := []struct {
		name   string
		opts   []roach.Option
		expErr bool
	}{
		{
			name: "valid",
			opts: []roach.Option{
				roach.WithDBName(conf.DBName),
				roach.WithDSN(conf.FormatDSN()),
			},
			expErr: false,
		},
		{
			name:   "valid (no options)",
			opts:   nil,
			expErr: false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := roach.NewRoach(tc.opts...)
			if r == nil {
				t.Fatalf("Got nil roach")
			}
		})
	}
}

func TestRoach_InitDBIfNot(t *testing.T) {

	conf := setup(t)
	defer tearDown(t, conf)
	r := newRoach(t, conf)
	rdb := getDB(t, conf)
	defer rdb.Close()
	if err := r.InitDBIfNot(); err != nil {
		t.Fatalf("Initial init call failed: %v", err)
	}

	tt := []struct {
		name       string
		hasVersion bool
		version    []byte
		expErr     bool
	}{
		{
			name:       "first use",
			hasVersion: false,
			expErr:     false,
		},
		{
			name:       "versions equal",
			hasVersion: true,
			version:    []byte(strconv.Itoa(roach.Version)),
			expErr:     false,
		},
		{
			name:       "db version smaller",
			hasVersion: true,
			version:    []byte(strconv.Itoa(roach.Version - 1)),
			expErr:     true,
		},
		{
			name:       "db version bigger",
			hasVersion: true,
			version:    []byte(strconv.Itoa(roach.Version + 1)),
			expErr:     true,
		},
	}

	cols := roach.ColDesc(roach.ColKey, roach.ColValue, roach.ColUpdateDate)
	updCols := roach.ColDesc(roach.ColValue, roach.ColUpdateDate)
	upsertQ := `
		INSERT INTO ` + roach.TblConfigurations + ` (` + cols + `)
			VALUES ('db.version', $1, CURRENT_TIMESTAMP)
			ON CONFLICT (` + roach.ColKey + `)
			DO UPDATE SET (` + updCols + `) = ($1, CURRENT_TIMESTAMP)`
	delQ := `
		DELETE FROM ` + roach.TblConfigurations + `
			WHERE ` + roach.ColKey + `='db.version'`

	for _, tc := range tt {
		if _, err := rdb.Exec(delQ); err != nil {
			t.Fatalf("Error setting up: clear previous config: %v", err)
		}
		if tc.hasVersion {
			if _, err := rdb.Exec(upsertQ, tc.version); err != nil {
				t.Fatalf("Error setting up: insert test config: %v", err)
			}
		}
		t.Run(tc.name, func(t *testing.T) {
			r = newRoach(t, conf)
			err := r.InitDBIfNot()
			if tc.expErr {
				if err == nil {
					t.Fatalf("Expected an error, got nil")
				}
				// set db to have correct version (init error should be cached not queried)
				if _, err := rdb.Exec(upsertQ, []byte(strconv.Itoa(roach.Version))); err != nil {
					t.Fatalf("Error setting up: insert test config: %v", err)
				}
				if err := r.InitDBIfNot(); err == nil {
					t.Fatalf("Subsequent init db not returning error")
				}
				return
			}
			if err != nil {
				t.Fatalf("Got an error: %v", err)
			}
			// set db to have incorrect version (isInit flag should be cached, not queried)
			if _, err := rdb.Exec(upsertQ, []byte(strconv.Itoa(roach.Version+10))); err != nil {
				t.Fatalf("Error setting up: insert test config: %v", err)
			}
			if err = r.InitDBIfNot(); err != nil {
				t.Fatalf("Subsequent init not working")
			}
		})
	}
}

func newRoach(t *testing.T, conf config.Database) *roach.Roach {
	r := roach.NewRoach(
		roach.WithDBName(conf.DBName),
		roach.WithDSN(conf.FormatDSN()),
	)
	if r == nil {
		t.Fatalf("Got nil roach")
	}
	return r
}

func getDB(t *testing.T, conf config.Database) *sql.DB {
	DB, err := sql.Open("postgres", conf.FormatDSN())
	if err != nil {
		t.Fatalf("new db instance: %s", err)
	}
	return DB
}
