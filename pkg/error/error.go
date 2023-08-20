package errhand

import "github.com/sirupsen/logrus"

func ErrorOnError(err error, msg string) {
	if err != nil {
		logrus.Errorf("%s: %s", msg, err)
	}
}
