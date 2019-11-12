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
	"bytes"
	"gerrit.o-ran-sc.org/r/com/golog"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestInitDebugLoggerSuccess(t *testing.T) {
	log, err := InitLogger(DebugLevel)
	assert.Nil(t, err)
	assert.NotNil(t, log)
	assert.True(t, log.Logger.LevelGet() == golog.DEBUG)
}

func TestInitInfoLoggerSuccess(t *testing.T) {
	log, err := InitLogger(InfoLevel)
	assert.Nil(t, err)
	assert.NotNil(t, log)
	assert.True(t, log.Logger.LevelGet() == golog.INFO)
}

func TestInitWarnLoggerSuccess(t *testing.T) {
	log, err := InitLogger(WarnLevel)
	assert.Nil(t, err)
	assert.NotNil(t, log)
	assert.True(t, log.Logger.LevelGet() == golog.WARN)
}

func TestInitErrorLoggerSuccess(t *testing.T) {
	log, err := InitLogger(ErrorLevel)
	assert.Nil(t, err)
	assert.NotNil(t, log)
	assert.True(t, log.Logger.LevelGet() == golog.ERR)
}

func TestInitInfoLoggerFailure(t *testing.T) {
	log, err := InitLogger(99)
	assert.NotNil(t, err)
	assert.Nil(t, log)
}

func TestDebugEnabledFalse(t *testing.T){
	entryNum, log := countRecords(InfoLevel, t)
	assert.False(t, log.DebugEnabled())
	assert.Equal(t,2, entryNum)
}

func TestDebugEnabledTrue(t *testing.T){
	entryNum, log := countRecords(DebugLevel, t)
	assert.True(t, log.DebugEnabled())
	assert.Equal(t,3, entryNum)
}

func TestErrorfDebugLevel(t *testing.T)  {
	assert.True(t,validateRecordExists(DebugLevel, golog.ERR, t))
}

func TestErrorfInfoLevel(t *testing.T)  {
	assert.True(t,validateRecordExists(InfoLevel, golog.ERR, t))
}

func TestInfofDebugLevel(t *testing.T)  {
	assert.True(t,validateRecordExists(DebugLevel, golog.INFO, t))
}

func TestInfofInfoLevel(t *testing.T)  {
	assert.True(t,validateRecordExists(InfoLevel, golog.INFO, t))
}

func TestDebugfDebugLevel(t *testing.T)  {
	assert.True(t,validateRecordExists(DebugLevel, golog.DEBUG, t))
}

func TestDebugfInfoLevel(t *testing.T)  {
	assert.False(t,validateRecordExists(InfoLevel, golog.DEBUG, t))
}

func TestWarnfWarnLevel(t *testing.T)  {
	assert.True(t,validateRecordExists(WarnLevel, golog.WARN, t))
}

func TestWarnfDebugLevel(t *testing.T)  {
	assert.True(t,validateRecordExists(DebugLevel, golog.WARN, t))
}

func TestWarnfInfoLevel(t *testing.T)  {
	assert.True(t,validateRecordExists(InfoLevel, golog.WARN, t))
}

func TestLogLevelTokenToLevel(t *testing.T) {
	level, ok := LogLevelTokenToLevel("deBug")
	assert.True(t, ok)
	assert.True(t, level == DebugLevel)

	level, ok = LogLevelTokenToLevel("infO")
	assert.True(t, ok)
	assert.True(t, level == InfoLevel)

	level, ok = LogLevelTokenToLevel("Warn")
	assert.True(t, ok)
	assert.True(t, level == WarnLevel)

	level, ok = LogLevelTokenToLevel("eRror")
	assert.True(t, ok)
	assert.True(t, level == ErrorLevel)

	level, ok = LogLevelTokenToLevel("zzz")
	assert.False(t, ok)
	assert.True(t, level > ErrorLevel)

}
func countRecords(logLevel LogLevel, t *testing.T) (int, *Logger){
	old := os.Stdout
	r, w, _ :=os.Pipe()
	os.Stdout = w
	log, err := InitLogger(logLevel)
	if err != nil {
		t.Errorf("logger_test.TestSyncFailure - failed to initialize logger, error: %s", err)
	}
	log.Infof("%v, %v, %v", 1, "abc", 0.1)
	log.Debugf("%v, %v, %v", 1, "abc", 0.1)
	log.Errorf("%v, %v, %v", 1, "abc", 0.1)

	err = w.Close()
	if err != nil {
		t.Errorf("logger_test.TestSyncFailure - failed to close writer, error: %s", err)
	}
	os.Stdout = old
	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Errorf("logger_test.TestSyncFailure - failed to copy bytes, error: %s", err)
	}
	entryNum := 0
	s,_:= buf.ReadString('\n')
	for len(s) > 0{
		entryNum +=1
		s,_= buf.ReadString('\n')
	}
	return entryNum, log
}

func validateRecordExists(logLevel LogLevel, recordLevel golog.Level, t *testing.T) bool {
	old := os.Stdout
	r, w, _ :=os.Pipe()
	os.Stdout = w
	log, err := InitLogger(logLevel)
	if err != nil {
		t.Errorf("logger_test.TestSyncFailure - failed to initialize logger, error: %s", err)
	}
	switch recordLevel{
	case  golog.DEBUG:
		log.Debugf("%v, %v, %v", 1, "abc", 0.1)
	case golog.INFO:
		log.Infof("%v, %v, %v", 1, "abc", 0.1)
	case golog.WARN:
		log.Warnf("%v, %v, %v", 1, "abc", 0.1)
	case golog.ERR:
		log.Errorf("%v, %v, %v", 1, "abc", 0.1)
	}
	err = w.Close()
	if err != nil {
		t.Errorf("logger_test.TestSyncFailure - failed to close writer, error: %s", err)
	}
	os.Stdout = old
	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		t.Errorf("logger_test.TestSyncFailure - failed to copy bytes, error: %s", err)
	}
	entryNum := 0
	s,_:= buf.ReadString('\n')
	for len(s) > 0{
		entryNum +=1
		s,_= buf.ReadString('\n')
	}
	return entryNum == 1
}