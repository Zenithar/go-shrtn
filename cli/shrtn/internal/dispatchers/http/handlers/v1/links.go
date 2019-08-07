/*
 * Copyright 2019 Thibault NORMAND
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/web/respond"
	apiv1 "go.zenithar.org/shrtn/internal/services/pkg/link/v1"
	linkv1 "go.zenithar.org/shrtn/pkg/gen/go/shrtn/link/v1"
)

type linkCtrl struct {
	links apiv1.Link
}

// -----------------------------------------------------------------------------

// LinkRoutes is used to setup the http router.
func LinkRoutes(links apiv1.Link) http.Handler {
	r := chi.NewRouter()

	ctrl := &linkCtrl{
		links: links,
	}

	// Map routes
	r.Get("/{hash:[a-zA-Z0-9]+}", ctrl.redirect())

	return r
}

// -----------------------------------------------------------------------------

func (c *linkCtrl) redirect() http.HandlerFunc {
	// Handler
	return func(w http.ResponseWriter, r *http.Request) {
		// Prepare context
		ctx := r.Context()

		// Do the call
		res, err := c.links.Resolve(ctx, &linkv1.ResolveRequest{
			Hash: chi.URLParamFromCtx(ctx, "hash"),
		})
		if err != nil {
			log.For(ctx).Error("Unable to query database", log.Error(err))
			respond.WithError(w, r, http.StatusInternalServerError, "Unable to query advisory database")
			return
		}

		// Send redirect to client
		http.Redirect(w, r, res.Link.Url, http.StatusMovedPermanently)
	}
}
