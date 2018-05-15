//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package cluster

import (
	"github.com/lastbackend/lastbackend/pkg/api/types/v1"
	"github.com/lastbackend/lastbackend/pkg/distribution/errors"
	"github.com/lastbackend/lastbackend/pkg/log"
	"net/http"
	"github.com/lastbackend/lastbackend/pkg/distribution/types"
)

const (
	logLevel  = 2
	logPrefix = "api:handler:cluster"
)

func ClusterInfoH(w http.ResponseWriter, r *http.Request) {

	// swagger:operation GET /cluster cluster clusterInfo
	//
	// Shows an info about current cluster
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: Cluster response
	//     schema:
	//       "$ref": "#/definitions/views_cluster"
	//   '404':
	//     description: Cluster not found
	//   '500':
	//     description: Internal server error

	log.V(logLevel).Debugf("%s:info:> get cluster", logPrefix)

	cl := new(types.Cluster)
	cl.Meta.Name = types.DefaultClusterName
	cl.Meta.Description = types.DefaultClusterDescription
	cl.SelfLink()

	response, err := v1.View().Cluster().New(cl).ToJson()
	if err != nil {
		log.V(logLevel).Errorf("%s:info:> convert struct to json err: %s", logPrefix, err.Error())
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.V(logLevel).Errorf("%s:info:> write response err: %s", logPrefix, err.Error())
		return
	}
}
