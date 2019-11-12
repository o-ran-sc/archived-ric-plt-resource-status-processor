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
package rmrmsghandlers

import (
	"encoding/json"
	"gerrit.o-ran-sc.org/r/ric-plt/nodeb-rnib.git/entities"
	"rsp/configuration"
	"rsp/e2pdus"
	"rsp/logger"
	"rsp/managers"
	"rsp/models"
)

type ResourceStatusInitiateNotificationHandler struct {
	logger                        *logger.Logger
	config                        *configuration.Configuration
	resourceStatusInitiateManager managers.IResourceStatusInitiateManager
	requestName                   string
}

func NewResourceStatusInitiateNotificationHandler(logger *logger.Logger, config *configuration.Configuration, resourceStatusInitiateManager managers.IResourceStatusInitiateManager, requestName string) ResourceStatusInitiateNotificationHandler {
	return ResourceStatusInitiateNotificationHandler{
		logger:                        logger,
		config:                        config,
		resourceStatusInitiateManager: resourceStatusInitiateManager,
		requestName:                   requestName,
	}
}

func (h ResourceStatusInitiateNotificationHandler) Handle(request *models.RmrRequest) {
	inventoryName := request.RanName
	h.logger.Infof("#ResourceStatusInitiateNotificationHandler - RAN name: %s - Received %s notification", inventoryName, h.requestName)

	payload := models.ResourceStatusPayload{}
	err := json.Unmarshal(request.Payload, &payload)

	if err != nil {
		h.logger.Errorf("#ResourceStatusInitiateNotificationHandler - RAN name: %s - Error unmarshaling RMR request payload: %v", inventoryName, err)
		return
	}

	h.logger.Infof("#ResourceStatusInitiateNotificationHandler - Unmarshaled payload successfully: %+v", payload)

	if payload.NodeType != entities.Node_ENB {
		h.logger.Debugf("#ResourceStatusInitiateNotificationHandler - RAN name: %s, Node type isn't ENB", inventoryName)
		return
	}

	resourceStatusInitiateRequestParams := &e2pdus.ResourceStatusRequestData{}

	_ = h.resourceStatusInitiateManager.Execute(inventoryName, resourceStatusInitiateRequestParams)
}

