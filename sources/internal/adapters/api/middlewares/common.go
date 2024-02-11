// HRACH_DEV © iMed Cloud Services, Inc.
package middlewares

import (
	"shtem-api/sources/internal/adapters/api/dto"
	"shtem-api/sources/internal/configs"
	"shtem-api/sources/internal/core/domain"
	"time"

	connLimit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

const maxConcurrentCon int = 10

var (
	allowedHeaders = []string{
		"Authorization", "X-Forwarded-For", "X-Shtem-Api-Key", "Content-Type",
	}
	allowedMethods = []string{
		"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS",
	}
	nonCompressables = []string{
		".png", ".gif", ".jpeg", ".jpg", ".pdf", ".mp4", ".avi", ".mov", ".webp", ".mp3",
	}
)

func apiKeyValidator(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-Shtem-Api-Key") != key {
			dto.WriteErrorResponse(c, domain.ErrAccessDenied)
			c.Abort()
		}
	}
}

func ApplyCommonMiddlewares(r *gin.Engine, cfg *configs.Configs) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: allowedMethods,
		AllowHeaders: allowedHeaders,
		MaxAge:       1 * time.Hour,
	}))
	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions(nonCompressables)))
	r.Use(connLimit.MaxAllowed(maxConcurrentCon))
	// r.Use(apiKeyValidator(cfg.Global.PublicAPIKey))
}
