This simple project provides a HTTP mock service for K8s and K3s `/livez` and `/readyz` end-points. You can configure on startup to fail or pass either.

```bash
# Example of a passing /readyz check
➜ ./health -config config.toml
INFO[2023-10-12T18:18:55+01:00] Checking config                              
INFO[2023-10-12T18:18:55+01:00] Config OK                                    
INFO[2023-10-12T18:18:55+01:00] /livez fail: false                           
INFO[2023-10-12T18:18:55+01:00] /readyz fail: false                          
INFO[2023-10-12T18:18:55+01:00] Listening on port: 8080                      

# Open up /localhost:8080/readyz?verbose on a browser
[+]ping ok
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

# Example of a passing /readyz check
➜ ./health -config config.toml -readyz=true
INFO[2023-10-12T18:20:06+01:00] Checking config                              
INFO[2023-10-12T18:20:06+01:00] Config OK                                    
INFO[2023-10-12T18:20:06+01:00] /livez fail: false                           
INFO[2023-10-12T18:20:06+01:00] /readyz fail: true                           
INFO[2023-10-12T18:20:06+01:00] Listening on port: 8080                      

# Open up /localhost:8080/readyz?verbose on a browser
[+]ping ok
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
```

### Why?

Because I needed to regularly test logic from K8s/K3s.

### Security

None. This is plain text and ignored Bearer token inputs etc.