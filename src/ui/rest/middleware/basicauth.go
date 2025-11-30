package middleware

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/aldinokemal/go-whatsapp-web-multidevice/config"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthorizationValue string

func BasicAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := string(c.Request().Header.Peek("Authorization"))
		if token != "" {
			ctx := context.WithValue(c.Context(), AuthorizationValue("BASIC_AUTH"), token)
			c.SetUserContext(ctx)

			// Debug logging for authentication
			if config.AppDebug {
				logrus.WithFields(logrus.Fields{
					"path":              c.Path(),
					"method":            c.Method(),
					"authorization":     token,
					"remote_ip":         c.IP(),
				}).Debug("[AUTH] Authorization header received")

				// Try to decode and show credentials (for debugging)
				if strings.HasPrefix(token, "Basic ") {
					encodedCreds := strings.TrimPrefix(token, "Basic ")
					if decoded, err := base64.StdEncoding.DecodeString(encodedCreds); err == nil {
						creds := strings.SplitN(string(decoded), ":", 2)
						if len(creds) == 2 {
							logrus.WithFields(logrus.Fields{
								"username":     creds[0],
								"password_len": len(creds[1]),
								"hash_sent":    encodedCreds,
							}).Debug("[AUTH] Decoded credentials")
						}
					} else {
						logrus.WithError(err).Debug("[AUTH] Failed to decode Basic Auth")
					}
				}
			}
		}

		return c.Next()
	}
}
