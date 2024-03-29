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

package rmrcgo

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"rsp/logger"
	"rsp/tests"
	"strconv"
	"testing"
)

var (
	msgr RmrMessenger
)

func TestLogger(t *testing.T) {
	log := initLog(t)
	data := map[string]interface{}{"messageType": 1001, "ranIp": "10.0.0.3", "ranPort": 879, "ranName": "test1"}
	b := new(bytes.Buffer)
	_ = json.NewEncoder(b).Encode(data)
	req := tests.GetHttpRequest()
	boo, _ := ioutil.ReadAll(req.Body)
	log.Debugf("#rmr_c_go_api_test.TestLogger - request header: %v\n; request body: %s\n", req.Header, string(boo))
}

func TestNewMBufSuccess(t *testing.T) {
	msg := NewMBuf(tests.MessageType, len(tests.DummyPayload), "RanName", &tests.DummyPayload, &tests.DummyXAction)
	assert.NotNil(t, msg)
	assert.NotEmpty(t, msg.Payload)
	assert.NotEmpty(t, msg.XAction)
	assert.Equal(t, msg.MType, tests.MessageType)
	assert.Equal(t, msg.Meid, "RanName")
	assert.Equal(t, msg.Len, len(tests.DummyPayload))
}

func TestSendRecvMsgSuccess(t *testing.T) {
	log := initLog(t)

	initRmr(tests.ReadyIntervalSec, tests.GetPort(), tests.MaxMsgSize, tests.Flags, log)
	if msgr == nil || !msgr.IsReady()  {
		t.Errorf("#rmr_c_go_api_test.TestSendRecvMsgSuccess - The rmr router is not ready")
	}

	msg := NewMBuf(1, tests.MaxMsgSize, "test 1", &tests.DummyPayload, &tests.DummyXAction)
	log.Debugf("#rmr_c_go_api_test.TestSendRecvMsgSuccess - Going to send the message: %#v\n", msg)
	result, err := msgr.SendMsg(msg)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	msgR, err := msgr.RecvMsg()

	assert.Nil(t, err)
	assert.NotNil(t, msgR)
	msgr.Close()
}

func TestSendMsgRmrInvalidPortError(t *testing.T) {
	log := initLog(t)

	initRmr(tests.ReadyIntervalSec, "tcp:" + strconv.Itoa(5555), tests.MaxMsgSize, tests.Flags, log)
	if msgr == nil || !msgr.IsReady()  {
		t.Errorf("#rmr_c_go_api_test.TestSendMsgRmrInvalidPortError - The rmr router is not ready")
	}

	msg := NewMBuf(1, tests.MaxMsgSize, "test 1", &tests.DummyPayload, &tests.DummyXAction)
	log.Debugf("#rmr_c_go_api_test.TestSendMsgRmrInvalidPortError - Going to send the message: %#v\n", msg)
	result, err := msgr.SendMsg(msg)

	assert.NotNil(t, err)
	assert.Nil(t, result)

	msgr.Close()
}

func TestSendMsgRmrInvalidMsgNumError(t *testing.T) {
	log := initLog(t)

	initRmr(tests.ReadyIntervalSec, tests.GetPort(), tests.MaxMsgSize, tests.Flags, log)
	if msgr == nil || !msgr.IsReady()  {
		t.Errorf("#rmr_c_go_api_test.TestSendMsgRmrInvalidMsgNumError - The rmr router is not ready")
	}

	msg := NewMBuf(10, tests.MaxMsgSize, "test 1", &tests.DummyPayload, &tests.DummyXAction)
	log.Debugf("#rmr_c_go_api_test.TestSendMsgRmrInvalidMsgNumError - Going to send the message: %#v\n", msg)
	result, err := msgr.SendMsg(msg)

	assert.NotNil(t, err)
	assert.Nil(t, result)

	msgr.Close()
}

func TestIsReadySuccess(t *testing.T) {
	log := initLog(t)

	initRmr(tests.ReadyIntervalSec, tests.GetPort(), tests.MaxMsgSize, tests.Flags, log)
	if msgr == nil || !msgr.IsReady()  {
		t.Errorf("#rmr_c_go_api_test.TestIsReadySuccess - The rmr router is not ready")
	}

	msgr.Close()
}

func initRmr(readyIntervalSec int, port string, maxMsgSize int, flags int, log *logger.Logger){
	var ctx *Context
	msgr = ctx.Init(readyIntervalSec, port, maxMsgSize, flags, log)
}

func initLog(t *testing.T) *logger.Logger {
	log, err := logger.InitLogger(logger.DebugLevel)
	if err != nil {
		t.Errorf("#rmr_c_go_api_test.initLog - failed to initialize logger, error: %s", err)
	}
	return log
}