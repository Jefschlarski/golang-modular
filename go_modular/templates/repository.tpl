package {{ .NameLower }}

import (
    "database/sql"
    model "{{ .Projeto }}/modules/{{ .NameLower }}/model"
)

type {{ .Name }}Repository struct {
    db *sql.DB
}

func New{{ .Name }}Repository(db *sql.DB) {{ .Name }}RepositoryInterface {
    return &{{ .Name }}Repository{db: db}
}

func (ur *{{ .Name }}Repository) Get{{ .Name }}(id int) (user *model.{{ .Name }}, err error) {
    return nil, nil
}

func (ur *{{ .Name }}Repository) Get{{ .Name }}s() (users []*model.{{ .Name }}, err error) {
    return nil, nil
}

func (ur *{{ .Name }}Repository) Create{{ .Name }}(create{{ .Name }} model.{{ .Name }}) (userId int, err error) {
    return 0, nil
}

func (ur *{{ .Name }}Repository) Update{{ .Name }}(id int, user model.{{ .Name }}) (rowsAffected int64, err error) {
    return 0, nil
}

func (ur *{{ .Name }}Repository) Delete{{ .Name }}(id int) (rowsAffected int64, err error) {
    return 0, nil
}
