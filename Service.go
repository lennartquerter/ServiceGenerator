package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"os"
	"text/template"
	"path/filepath"
)

func GenerateService(name string, context *string, postgres *bool) (err error) {

	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	var dirs = strings.Split(dir, "/")
	var root = dirs[len(dirs)-1]

	baseDir := "/Users/lennartquerter/Documents/Go/src/Generator/Template/"

	if context == nil {
		t := "DatabaseContext"
		context = &t
	}

	data := ServiceNaming{
		Name: name,

		Controller: fmt.Sprintf("%sController", name),

		Service:   fmt.Sprintf("%sService", name),
		NameLower: strings.ToLower(name),
		Interface: fmt.Sprintf("I%sService", name),
		Root:      strings.Split(root, ".")[0],
		NameSpace: root,

		Context: *context,

		CreateRequest: fmt.Sprintf("Create%sRequest", name),
		UpdateRequest: fmt.Sprintf("Update%sRequest", name),
		Response:      fmt.Sprintf("%sResponse", name),

		TableName:       fmt.Sprintf(`[%s]`, name),
		TableSelectName: fmt.Sprintf(`[%s]`, name[1:3]),
		Pagination:      `ORDER BY {pagination.OrderBy} {desc} OFFSET {pagination.Offset} ROWS FETCH NEXT {pagination.Limit} ROWS ONLY`,
	}

	if postgres != nil {
		data.TableName = fmt.Sprintf(`\"%s\"`, name)
		data.TableSelectName = fmt.Sprintf(`\"%s\"`, name[1:3])
		data.Pagination = `ORDER BY {TableSelectName}.\"{pagination.OrderBy}\" {(pagination.OrderByDesc ? "ASC" : "DESC")} LIMIT {pagination.Limit} OFFSET {pagination.Offset}`
	}

	files := []FileSystem{
		{
			Name: "Controller",
			Src:  "Controller",
			Dest: fmt.Sprintf("%s.cs", data.Controller),
			Dir:  []string{"Controllers", data.Controller},
		},
		{
			Name: "Interface",
			Src:  "Service",
			Dest: fmt.Sprintf("%s.cs", data.Interface),
			Dir:  []string{"Services", data.Service, "Abstract"},
		},
		{
			Name: "Service",
			Src:  "Service",
			Dest: fmt.Sprintf("%s.cs", data.Service),
			Dir:  []string{"Services", data.Service},
		},
		{
			Name: "Create",
			Src:  "Service",
			Dest: fmt.Sprintf("%s.%s.cs", data.Service, "Create"),
			Dir:  []string{"Services", data.Service},
		},
		{
			Name: "Update",
			Src:  "Service",
			Dest: fmt.Sprintf("%s.%s.cs", data.Service, "Update"),
			Dir:  []string{"Services", data.Service},
		},
		{
			Name: "Get",
			Src:  "Service",
			Dest: fmt.Sprintf("%s.%s.cs", data.Service, "Get"),
			Dir:  []string{"Services", data.Service},
		},
		{
			Name: "GetOne",
			Src:  "Service",
			Dest: fmt.Sprintf("%s.%s.cs", data.Service, "GetOne"),
			Dir:  []string{"Services", data.Service},
		},
		{
			Name: "Delete",
			Src:  "Service",
			Dest: fmt.Sprintf("%s.%s.cs", data.Service, "Delete"),
			Dir:  []string{"Services", data.Service},
		},
	}

	for _, v := range files {

		dat, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/%s.template", baseDir, v.Src, v.Name))

		if err != nil {
			return err
		}

		t, err := template.New(v.Name).Parse(string(dat))

		if err != nil {
			return err
		}

		path := currentDir

		for _, x := range v.Dir {
			path += fmt.Sprintf("/%s", x)
		}

		os.MkdirAll(fmt.Sprintf("%s", path), os.ModePerm)

		f, err := os.Create(fmt.Sprintf("%s/%s", path, v.Dest))

		if err != nil {
			return err
		}

		err = t.Execute(f, data)
		if err != nil {
			return err
		}

		f.Close()

		if err != nil {
			return err
		}
	}

	return nil
}
