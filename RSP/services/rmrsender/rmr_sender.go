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

package rmrsender

import (
	"rsp/logger"
	"rsp/models"
	"rsp/rmrcgo"
)

type RmrSender struct {
	logger    *logger.Logger
	messenger rmrcgo.RmrMessenger
}

func NewRmrSender(logger *logger.Logger, messenger rmrcgo.RmrMessenger) *RmrSender {
	return &RmrSender{
		logger:    logger,
		messenger: messenger,
	}
}

func (r *RmrSender) Send(rmrMessage *models.RmrMessage) error {
	transactionIdByteArr := []byte(rmrMessage.RanName)
	msg := rmrcgo.NewMBuf(rmrMessage.MsgType, len(rmrMessage.Payload), rmrMessage.RanName, &rmrMessage.Payload, &transactionIdByteArr)

	_, err := r.messenger.SendMsg(msg)

	if err != nil {
		r.logger.Errorf("#RmrSender.Send - RAN name: %s , Message type: %d - Failed sending message. Error: %v", rmrMessage.RanName, rmrMessage.MsgType, err)
		return err
	}

	r.logger.Infof("#RmrSender.Send - RAN name: %s , Message type: %d - Successfully sent RMR message", rmrMessage.RanName, rmrMessage.MsgType)
	return nil
}