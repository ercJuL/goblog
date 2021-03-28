package utils

import (
	"github.com/rs/zerolog/log"
)

func EntLog(iVal ...interface{}) {
	l := log.Debug()
	for i := 0; i < len(iVal); i++ {
		l.Interface(string(i), iVal[i])
	}
	l.Msg("ent log")
}
