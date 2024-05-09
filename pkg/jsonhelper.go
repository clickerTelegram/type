package pkg

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func Encode[T any](t T) []byte {
	b, err := json.Marshal(t)
	if err != nil {
		logrus.WithError(err).WithField("t", t).
			Fatalln("could`t encode the variable")
	}
	return b
}

func Decode[T any](b []byte) T {
	var t T
	err := json.Unmarshal(b, &t)
	if err != nil {
		logrus.WithError(err).WithField("t", t).
			Fatalln("could`t encode the variable")
	}
	return t
}
