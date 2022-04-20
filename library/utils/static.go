package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 静态资源编译
type embeddingFileSystem struct {
	fs http.FileSystem
}

func (b *embeddingFileSystem) Exists(c *gin.Context, prefix string, filepath string) bool {
	if prefix == "" {
		if _, err := b.fs.Open(filepath); err != nil {
			return false
		}
		return true
	} else {
		if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
			if _, err := b.fs.Open(p); err != nil {
				return false
			}
			return true
		}
	}
	return false
}

func EmbeddingFileSystem(fs http.FileSystem) *embeddingFileSystem {
	return &embeddingFileSystem{
		fs: fs,
	}
}

func Serve(urlPrefix string, fs *embeddingFileSystem) gin.HandlerFunc {
	fileServer := http.FileServer(fs.fs)
	if urlPrefix != "" {
		fileServer = http.StripPrefix(urlPrefix, fileServer)
	}
	return func(c *gin.Context) {
		if fs.Exists(c, urlPrefix, c.Request.URL.Path) {
			fileServer.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
