package {{ .NameLower }}

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    "{{ .Projeto }}/modules"
    controller "{{ .Projeto }}/modules/{{ .NameLower }}/controller"
    repository "{{ .Projeto }}/modules/{{ .NameLower }}/repository"
    service "{{ .Projeto }}/modules/{{ .NameLower }}/service"
    // m "{{ .Projeto }}/go-task-manager/pkg/middlewares"
)

type {{ .Name }}Module struct{}

func New{{ .Name }}Module() modules.ModuleInterface {
    return &{{ .Name }}Module{}
}

func (um *{{ .Name }}Module) Init(router *gin.RouterGroup, db *sql.DB) {
    {{.NameLower}}Repository := repository.New{{ .Name }}Repository(db)
    {{.NameLower}}Service := service.New{{ .Name }}Service({{.NameLower}}Repository)
    {{.NameLower}}Controller := controller.New{{ .Name }}Controller({{.NameLower}}Service)

    // middlewares examples
    // router.GET("/{{ .NameLower }}s", m.Logger(m.Authenticate({{.NameLower}}Controller.Get{{ .Name }}s)))
    
    router.GET("/{{ .NameLower }}s", {{.NameLower}}Controller.Get{{ .Name }}s)
    router.GET("/{{ .NameLower }}/:id", {{.NameLower}}Controller.Get{{ .Name }})
    router.POST("/{{ .NameLower }}", {{.NameLower}}Controller.Create{{ .Name }})
    router.PUT("/{{ .NameLower }}/:id", {{.NameLower}}Controller.Update{{ .Name }})
    router.DELETE("/{{ .NameLower }}/:id", {{.NameLower}}Controller.Delete{{ .Name }})
}
