package sqldb

import (
	"database/sql"
	"errors"
	"fmt"
)

type myDB struct {
	DataName string
	sqlpath  string
	db       *sql.DB
}

var (
	mndb myDB
)

// GetSQLPath 组装sql连接路径 username:password@tcp(ip:port)/dbname?charset=CHARSET&timeout=TIME
func GetSQLPath(user, passwd, host, dbname string) string {
	sqlpath := user + ":" + passwd + "@tcp(" + host + ")/" + dbname + "?charset=utf8&timeout=10s"
	return sqlpath
}

// DBInit 数据库连接 sqlpath = GetSQLPath(...)
func DBInit(dbname, sqlpath string) bool {
	ok, dbnm := dbInit(dbname, sqlpath)
	mndb = myDB{dbname, sqlpath, dbnm}
	return ok
}

func dbInit(dbname, address string) (bool, *sql.DB) {
	//fmt.Println("conn db:", dbname)
	var (
		dbobj *sql.DB
		err   error
	)
	dbobj, err = sql.Open("mysql", address)

	if err != nil {
		fmt.Println("open mysql err:", err)
		return false, dbobj
	}

	dbobj.SetMaxOpenConns(1000)
	dbobj.SetMaxIdleConns(500)
	err = dbobj.Ping()

	if err != nil {
		fmt.Println("Ping database error: ", err)
		return false, dbobj
	}
	//fmt.Println("conn_mysql_ok:", dbname)
	return true, dbobj
}

const (
	// MaxTryCount 最大重试次数
	MaxTryCount = 3
)

// Select 执行sql语句 select
func Select(sid, msql string) ([]map[string]string, error) {
	if nil == mndb.db {
		bl := false
		if bl, mndb.db = dbInit(mndb.DataName, mndb.sqlpath); !bl {
			fmt.Println(sid, " DBInit return false.")
			return nil, errors.New(sid + " DBInit return false")
		}
	}
	rows, err := mndb.db.Query(msql)
	if err != nil {
		fmt.Println(sid, "Query error: ", err)
		return nil, errors.New(sid + " Query error")
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println(sid, " Failed to get columns", err)
		return nil, errors.New(sid + " Failed to get columns")
	}
	var mpAr []map[string]string
	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println(sid, " rows.Scan error: ", err)
			return nil, errors.New(sid + " error")
		}
		record := make(map[string]string)
		for k, v := range rawResult {
			if v != nil {
				record[cols[k]] = string(v)
			} else {
				record[cols[k]] = ""
			}
		}
		mpAr = append(mpAr, record)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(sid, " rows error: ", err)
		return nil, errors.New(sid + " rows error")
	}

	if len(mpAr) > 0 {
		return mpAr, nil
	}
	return nil, errors.New(sid + " find result is null")
}
