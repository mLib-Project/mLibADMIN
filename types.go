package main

type book struct {
	category    int
	subcategory string
	name        string
	author      string
	source      string
}

type category struct {
	name   string
	family int
}
