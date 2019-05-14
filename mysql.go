package storage

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/luopengift/log"
	"github.com/luopengift/types"
)

// MySQL mysql
type MySQL struct {
	*sql.DB
}

// MySQLInit init
func MySQLInit(dsn string) (*MySQL, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(6 * time.Hour)
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	return &MySQL{db}, nil
}

// Put insert into sql
// key is tablename
// value is rows message
func (mysql *MySQL) Put(ctx context.Context, key string, value map[string]interface{}, opts ...interface{}) (interface{}, error) {
	if err := mysql.Ping(); err != nil {
		return nil, err
	}
	var keyList []string
	var valueList []interface{}
	var markList []string
	for k, v := range value {
		keyList = append(keyList, k)
		valueList = append(valueList, v)
		markList = append(markList, "?")
	}
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", key, strings.Join(keyList, ", "), strings.Join(markList, ", "))
	fmt.Println(sql, fmt.Sprintf("%#v", valueList))
	stmt, err := mysql.Prepare(sql)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(valueList...)
}

// Get xx
func (mysql *MySQL) Get(ctx context.Context, key string, value map[string]interface{}, opts ...interface{}) (interface{}, error) {
	if err := mysql.Ping(); err != nil {
		return nil, err
	}
	var keyList []string
	var valueList []interface{}

	for k, v := range value {
		keyList = append(keyList, k+"=?")
		valueList = append(valueList, v)

	}

	rows, err := mysql.QueryContext(ctx, opts[0].(string), opts[1:]...)
	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	columnsType, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	var records []map[string]interface{}
	for rows.Next() {
		//将行数据保存到record字典
		if err = rows.Scan(scanArgs...); err != nil {
			return nil, err
		}
		record := make(map[string]interface{})
		for i, value := range values {
			if value != nil {
				name := columnsType[i].ScanType().Name()
				switch name {
				case "int32", "NullInt64":
					if record[columns[i]], err = types.ToInt(value); err != nil {
						return nil, err
					}
				case "RawBytes", "[]byte":
					if record[columns[i]], err = types.ToString(value); err != nil {
						return nil, err
					}
				case "NullTime":
					if record[columns[i]], err = types.ToString(value); err != nil {
						return nil, err
					}
				default:
					return nil, fmt.Errorf("%v", name)
				}
			}
		}
		records = append(records, record)

	}
	return records, nil
}

// Do xx
func (mysql *MySQL) Do(ctx context.Context, op interface{}) (interface{}, error) {
	return nil, nil
}

// Close xx
func (mysql *MySQL) Close() error {
	return nil
}

func (mysql *MySQL) Update(ctx context.Context, table string, row map[string]interface{}, selector string, value interface{}) (sql.Result, error) {
	if err := mysql.Ping(); err != nil {
		return nil, err
	}
	var updateList []string
	var valueList []interface{}
	for k, v := range row {
		updateList = append(updateList, k+"=?")
		valueList = append(valueList, v)
	}
	sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s=?", table, strings.Join(updateList, ", "), selector)
	valueList = append(valueList, value)
	log.Debugf("%s, %s", sql, fmt.Sprintf("%#v", valueList))
	stmt, err := mysql.Prepare(sql)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(valueList...)
}

func (mysql *MySQL) Delete(table string, selector string, value interface{}) (sql.Result, error) {
	err := mysql.Ping()
	if err != nil {
		return nil, err
	}
	sql := fmt.Sprintf("DELETE FROM %s WHERE %s=?", table, selector)
	fmt.Println(sql, fmt.Sprintf("%#v", value))
	stmt, err := mysql.Prepare(sql)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(value)
}
