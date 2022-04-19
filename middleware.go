package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aerogo/aero"
)

/* Middleware that logs each request and it's response time */
func RequestLog(next aero.Handler) aero.Handler {
	return func(ctx aero.Context) error {
		start        := time.Now()
		err          := next(ctx)
		responseTime := time.Since(start)
		request      := ctx.Request()
		logString    := fmt.Sprintf(
			"%12s %15s %6s %s",
			responseTime,
			ctx.IP(),
			request.Method(),
			ctx.Path(),
		)

		log.Println(logString)

		return err
	}
}
