package main

import (
	"flag"
	"log"
	"project-generator/internal/generator" // Certifique-se de que o caminho do import está correto
)

func main() {
	projectName := flag.String("name", "", "Name of the project")
	flag.Parse()

	if *projectName == "" {
		log.Fatal("Please provide a project name using the -name flag.")
	}

	err := generator.GenerateProject(*projectName)
	if err != nil {
		log.Fatalf("Error generating project: %v", err)
	}

	log.Printf("Project %s generated successfully.", *projectName)
}
