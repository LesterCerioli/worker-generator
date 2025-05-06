package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type WorkerConfig struct {
	ProjectName string
	GoVersion   string
}

func GenerateWorker(config WorkerConfig, outputDir string) error {

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	templates := map[string]string{
		"go.mod.tmpl":    filepath.Join(outputDir, "go.mod"),
		"worker.go.tmpl": filepath.Join(outputDir, "worker.go"),
	}

	for tmplFile, outputFile := range templates {
		tmpl, err := template.ParseFiles(filepath.Join("pkg/templates", tmplFile))
		if err != nil {
			return fmt.Errorf("failed to parse template: %v", err)
		}

		f, err := os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		defer f.Close()

		if err := tmpl.Execute(f, config); err != nil {
			return fmt.Errorf("template execution failed: %v", err)
		}
	}

	return nil
}
