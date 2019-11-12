//
// Copyright 2019 AT&T Intellectual Property
// Copyright 2019 Nokia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package logger

import (
	"fmt"
	"gerrit.o-ran-sc.org/r/com/golog"
	"strings"
	"time"
)

type Logger struct {
	Logger     *golog.MdcLogger
}


// A Level is a logging priority. Higher levels are more important.
type LogLevel int8

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel LogLevel = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel

	_maxLevel = ErrorLevel
)

var logLevelTokenToLevel = map[string] LogLevel {
	"debug" : DebugLevel,
	"info": InfoLevel,
	"warn": WarnLevel,
	"error": ErrorLevel,
}

func LogLevelTokenToLevel(level string) (LogLevel, bool) {
	if level, ok := logLevelTokenToLevel[strings.TrimSpace(strings.ToLower(level))];ok {
		return level, true
	}
	return _maxLevel+1, false
}

func InitLogger(requested LogLevel) (*Logger, error) {
	var logger *golog.MdcLogger
	var err error
	switch requested {
	case DebugLevel:
		logger, err = initLoggerByLevel(golog.DEBUG)
	case InfoLevel:
		logger, err = initLoggerByLevel(golog.INFO)
	case WarnLevel:
		logger, err = initLoggerByLevel(golog.WARN)
	case ErrorLevel:
		logger, err = initLoggerByLevel(golog.ERR)
	default:
		err = fmt.Errorf("invalid logging Level :%d",requested)
	}
	if err != nil {
		return nil, err
	}
	return &Logger{Logger:logger}, nil

}

func (l *Logger)Infof(formatMsg string, a ...interface{})  {
	if l.InfoEnabled() {
		l.Logger.MdcAdd("time", l.getTimeStampMdc())
		l.Logger.Info(formatMsg, a...)
	}
}

func (l *Logger)Debugf(formatMsg string, a ...interface{})  {
	if l.DebugEnabled(){
		l.Logger.MdcAdd("time", l.getTimeStampMdc())
		l.Logger.Debug(formatMsg, a...)
	}
}

func (l *Logger)Errorf(formatMsg string, a ...interface{})  {
	l.Logger.MdcAdd("time", l.getTimeStampMdc())
	l.Logger.Error(formatMsg, a...)
}

func (l *Logger)Warnf(formatMsg string, a ...interface{})  {
	l.Logger.MdcAdd("time", l.getTimeStampMdc())
	l.Logger.Warning(formatMsg, a...)
}

func (l *Logger) getTimeStampMdc() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}

func (l *Logger)InfoEnabled()bool{
	return l.Logger.LevelGet() > golog.WARN
}

func (l *Logger)DebugEnabled()bool{
	return l.Logger.LevelGet() > golog.INFO
}

func initLoggerByLevel(l golog.Level) (*golog.MdcLogger, error) {

	logger, err := golog.InitLogger("RSP")
	if logger != nil {
		logger.LevelSet(l)
	}
	return logger, err
}
