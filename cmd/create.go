package cmd

import (
    "fmt"
    "github.com/alexraileanu/aoc-init/templates"
    "github.com/fatih/color"
    "os"
    "path/filepath"
    "strings"
)

type File struct {
    Name     string
    Template string
    Args     interface{}
}

var (
    errColor     = color.New(color.FgRed)
    successColor = color.New(color.FgGreen)

    filesToCreate = []File{
        {
            Name:     "main.go",
            Template: templates.MainTemplate(),
            Args:     "",
        },
        {
            Name:     "go.mod",
            Template: templates.ModTemplate(),
            Args:     "",
        },
        {
            Name:     ".gitignore",
            Template: templates.GitignoreTemplate(),
            Args:     "",
        },
        {
            Name:     "input",
            Template: "",
            Args:     "",
        },
    }
)

func Create(path string) {
    fullPath, err := filepath.Abs(path)
    if err != nil {
        errColor.Printf("Incorrect path: %v", err.Error())
        return
    }

    err = os.MkdirAll(fullPath, 0755)
    if err != nil {
        errColor.Printf("Error creating folder(s): %v\n", err.Error())
        return
    }

    var filesNotCreated []File

    for _, file := range filesToCreate {
        f, err := createFile(file.Name, fullPath)
        if err != nil {
            errColor.Printf("Error creating file %v: %v\n", file, err.Error())
            filesNotCreated = append(filesNotCreated, file)
            continue
        }

        _, err = f.WriteString(strings.Trim(file.Template, "\n"))
        if err != nil {
            errColor.Printf("Error writing contents to file %v: %v", file, err.Error())
            continue
        }

        err = f.Close()
        if err != nil {
            errColor.Printf("Error closing file handler for file %v: %v\n", file, err.Error())
            continue
        }
    }

    if len(filesNotCreated) == 0 {
        successColor.Println("All files created successfully!")
        return
    }

    errColor.Printf("Some files were not created: %v\n", filesNotCreated)
}

func createFile(fileName, parentPath string) (*os.File, error) {
    absPath, err := filepath.Abs(fmt.Sprintf("%s/%s", parentPath, fileName))
    if err != nil {
        return &os.File{}, err
    }
    f, err := os.OpenFile(absPath, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        return &os.File{}, err
    }

    return f, nil
}
