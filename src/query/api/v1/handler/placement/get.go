// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package placement

import (
	"net/http"
	"strconv"

	"github.com/m3db/m3/src/cmd/services/m3query/config"
	"github.com/m3db/m3/src/query/api/v1/handler"
	"github.com/m3db/m3/src/query/generated/proto/admin"
	"github.com/m3db/m3/src/query/util/logging"
	clusterclient "github.com/m3db/m3cluster/client"
	"github.com/m3db/m3cluster/placement"

	"go.uber.org/zap"
)

const (
	// OldM3DBGetURL is the old url for the placement get handler, maintained for
	// backwards compatibility.
	OldM3DBGetURL = handler.RoutePrefixV1 + "/placement"

	// M3DBGetURL is the url for the placement get handler (with the GET method).
	M3DBGetURL = handler.RoutePrefixV1 + "/m3db/services/placement"

	// GetHTTPMethod is the HTTP method used with this resource.
	GetHTTPMethod = http.MethodGet
)

// GetHandler is the handler for placement gets.
type GetHandler Handler

// NewGetHandler returns a new instance of GetHandler.
func NewGetHandler(client clusterclient.Client, cfg config.Configuration) *GetHandler {
	return &GetHandler{client: client, cfg: cfg}
}

func (h *GetHandler) ServeHTTP(serviceName string, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logging.WithContext(ctx)

	service, err := Service(
		h.client, NewServiceOptions(serviceName))
	if err != nil {
		handler.Error(w, err, http.StatusInternalServerError)
		return
	}

	var placement placement.Placement
	var version int
	status := http.StatusNotFound
	if vs := r.FormValue("version"); vs != "" {
		version, err = strconv.Atoi(vs)
		if err == nil {
			placement, err = service.PlacementForVersion(version)
		} else {
			status = http.StatusBadRequest
		}
	} else {
		placement, version, err = service.Placement()
	}

	if err != nil {
		handler.Error(w, err, status)
		return
	}

	placementProto, err := placement.Proto()
	if err != nil {
		logger.Error("unable to get placement protobuf", zap.Any("error", err))
		handler.Error(w, err, http.StatusInternalServerError)
		return
	}

	resp := &admin.PlacementGetResponse{
		Placement: placementProto,
		Version:   int32(version),
	}

	handler.WriteProtoMsgJSONResponse(w, resp, logger)
}
