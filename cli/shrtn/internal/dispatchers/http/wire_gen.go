// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package http

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/plugin/ochttp/propagation/b3"
	"go.opencensus.io/stats/view"
	"go.zenithar.org/pkg/log"
	"go.zenithar.org/pkg/tlsconfig"
	"go.zenithar.org/shrtn/cli/shrtn/internal/config"
	"go.zenithar.org/shrtn/cli/shrtn/internal/core"
	v1_2 "go.zenithar.org/shrtn/cli/shrtn/internal/dispatchers/http/handlers/v1"
	"go.zenithar.org/shrtn/internal/repositories/pkg/badger"
	v1 "go.zenithar.org/shrtn/internal/services/pkg/link/v1"
)

// Injectors from wire.go:

func setup(ctx context.Context, cfg *config.Configuration) (*http.Server, error) {
	db, err := core.BadgerConfig(cfg)
	if err != nil {
		return nil, err
	}
	link := badger.Link(db)
	hasher := core.Hasher(cfg)
	v1Link := v1.New(link, hasher)
	server, err := httpServer(ctx, cfg, v1Link)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// wire.go:

func httpServer(ctx context.Context, cfg *config.Configuration, links v1.Link) (*http.Server, error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/", func(r chi.Router) {
		r.Mount("/", ochttp.WithRouteTag(v1_2.LinkRoutes(links), "/"))
	})

	server := &http.Server{
		Handler: &ochttp.Handler{
			Handler:     r,
			Propagation: &b3.HTTPFormat{},
		},
	}

	if cfg.Server.HTTP.UseTLS {

		clientAuth := tls.VerifyClientCertIfGiven
		if cfg.Server.HTTP.TLS.ClientAuthenticationRequired {
			clientAuth = tls.RequireAndVerifyClientCert
		}

		tlsConfig, err := tlsconfig.Server(tlsconfig.Options{
			KeyFile:    cfg.Server.HTTP.TLS.PrivateKeyPath,
			CertFile:   cfg.Server.HTTP.TLS.CertificatePath,
			CAFile:     cfg.Server.HTTP.TLS.CACertificatePath,
			ClientAuth: clientAuth,
		})
		if err != nil {
			log.For(ctx).Error("Unable to build TLS configuration from settings", log.Error(err))
			return nil, err
		}

		server.TLSConfig = tlsConfig
	} else {
		log.For(ctx).Info("No transport encryption enabled for HTTP server")
	}

	err := view.Register(ochttp.DefaultServerViews...,
	)
	if err != nil {
		log.For(ctx).Fatal("Unable to register stat views", log.Error(err))
	}

	return server, nil
}
