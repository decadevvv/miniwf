package utils

import (
	validator "github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

func ValidateStruct(s interface{}) error {
	valid := validator.New()
	if err := valid.Struct(s); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			log.Errorf("validation error: Namespace = %s, StructNamespace = %s, StructField = %s, Field = %s, Tag = %s, Kind = %#v, Type = %#v, Value = %#v, Param = %s, Error = %s",
				e.Namespace(), e.StructNamespace(), e.StructField(), e.Field(), e.Tag(), e.Kind(), e.Type(), e.Value(), e.Param(), e.Error())
		}
		return err
	}
	return nil
}
