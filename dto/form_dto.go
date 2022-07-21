package dto

import "gin-demo/model"

type FormDto struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func ToFormDto(form model.Form) FormDto {
	return FormDto{
		Id:   form.Id,
		Name: form.Name,
	}
}
