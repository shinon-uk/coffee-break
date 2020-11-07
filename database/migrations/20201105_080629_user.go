package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20201105_080629 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20201105_080629{}
	m.Created = "20201105_080629"

	migration.Register("User_20201105_080629", m)
}

// Run the migrations
func (m *User_20201105_080629) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(
		"CREATE TABLE user" +
			"(" +
			"	id         INT AUTO_INCREMENT NOT NULL," +
			"	user_name  TEXT               NOT NULL," +
			"	password   VARCHAR(32)        NOT NULL," +
			"	created_at DATETIME  DEFAULT current_timestamp NOT NULL," +
			"	updated_at TIMESTAMP DEFAULT current_timestamp ON UPDATE current_timestamp NOT NULL," +
			"	PRIMARY KEY (id)" +
			");",
	)
}

// Reverse the migrations
func (m *User_20201105_080629) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE user;")
}
