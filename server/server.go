package server

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tcdw/config-server/config"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

type IncData struct {
	Variant string
}

func Start(conf config.Config) {
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.GET("/:token/:data/:variant", func(c *gin.Context) {
		token := c.Param("token")
		data := c.Param("data")
		variant := c.Param("variant")
		if token != conf.Token {
			c.String(http.StatusNotFound, "")
			return
		}

		appendPath := fmt.Sprintf("%s/%s.txt", conf.TemplatePath, data)
		tmplPath, err := filepath.Abs(appendPath)
		tmplData := IncData{
			Variant: variant,
		}
		tmpl, err := template.New(fmt.Sprintf("%s.txt", data)).ParseFiles(tmplPath)
		if err != nil {
			log.Print(err)
			c.String(http.StatusNotFound, "")
			return
		}

		var out bytes.Buffer
		err = tmpl.Execute(&out, tmplData)
		if err != nil {
			log.Print(err)
			c.String(http.StatusNotFound, "")
			return
		}

		c.String(http.StatusOK, "%s", out.String())
	})
	r.Run(fmt.Sprintf("127.0.0.1:%d", int(conf.Port))) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
