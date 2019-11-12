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

package rmrmsghandlerprovider

import (
	"fmt"
	"github.com/pkg/errors"
	"rsp/configuration"
	"rsp/converters"
	"rsp/handlers/rmrmsghandlers"
	"rsp/logger"
	"rsp/managers"
	"rsp/rmrcgo"
	"rsp/services"
	"rsp/services/rmrsender"
)

const (
	RanConnected string = "Ran connected"
	RanRestarted string = "Ran restarted"
)

type MessageHandlerProvider struct {
	msgHandlers map[int]rmrmsghandlers.RmrMessageHandler
}

func NewMessageHandlerProvider(logger *logger.Logger, config *configuration.Configuration, rnibDataService services.RNibDataService, rmrSender *rmrsender.RmrSender, resourceStatusInitiateManager *managers.ResourceStatusInitiateManager, unpacker converters.Asn1PduUnpacker) *MessageHandlerProvider {
	return &MessageHandlerProvider{
		msgHandlers: initMessageHandlersMap(logger, config, rnibDataService, rmrSender, resourceStatusInitiateManager, unpacker),
	}
}

func initMessageHandlersMap(logger *logger.Logger, config *configuration.Configuration, rnibDataService services.RNibDataService, rmrSender *rmrsender.RmrSender, resourceStatusInitiateManager *managers.ResourceStatusInitiateManager, unpacker converters.Asn1PduUnpacker) map[int]rmrmsghandlers.RmrMessageHandler {
	return map[int]rmrmsghandlers.RmrMessageHandler{
		rmrcgo.RanConnected:        rmrmsghandlers.NewResourceStatusInitiateNotificationHandler(logger, config, resourceStatusInitiateManager, RanConnected),
		rmrcgo.RanRestarted:        rmrmsghandlers.NewResourceStatusInitiateNotificationHandler(logger, config, resourceStatusInitiateManager, RanRestarted),
		rmrcgo.RicResStatusFailure: rmrmsghandlers.NewResourceStatusFailureHandler(logger, unpacker),
		rmrcgo.RicResStatusResp:    rmrmsghandlers.NewResourceStatusResponseHandler(logger),
	}
}

func (provider MessageHandlerProvider) GetMessageHandler(messageType int) (rmrmsghandlers.RmrMessageHandler, error) {
	handler, ok := provider.msgHandlers[messageType]

	if !ok {
		msg := fmt.Sprintf("#MessageHandlerProvider.GetMessageHandler - notification handler not found for message %d",messageType )
		return nil, errors.New(msg)
	}

	return handler, nil

}
