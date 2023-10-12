package main

import (
	"fmt"
	"net/http"
)

const goodReadyzChecks = `[+]ping ok
[+]log ok
[+]etcd ok
[+]informer-sync ok
[+]poststarthook/start-kube-apiserver-admission-initializer ok
[+]poststarthook/generic-apiserver-start-informers ok
[+]poststarthook/start-apiextensions-informers ok
[+]poststarthook/start-apiextensions-controllers ok
[+]poststarthook/crd-informer-synced ok
[+]poststarthook/bootstrap-controller ok
[+]poststarthook/start-system-namespaces-controller ok
[+]poststarthook/start-service-ip-repair-controllers ok
[+]poststarthook/scheduling/bootstrap-system-priority-classes ok
[+]poststarthook/start-cluster-authentication-info-controller ok
[+]poststarthook/start-kube-aggregator-informers ok
[+]poststarthook/apiservice-registration-controller ok
[+]poststarthook/apiservice-status-available-controller ok
[+]poststarthook/kube-apiserver-autoregistration ok
[+]autoregister-completion ok
[+]poststarthook/apiservice-openapi-controller ok
readyz check passed
`

const badReadyzChecks = `[+]ping ok
[+]log ok
[-]etcd failed: reason withheld
[+]informer-sync ok
[+]poststarthook/start-kube-apiserver-admission-initializer ok
[+]poststarthook/generic-apiserver-start-informers ok
[+]poststarthook/start-apiextensions-informers ok
[+]poststarthook/start-apiextensions-controllers ok
[+]poststarthook/crd-informer-synced ok
[+]poststarthook/bootstrap-controller ok
[+]poststarthook/start-system-namespaces-controller ok
[+]poststarthook/start-service-ip-repair-controllers ok
[+]poststarthook/scheduling/bootstrap-system-priority-classes ok
[+]poststarthook/start-cluster-authentication-info-controller ok
[+]poststarthook/start-kube-aggregator-informers ok
[+]poststarthook/apiservice-registration-controller ok
[+]poststarthook/apiservice-status-available-controller ok
[+]poststarthook/kube-apiserver-autoregistration ok
[+]autoregister-completion ok
[+]poststarthook/apiservice-openapi-controller ok
readyz check failed
`

func (c ConfigData) readyz(w http.ResponseWriter, r *http.Request) {
	// Get the verbose query and either return sparse, or enrich the response
	if _, verbose := r.URL.Query()["verbose"]; !verbose {
		if c.ReadyzFail {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "readyz check failed")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
		return
	}

	// enrich and return
	if c.ReadyzFail {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, badReadyzChecks)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, goodReadyzChecks)
}
