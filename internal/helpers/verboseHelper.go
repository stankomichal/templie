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

func VerbosePrintln(cmd *cobra.Command, message string) {
	if IsVerbose(cmd.Context()) {
		cmd.Println(message)
	}
}

func VerbosePrintf(cmd *cobra.Command, format string, args ...interface{}) {
	if IsVerbose(cmd.Context()) {
		cmd.Printf(format, args...)
	}
}
