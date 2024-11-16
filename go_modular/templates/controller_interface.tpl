package {{ .NameLower }}

import (
    "github.com/gin-gonic/gin"
)

type {{ .Name }}ControllerInterface interface {
    Get{{ .Name }}s(ctx *gin.Context)
    Get{{ .Name }}(ctx *gin.Context)
    Create{{ .Name }}(ctx *gin.Context)
    Update{{ .Name }}(ctx *gin.Context)
    Delete{{ .Name }}(ctx *gin.Context)
}