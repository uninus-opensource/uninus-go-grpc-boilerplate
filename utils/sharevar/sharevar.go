package sharevar

import (
	"context"

	logkit "github.com/go-kit/kit/log"
	"go.elastic.co/apm"
)

// Logger is global variable for logging
var Logger logkit.Logger

// Tracer APM
var TracerAPM *apm.Tracer

// Context : fill context with auth to call
var Context context.Context
