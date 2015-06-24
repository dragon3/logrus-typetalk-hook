logrus-typetalk-hook
========

[Typetalk](https://typetalk.in) hook for [Logrus](https://github.com/Sirupsen/logrus). 

## Use

```go
package main

import (
	logrus "github.com/Sirupsen/logrus"
	"github.com/dragon3/logrus-typetalk-hook"
)

func main() {
	logrus.AddHook(&typetalk.TypetalkHook{
		BotURL:        "https://typetalk.in/api/v1/topics/000?typetalkToken=abcdefghijklmnopqrstuvwxyz",
		AcceptedLevels: typetalk.LevelThreshold(logrus.DebugLevel),
	})

	logrus.Warn("warn")
	logrus.Info("info")
	logrus.Debug("debug")
}

```

## Parameters

#### Required
  * ```BotURL``` : 
    You can add a Bot from "Edit topic" and get BotURL from "Get or post messages URL" ( see: https://nulab-inc.com/blog/typetalk/typetalk-update-create-bots-easily-use-webhook/ )
  
## Installation

    go get github.com/dragon3/logrus-typetalk-hook
    
## Credits 

Based on hipchat handler by [nuboLAB](https://github.com/nubo/hiprus) and slack handler by [John Dyer](https://github.com/johntdyer/slackrus)
