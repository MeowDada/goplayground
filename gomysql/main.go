package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	
	// Get a handle for a SQL database, not a connection!
	db, err := sql.Open("mysql", "jack:password@(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	// Validate DSN data
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Create a database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS testDB;")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created a database")
	}

	// Choose the database
	_, err = db.Exec("USE testDB;")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully")
	}

	// Create a vaultdir table
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS vaultdir(
		did int  NOT NULL AUTO_INCREMENT,
		name varchar(256) NOT NULL,
		PRIMARY KEY(did, name));`)

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Create a vault table
	stmt, err = db.Prepare(`CREATE TABLE IF NOT EXISTS vault(
		did int,
		name varchar(256) NOT NULL,
		size int NOT NULL DEFAULT 0,
		status int NOT NULL DEFAULT 0,
		FOREIGN KEY (did) REFERENCES vaultdir(did) ON DELETE RESTRICT ON UPDATE CASCADE,
		PRIMARY KEY (did, name));`)

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Create a file table
	stmt, err = db.Prepare(`CREATE TABLE IF NOT EXISTS file(
		filename varchar(256) NOT NULL,
		orifsize BIGINT NOT NULL,
		dedupfsize BIGINT NOT NULL,
		first_created DATETIME NOT NULL,
		last_modified DATETIME NOT NULL,
		status int NOT NULL DEFAULT 0);`)
	
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
	
	// Create a file-vault mapping table
	stmt, err = db.Prepare(`CREATE TABLE IF NOT EXISTS filemap(
		vid int NOT NULL,
		fid int NOT NULL,
		FOREIGN KEY (vid) REFERENCES vault(vid) ON DELETE RESTRICT ON UPDATE CASCADE,
		FOREIGN KEY (fid) REFERENCES file(fid) ON DELETE RESTRICT ON UPDATE CASCADE,
		PRIMARY KEY (vid, fid));`)
	
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := db.Query(`SELECT v.NAME from vault v INNER JOIN filemap fm ON fm.vid = v.vid
		INNER JOIN file f ON fm.fid = f.fid WHERE f.FILENAME = ?`, "jack.tmp")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var vname string

	for rows.Next() {
		err := rows.Scan(&vname)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(vname)
	}
}