package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (c *ConfigData) mode(w http.ResponseWriter, r *http.Request) {
	// Get the verbose query and either return sparse, or enrich the response
	if val, readyz := r.URL.Query()["readyz"]; readyz {

		// check the first value of val, if it exists set it to true
		if len(val) > 0 {
			if val[0] == "true" {
				c.SetReadyzFail(false)
				log.Info("Setting readyz to pass")
			} else {
				c.SetReadyzFail(true)
				log.Info("Setting readyz to fail")
			}
		}
	}

	// Get the verbose query and either return sparse, or enrich the response
	if val, livez := r.URL.Query()["livez"]; livez {

		// check the first value of val, if it exists set it to true
		if len(val) > 0 {
			if val[0] == "true" {
				c.SetLivezFail(false)
				log.Info("Setting livez to pass")
			} else {
				c.SetLivezFail(true)
				log.Info("Setting livez to fail")
			}
		}

	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "readyz will pass: "+fmt.Sprint(!c.ReadyzFail)+", livez will pass: "+fmt.Sprint(!c.LivezFail))
}
