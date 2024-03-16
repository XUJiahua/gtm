package gtm

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestInvalidCursorCheck(t *testing.T) {
	positionLost := mongo.CommandError{Code: 136}
	assert.True(t, invalidCursor(positionLost))
	err := mongo.CommandError{Code: 999}
	assert.False(t, invalidCursor(err))
	assert.False(t, invalidCursor(nil))
}

func tm(tStr string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", tStr)
	if err != nil {
		panic(err)
	}
	return t
}

func TestCheckOrder(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	times := []time.Time{
		tm("2023-11-23 02:18:54"),
		tm("2023-09-22 13:47:50"),
		tm("2023-11-23 02:18:54"),
		tm("2023-07-28 13:26:06"),
		tm("2023-10-27 13:15:54"),
		tm("2023-07-28 13:26:07"),
	}

	lastTs := time.Time{}
	for _, ts := range times {
		logrus.Debugf("ts: %s", ts.String())
		if ts.Before(lastTs) {
			logrus.Debugf("ts/_id is not monotonic increasing")
			logrus.Warnf("not work for DirectReadResumable mode")
		}
		lastTs = ts
	}
}
