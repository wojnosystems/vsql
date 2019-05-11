# Getting started



# Motivation

I started learning Go for a work project not too long ago. I've been struggling with the database interfaces that come stock with Go. While I can see why some decisions were made, the interface is cumbersome and difficult to extend.

To overcome these problems, this library is essentially an interface facade that breaks up the parts based on functionality (imagine that!) and capability. For example, I was extremely perturbed using the database.sql library that I had to cast to a txn type because my calls could either be top-levels themselves OR transactions. These interfaces will hide the transaction status.

However, because some databases (cough!--MYSQL--cough!) don't support nested transactions, I've split the interfaces based on the capabilities of the underlying database. One supports Nested Transactions, the other does not. ONLY difference is that the NestedTransactionStarter returns a NestedTransactioner instead of just a Transactioner. This will help ensure that you write code to conform to the interfaces. If you decide you want to use nested transactions and move to a database that supports them, it should be as easy as changing which interfaces you use and adding the Begin calls that you need.

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

