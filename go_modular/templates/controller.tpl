package {{ .NameLower }}

import (
    "github.com/gin-gonic/gin"
    service "{{ .Projeto }}/modules/{{ .NameLower }}/service"
)

type {{ .Name }}Controller struct {
    s service.{{ .Name }}ServiceInterface
}

func New{{ .Name }}Controller(service service.{{ .Name }}ServiceInterface) {{ .Name }}ControllerInterface {
    return &{{ .Name }}Controller{s: service}
}

func (uc *{{ .Name }}Controller) Get{{ .Name }}s(ctx *gin.Context) {
    // Implementação do método Get{{ .Name }}s
}

func (uc *{{ .Name }}Controller) Get{{ .Name }}(ctx *gin.Context) {
    // Implementação do método Get{{ .Name }}
}

func (uc *{{ .Name }}Controller) Create{{ .Name }}(ctx *gin.Context) {
    // Implementação do método Create{{ .Name }}
}

func (uc *{{ .Name }}Controller) Update{{ .Name }}(ctx *gin.Context) {
    // Implementação do método Update{{ .Name }}
}

func (uc *{{ .Name }}Controller) Delete{{ .Name }}(ctx *gin.Context) {
    // Implementação do método Delete{{ .Name }}
}
