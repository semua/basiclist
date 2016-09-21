BasicList: A sorted Linked List implement in Go
==================
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/basiclist/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/basiclist?status.svg)](https://godoc.org/github.com/kkdai/basiclist)  [![Build Status](https://travis-ci.org/kkdai/basiclist.svg?branch=master)](https://travis-ci.org/kkdai/basiclist)


What is Basic List
---------------

It is not a new data structure, it is a sorted linked list implement in Golang. I implement this for study linked list part for [Skip List](https://en.wikipedia.org/wiki/Skip_list).

fork change
---------------
I use a two-way linked-list data structure + map[key], to sort, compare and get interface{}

all codes are coveraged.
github.com\semua\basiclist\basiclist.go:23:	NewBasicList		100.0%
github.com\semua\basiclist\basiclist.go:31:	Insert			100.0%
github.com\semua\basiclist\basiclist.go:76:	GetByKey		100.0%
github.com\semua\basiclist\basiclist.go:83:	Equal			100.0%
github.com\semua\basiclist\basiclist.go:105:	LessThanOrEqual		100.0%
github.com\semua\basiclist\basiclist.go:127:	GreaterThanOrEqual	100.0%
github.com\semua\basiclist\basiclist.go:148:	Remove			100.0%
github.com\semua\basiclist\basiclist.go:169:	List			100.0%
github.com\semua\basiclist\basiclist.go:191:	Print			100.0%
total:						(statements)		100.0%

Install
---------------
`go get github.com/semua/basiclist`


Usage
---------------
	// first create a compare function
	func Compare(val0 interface{}, val1 interface{}) int {
		pubt0 := val0.(Article).pubtime - val0.(Article).pubtime%(3600*12)
		pubt1 := val1.(Article).pubtime - val1.(Article).pubtime%(3600*12)
		if pubt0-pubt1 == 0 {
			return 0
		}
		if pubt0-pubt1 > 0 {
			return 1
		}
		return -1
	}
	type Article struct {
		title   string
		pubtime int64
	}
	// second on your main

    // New a BasicList
    bList := NewBasicList()
    
    //Insert key could be any `interface{}`
    bList.Insert(Compare, "string3", Article{title: "title3", pubtime: time.Now().Unix()})
    bList.Insert(Compare, "string4", Article{title: "title4", pubtime: time.Now().Unix()})
    bList.Insert(Compare, "string2", Article{title: "title2", pubtime: time.Now().Unix()})
    
    //Display all linked list.
    bList.Print()
    //head->[val:{title3 1474461913}]->[val:{title4 1474461914}]->[val:{title2 1474461915}]->nil
    
    //Remove from list
    bList.Remove("title3")    
    
    bList.Print()
    //head->[val:{title4 1474461914}]->[val:{title2 1474461915}]->nil
    
	fmt.Println(bList.GetByKey("title4"))
	//{title4 1474461914} <nil>
	
	cmpObj := Article{title: "title6", pubtime: 1474461915}
	fmt.Println(bList.LessThanOrEqual(Compare, cmpObj))
	//{title2 1474461915} <nil>


License
---------------

This package is licensed under MIT license. See LICENSE for details.


