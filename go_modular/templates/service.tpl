package {{ .NameLower }}

import (
    dto "{{ .Projeto }}/modules/{{ .NameLower }}/dto"
    model "{{ .Projeto }}/modules/{{ .NameLower }}/model"
    repository "{{ .Projeto }}/modules/{{ .NameLower }}/repository"
)

type {{ .Name }}Service struct {
    r repository.{{ .Name }}RepositoryInterface
}

func New{{ .Name }}Service(repository repository.{{ .Name }}RepositoryInterface) {{ .Name }}ServiceInterface {
    return &{{ .Name }}Service{r: repository}
}

func (us *{{ .Name }}Service) Get{{ .Name }}(id int) (user *model.{{ .Name }}, err error) {
    return nil, nil
}

func (us *{{ .Name }}Service) Get{{ .Name }}s() (users []*model.{{ .Name }}, err error) {
    return nil, nil
}

func (us *{{ .Name }}Service) Create{{ .Name }}(userDto dto.Create{{ .Name }}Dto) (user *model.{{ .Name }}, err error) {
    return nil, nil
}

func (us *{{ .Name }}Service) Update{{ .Name }}(id int, userDto dto.Update{{ .Name }}Dto) (rowsAffected int64, err error) {
    return 1, nil
}

func (us *{{ .Name }}Service) Delete{{ .Name }}(id int) (rowsAffected int64, err error) {
	return 1, nil
}