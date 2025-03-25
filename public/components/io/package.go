// Copyright 2025 Redpanda Data, Inc.

// Package io contains component implementations that have a small dependency
// footprint (mostly standard library) and interact with external systems via
// the filesystem and/or network sockets.
//
// EXPERIMENTAL: The specific components excluded by this package may change
// outside of major version releases. This means we may choose to remove certain
// plugins if we determine that their dependencies are likely to interfere with
// the goals of this package.
package io

import (
	"github.com/redpanda-data/benthos/v4/internal/impl/io"
	"github.com/redpanda-data/benthos/v4/public/service"
)

// HTTPInputMiddlewareMetaConst is a public type that is used to construct the http input middleware with optional config fields.
type HTTPInputMiddlewareMetaConst = io.HTTPInputMiddlewareMetaConst

// HTTPInputMiddlewareMeta is a public type that is used to register custom middleware for adding metadata to a message.
type HTTPInputMiddlewareMeta = io.HTTPInputMiddlewareMeta

// RegisterCustomHTTPServerInput registers a custom HTTP server input with a given name and optional middleware.
func RegisterCustomHTTPServerInput(name string, middlewareConst HTTPInputMiddlewareMetaConst, conf *service.ConfigField) {
	io.RegisterCustomHTTPServerInput(name, middlewareConst, conf)
}
