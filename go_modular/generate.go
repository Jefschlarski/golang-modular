package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Templates para os arquivos do módulo
var templates = map[string]string{
	"module.go":               "go_modular/templates/module.tpl",
	"controller.go":           "go_modular/templates/controller.tpl",
	"dto.go":                  "go_modular/templates/dto.tpl",
	"model.go":                "go_modular/templates/model.tpl",
	"repository.go":           "go_modular/templates/repository.tpl",
	"service.go":              "go_modular/templates/service.tpl",
	"controller_interface.go": "go_modular/templates/controller_interface.tpl",
	"service_interface.go":    "go_modular/templates/service_interface.tpl",
	"repository_interface.go": "go_modular/templates/repository_interface.tpl",
}

type Module struct {
	Name      string
	NameLower string
	Projeto   string
}

// Função para converter o nome dos arquivos para snake_case (minúsculo e com underscores)
func toSnakeCase(moduleName, fileName string) string {
	// Converter o nome do módulo para lowercase
	moduleName = strings.ToLower(moduleName)

	// Separar o nome do arquivo por underscores
	parts := strings.Split(fileName, "_")

	// Iniciar o nome do arquivo com o módulo (já em lowercase)
	var snakeCaseName string
	snakeCaseName = moduleName

	// Adicionar os outros termos (completando com underscores)
	for _, part := range parts {
		if len(part) > 0 {
			snakeCaseName += "_" + strings.ToLower(part)
		}
	}

	// Retornar o nome do arquivo
	return snakeCaseName
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: <module_name> <project_name>")
		return
	}
	moduleName := os.Args[1]
	Projeto := os.Args[2]
	module := Module{Name: moduleName, NameLower: strings.ToLower(moduleName), Projeto: Projeto}

	// Caminho base para o novo módulo
	basePath := fmt.Sprintf("modules/%s", module.NameLower)
	os.MkdirAll(basePath, os.ModePerm)

	// Criar subdiretórios para os arquivos do módulo
	subdirs := []string{"controller", "dto", "model", "repository", "service"}
	for _, subdir := range subdirs {
		os.MkdirAll(fmt.Sprintf("%s/%s", basePath, subdir), os.ModePerm)
	}

	// Gerar arquivos a partir dos templates
	for fileName, templatePath := range templates {
		// Determinar a pasta onde o arquivo deve ser colocado baseado no nome do arquivo
		var folder string
		switch {
		case strings.Contains(fileName, "controller"):
			folder = "controller"
		case strings.Contains(fileName, "dto"):
			folder = "dto"
		case strings.Contains(fileName, "model"):
			folder = "model"
		case strings.Contains(fileName, "repository"):
			folder = "repository"
		case strings.Contains(fileName, "service"):
			folder = "service"
		}

		// Criar o caminho do arquivo de saída com o nome em snake_case
		snakeCaseFileName := toSnakeCase(module.NameLower, fileName)
		outFilePath := fmt.Sprintf("%s/%s/%s", basePath, folder, snakeCaseFileName)

		// Parse do arquivo template
		tpl, err := template.ParseFiles(templatePath)
		if err != nil {
			fmt.Printf("Erro ao analisar o template: %v\n", err)
			continue
		}

		// Criar o arquivo de saída
		outFile, err := os.Create(outFilePath)
		if err != nil {
			fmt.Printf("Erro ao criar o arquivo: %v\n", err)
			continue
		}

		// Executar o template e escrever no arquivo de saída
		err = tpl.Execute(outFile, module)
		if err != nil {
			fmt.Printf("Erro ao executar o template: %v\n", err)
		}
		outFile.Close()
	}
}
