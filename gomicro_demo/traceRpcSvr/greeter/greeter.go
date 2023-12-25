package greeter

import (
	"context"
	"fmt"

	"github.com/go-micro/cli/debug/trace"
)

func Greet(ctx context.Context, name string) string {
	defer trace.NewSpan(ctx).Finish()
	return fmt.Sprint("Hello " + name)
}
