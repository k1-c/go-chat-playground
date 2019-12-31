package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Comment_20191231_105233 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Comment_20191231_105233{}
	m.Created = "20191231_105233"

	migration.Register("Comment_20191231_105233", m)
}

// Run the migrations
func (m *Comment_20191231_105233) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE comment(`id` int(11) NOT NULL AUTO_INCREMENT,`content` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Comment_20191231_105233) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `comment`")
}
