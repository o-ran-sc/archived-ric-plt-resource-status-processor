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
	"rsp/logger"
	"rsp/managers/rmrmanagers"
	"rsp/rmrcgo"
)

type RmrReceiver struct {
	logger    *logger.Logger
	nManager  *rmrmanagers.RmrMessageManager
	messenger rmrcgo.RmrMessenger
}

func NewRmrReceiver(logger *logger.Logger, messenger rmrcgo.RmrMessenger, nManager *rmrmanagers.RmrMessageManager) *RmrReceiver {
	return &RmrReceiver{
		logger:    logger,
		nManager:  nManager,
		messenger: messenger,
	}
}

func (r *RmrReceiver) ListenAndHandle() {

	for {
		mbuf, err := r.messenger.RecvMsg()
		r.logger.Debugf("#RmrReceiver.ListenAndHandle - Going to handle received message: %#v\n", mbuf)

		if err != nil {
			// TODO: error handling?
			continue
		}

		// TODO: go routine?
		r.nManager.HandleMessage(mbuf)
	}
}