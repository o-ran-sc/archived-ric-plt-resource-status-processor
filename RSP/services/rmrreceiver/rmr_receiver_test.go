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

package rmrreceiver

import (
	"fmt"
	"rsp/configuration"
	"rsp/logger"
	"rsp/managers"
	"rsp/managers/rmrmanagers"
	"rsp/mocks"
	"rsp/rmrcgo"
	"rsp/tests"
	"rsp/tests/testhelper"
	"testing"
	"time"
)

func TestListenAndHandle(t *testing.T) {
	rmrReceiver := initRmrReceiver(t)
	go rmrReceiver.ListenAndHandle()
	time.Sleep(time.Microsecond * 10)
}

func initRmrMessenger(log *logger.Logger) rmrcgo.RmrMessenger {
	rmrMessengerMock := &mocks.RmrMessengerMock{}
	rmrMessenger := rmrcgo.RmrMessenger(rmrMessengerMock)
	rmrMessengerMock.On("Init", tests.GetPort(), tests.MaxMsgSize, tests.Flags, log).Return(&rmrMessenger)

	// TODO: that's not good since we don't actually test anything. if the error is populated then the loop will just continue and it's sort of a "workaround" for that method to be called
	var buf *rmrcgo.MBuf
	e := fmt.Errorf("test error")
	rmrMessengerMock.On("RecvMsg").Return(buf, e)
	return rmrMessenger
}

func initRmrReceiver(t *testing.T) *RmrReceiver {
	rnibDataService, rmrSender, logger := testhelper.InitTestCase(t)
	config, err := configuration.ParseConfiguration()
	if err != nil {
		t.Errorf("#... - failed to parse configuration error: %s", err)
	}
	resourceStatusInitiateManager := managers.NewResourceStatusInitiateManager(logger, rnibDataService, rmrSender)

	rmrMessenger := initRmrMessenger(logger)
	manager := rmrmanagers.NewRmrMessageManager(logger, config, rnibDataService, rmrSender, resourceStatusInitiateManager, nil)

	return NewRmrReceiver(logger, rmrMessenger, manager)
}
