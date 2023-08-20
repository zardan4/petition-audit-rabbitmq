package errhand

import "github.com/sirupsen/logrus"

func FailOnError(err error, msg string) {
	if err != nil {
		logrus.Panicf("%s: %s", msg, err)
	}
}
