package generator

import (
	"os"
	"path/filepath"
	"text/template"
)

func GenerateProject(name string) error {
	baseDir := name

	dirs := []string{
		"cmd",
		"internal/handler",
		"internal/service",
		"pkg/utils",
	}

	files := map[string]string{
		"templates/service/cmd/main.go.tpl":                 "cmd/main.go",
		"templates/service/internal/handler/handler.go.tpl": "internal/handler/handler.go",
		"templates/service/internal/service/service.go.tpl": "internal/service/service.go",
		"templates/service/pkg/utils/utils.go.tpl":          "pkg/utils/utils.go",
		"templates/service/go.mod.tpl":                      "go.mod",
		"templates/service/go.sum.tpl":                      "go.sum",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Join(baseDir, dir), os.ModePerm)
		if err != nil {
			return err
		}
	}

	for tplPath, targetPath := range files {
		err := generateFileFromTemplate(tplPath, filepath.Join(baseDir, targetPath), name)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateFileFromTemplate(templatePath, targetPath, projectName string) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	targetFile, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	data := map[string]string{
		"ProjectName": projectName,
	}

	return tmpl.Execute(targetFile, data)
}
