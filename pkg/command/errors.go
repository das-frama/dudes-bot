package command

import "errors"

// ErrWrongDateFormat is used when file doesn't have an .md extension.
var ErrWrongDateFormat = errors.New("ой простити-извините, вы не совсем верно написали дату. Ой, пожалуйста, укажите правильный формат dd.mm.yyyy или dd.mm")
