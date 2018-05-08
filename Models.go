package main


type ServiceNaming struct {
	Name string

	Controller string

	Service   string
	NameLower string
	Interface string
	Root      string
	NameSpace string

	Context string

	CreateRequest string
	UpdateRequest string
	Response      string

	TableName       string
	TableSelectName string

	Pagination string
}

type FileSystem struct {
	Name string
	Dest string
	Src string
	Dir  []string
}