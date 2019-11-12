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

package httpserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"rsp/controllers"
)

func Run(port int, controller controllers.IRootController) error {

	router := mux.NewRouter()
	initializeRoutes(router, controller)

	addr := fmt.Sprintf(":%d", port)

	err := http.ListenAndServe(addr, router)

	return fmt.Errorf("#http_server.Run - Fail initiating HTTP server. Error: %v", err)
}

func initializeRoutes(router *mux.Router, rootController controllers.IRootController) {
	r := router.PathPrefix("/v1").Subrouter()
	r.HandleFunc("/health", rootController.HandleHealthCheckRequest).Methods("GET")
}
