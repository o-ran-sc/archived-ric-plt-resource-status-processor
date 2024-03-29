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

package controllers

import (
	"net/http"
	"rsp/services"
)

type IRootController interface {
	HandleHealthCheckRequest(writer http.ResponseWriter, request *http.Request)
}

type RootController struct {
	rnibDataService services.RNibDataService
}

func NewRootController(rnibDataService services.RNibDataService) *RootController {
	return &RootController{
		rnibDataService: rnibDataService,
	}
}

func (rc *RootController) HandleHealthCheckRequest(writer http.ResponseWriter, request *http.Request) {
	httpStatus := http.StatusNoContent
	//TODO: use RSP data-service
	isOn := rc.rnibDataService.PingRnib()
	if !isOn {
		httpStatus = http.StatusInternalServerError
	}

	writer.WriteHeader(httpStatus)
}
