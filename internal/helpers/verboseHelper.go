package helpers

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/stankomichal/templie/internal/contextKey"
)

func IsVerbose(ctx context.Context) bool {
	verbose, ok := ctx.Value(contextKey.VerboseKey).(bool)
	if !ok {
		return false
	}
	return verbose
}

func VerbosePrintln(cmd *cobra.Command, ctx context.Context, message string) {
	if IsVerbose(ctx) {
		cmd.Println(message)
	}
}

func VerbosePrintf(cmd *cobra.Command, ctx context.Context, format string, args ...interface{}) {
	if IsVerbose(ctx) {
		cmd.Printf(format, args...)
	}
}
