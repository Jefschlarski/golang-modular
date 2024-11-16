package {{ .NameLower }}

import (
    model "{{ .Projeto }}/modules/{{ .NameLower }}/model"
)

type {{ .Name }}RepositoryInterface interface {
    Get{{ .Name }}(id int) (user *model.{{ .Name }}, err error)
    Get{{ .Name }}s() (users []*model.{{ .Name }}, err error)
    Create{{ .Name }}(create{{ .Name }} model.{{ .Name }}) (userId int, err error)
    Update{{ .Name }}(id int, user model.{{ .Name }}) (rowsAffected int64, err error)
    Delete{{ .Name }}(id int) (rowsAffected int64, err error)
}
