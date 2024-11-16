package {{ .NameLower }}

import (
   dto "{{ .Projeto }}/modules/{{ .NameLower }}/dto"
   model "{{ .Projeto}}/modules/{{ .NameLower }}/model"
)

type {{ .Name }}ServiceInterface interface {
    Get{{ .Name }}(id int) (user *model.{{ .Name }}, err error)
    Get{{ .Name }}s() (users []*model.{{ .Name }}, err error)
    Create{{ .Name }}(userDto dto.Create{{ .Name }}Dto) (user *model.{{ .Name }}, err error)
    Update{{ .Name }}(id int, userDto dto.Update{{ .Name }}Dto) (rowsAffected int64, err error)
    Delete{{ .Name }}(id int) (rowsAffected int64, err error)
}