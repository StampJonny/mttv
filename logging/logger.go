package logging

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/stampjohnny/mttv/e"
	"github.com/stampjohnny/mttv/utils"
)

var BuyLog = "logs/buy"

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

func LogBuyContext(buyContext interface {
	GetAmount() float64
	GetMoneyBalance() float64
}) error {
	l, err := Get(BuyLog)
	if err != nil {
		return e.Err("can't get logger %s: %v", BuyLog, err)
	}

	amount := buyContext.GetAmount()
	Debug("amount=%v", amount)
	money := buyContext.GetMoneyBalance()
	Debug("money=%v", money)

	l.log.WithFields(Fields{
		"amount": amount,
		"money":  money,
	}).Info("buy")

	return nil
}

func (l *Logger) GetFilepath() string {
	return l.filepath
}

func (l *Logger) Log(message string) {
	l.log.Info(message)
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
	l.log.Formatter = new(logrus.TextFormatter)
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
