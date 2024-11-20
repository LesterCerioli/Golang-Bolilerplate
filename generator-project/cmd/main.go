package main



func main() {
	projectName ;= flag.string("name", "", "Name of project")
	flag.Parse()

	if *projectName == "" {
		log.Fatal("Please provide a project name using the -name flag.")
	}

	err := generator.GenerateProject(*projectName)
	if err != nil {
		log.Fatal("Error generating project: %v", err)
	}
	log.Printf("Project %s generated successfully.", *projectName)
}