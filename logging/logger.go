package logging

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/stampjohnny/mttv/utils"
)

var TradingLog = "logs/trading"

type Logger struct {
	filepath string
	log      *logrus.Logger
}

var registry = map[string]*Logger{}

type Fields = logrus.Fields

func Get(name string) (*Logger, error) {
	l, ok := registry[name]
	if !ok {
		l = create(name)
		registry[name] = l
	}

	return l, nil
}

func (l *Logger) GetFilepath() string {
	return l.filepath
}

func (l *Logger) Log(fields Fields, message string) {
	l.log.WithFields(fields).Info(message)
}
func (l *Logger) Remove() error {
	return os.RemoveAll(l.filepath)
}

func (l *Logger) Init() error {
	f, err := os.OpenFile(l.filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	l.log = logrus.New()
	l.log.Formatter = new(logrus.JSONFormatter)
	l.log.Level = logrus.DebugLevel
	l.log.Out = f

	return nil
}

func create(name string) *Logger {
	filePath := filepath.Join("/mttv/", name)
	l := &Logger{filepath: filePath}
	if err := l.Init(); err != nil {
		utils.Crash("error while initializing looger %s: %s", name, err)
	}
	return l
}
