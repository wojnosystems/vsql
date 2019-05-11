# Getting started

## What is this?

This? This is a library, really. It's a facade around the built-in database/sql package provided by Go. This module itself is just the facade and a few helper methods I've written over and over again in various forms throughout my Go+Database career. To use this module, you need to pair it with another module that implements these interfaces, but for the database you wish to use, such as vsql_mysql or vsql_postgres.

## Quickstart

1. Get a database-specific implementation: ```go get github.com/wojnosystems/vsql_mysql```
1. Create your program:

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/wojnosystems/vsql"
    "github.com/wojnosystems/vsql/param"
    "github.com/wojnosystems/vsql_mysql"
	"log"
	"strings"
)

func main() {
	ctx := context.Background()
	// Connect to your database
	db := vsql_mysql.NewMySQL(func() (db *sql.DB) {
		cfg := mysql.Config{
			User:                 os.Getenv("MYSQL_USER"),
			Passwd:               os.Getenv("MYSQL_PASSWORD"),
			Addr:                 os.Getenv("MYSQL_ADDR"),
			DBName:               os.Getenv("MYSQL_DBNAME"),
			AllowNativePasswords: true,
		}
		if strings.HasPrefix(cfg.Addr, "unix") {
			cfg.Net = "unix"
		} else {
			cfg.Net = "tcp"
		}
		db, err := sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Fatal("unable to initialize the MySQL driver", err)
		}
		return db
	} )
	
	// interact with your database
	err := db.Ping(ctx)
	if err != nil {
		log.Fatal("cannot connect to the database")
	}
	
	// create your tables
	db.Exec(context.Background(), param.NewAppend(
       		fmt.Sprint(`CREATE TABLE IF NOT EXISTS `,
       			vsql.BT("mytable"),
       			` ( id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), age TINYINT UNSIGNED )`)))
}

```

# Motivation

Look, Go has a lot going for it and the packages that are part of core are largely awesome. Nobody writes perfect code and this module itself isn't perfect. Thank you to the Go team for all your hard work. Let me give back with a bit of my own for this module ;).

I started learning Go for a work project not too long ago. I've been struggling with the database interfaces that come stock with Go. While I can see why some decisions were made, the interface is cumbersome and difficult to extend. It's impossible to mock out of the box and the interfaces are non-existent. Instead of thinking through what to expose, the objects are shared directly and expose interface calls to the developer, even if they are not sensical in that sense. I've written a version of this library for work, but it's hodgepodge and also didn't consider the larger implications for interfaces (mostly because I had no idea how to use them properly at the time).

One of my biggest pet peeves was in writing a function that implemented a database request and took in a `*sql.DB`. But that means that the code assumes that it's only run OUTSIDE of a transaction. If you wanted that call to run within a transaction, you had to write the method over again or take in an optional argument to represent an optional transaction state. But this is error-prone and extremely un-clean code. The method implementing a database call should not really care if it's in a transaction or not and should take in a vquery.Queryer instead of an object that advertises details about transactions.

To overcome these problems, this library is essentially an interface facade that breaks up the parts based on functionality (imagine that!) and capability. For example, I was extremely perturbed using the database.sql library that I had to cast to a txn type because my calls could either be top-levels themselves OR transactions. These interfaces will hide the transaction status.

However, because some databases (cough!--MYSQL--cough!) don't support nested transactions, I've split the interfaces based on the capabilities of the underlying database. One supports Nested Transactions, the other does not. The ONLY difference is that the NestedTransactionStarter returns a NestedTransactioner instead of just a Transactioner. This will help ensure that you write code to conform to the interfaces. If you decide you want to use nested transactions and move to a database that supports them, it should be as easy as changing which interfaces you use and adding the Begin calls that you need. However, the library that is implementing the vsql interfaces will need to support this.

## Contexts

I know the database/sql library has versions of the database calls that do not have contexts, I've opted to force all calls to have contexts. If you want to ignore it, use a context.Background(). This simplifies the interface and gives you more control.

# Mocks

I've added stretchr/mock structs for your convenience. I'm constantly mocking database interactions and you can use these mocks to avoid having to write them yourself.

All of the mock'ed structs are named after the interface that they implement+"Mock" so if you want to mock the Pinger interface, PingerMock is what you want to add your method mocks onto.

# License 

Copyright 2019 Chris Wojno

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

