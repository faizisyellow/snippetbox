package config

import "log"

type Aplication struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}
