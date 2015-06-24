package typetalk

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"net/http"
	"net/url"
)

const (
	VERISON = "1.0.0"
)

type TypetalkHook struct {
	AcceptedLevels []logrus.Level
	BotURL         string
}

func (th *TypetalkHook) Levels() []logrus.Level {
	if th.AcceptedLevels == nil {
		return AllLevels
	}
	return th.AcceptedLevels
}

func (th *TypetalkHook) Fire(e *logrus.Entry) error {
	resp, err := http.PostForm(th.BotURL, url.Values{"message": {e.Message}})
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(e.Message)
	}
	return nil
}
