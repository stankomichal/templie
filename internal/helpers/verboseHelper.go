package helpers

import (
	"context"
	"fmt"
	"github.com/stankomichal/templie/internal/contextKey"
)

func IsVerbose(ctx context.Context) bool {
	verbose, ok := ctx.Value(contextKey.VerboseKey).(bool)
	if !ok {
		return false
	}
	return verbose
}

func VerbosePrintln(ctx context.Context, message string) {
	if IsVerbose(ctx) {
		fmt.Println(message)
	}
}

func VerbosePrintf(ctx context.Context, format string, args ...interface{}) {
	if IsVerbose(ctx) {
		fmt.Printf(format, args...)
	}
}
