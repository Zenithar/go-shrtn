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

package cmd

import (
	"context"

	"github.com/cloudflare/tableflip"
	"github.com/oklog/run"
	"github.com/spf13/cobra"

	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/platform"
	"go.zenithar.org/shrtn/cli/shrtn/internal/dispatchers/graphql"
	"go.zenithar.org/shrtn/internal/version"
)

// -----------------------------------------------------------------------------

var graphqlCmd = &cobra.Command{
	Use:   "graphql",
	Short: "Starts the shrtn GraphQL server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Starting banner
		log.For(ctx).Info("Starting shrtn GraphQL server ...")

		// Start goroutine group
		err := platform.Serve(ctx, &platform.Server{
			Debug:           conf.Debug.Enable,
			Name:            "shrtn-graphql",
			Version:         version.Version,
			Revision:        version.Revision,
			Instrumentation: conf.Instrumentation,
			Builder: func(upg *tableflip.Upgrader, group *run.Group) {
				ln, err := upg.Fds.Listen(conf.Server.HTTP.Network, conf.Server.HTTP.Listen)
				if err != nil {
					log.For(ctx).Fatal("Unable to start GraphQL server", log.Error(err))
				}

				server, err := graphql.New(ctx, conf)
				if err != nil {
					log.For(ctx).Fatal("Unable to start GraphQL server", log.Error(err))
				}

				group.Add(
					func() error {
						log.For(ctx).Info("Starting GraphQL server", log.String("address", ln.Addr().String()))
						return server.Serve(ln)
					},
					func(e error) {
						log.For(ctx).Info("Shutting GraphQL server down")
						log.SafeClose(server, "Unable to close GraphQL server")
					},
				)
			},
		})
		log.CheckErrCtx(ctx, "Unable to run application", err)
	},
}
