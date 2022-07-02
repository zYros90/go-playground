<!--- ################### Profiling ################### --->

# Signals
```
sig := make(chan os.Signal, 1)
signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
<-sig
```

# Profiling

#### 1. github.com/pkg/profile

* import `"github.com/pkg/profile"`
* `defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()`
* after ctrg+C a\*.pprof file will be generated in ./
* `pprof -http :9999 mem.pprof`
* see to localhost:9999
* `defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()`
* \`go tool trace trace.out
* `defer profile.Start(profile.GoroutineProfile, profile.ProfilePath(".")).Stop()`
* 

#### 2. pprof

* import \_ "net/http/pprof"

```
 go func() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

* wget -O trace.out http://localhost:6060/debug/pprof/trace?seconds=10
* go tool trace trace.out
* go tool pprof http://localhost:6060/debug/pprof/heap?seconds=30
* go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
* go tool pprof http://localhost:6060/debug/pprof/block

#### 3. statsview

https://github.com/go-echarts/statsview

`go get -u github.com/go-echarts/statsview/...`

```
go func () {
    mgr := statsview.New()

    // Start() runs a HTTP server at `localhost:18066` by default.
    go mgr.Start()
}()

```

<!--- ################### Pre-Commit ################### --->

# Pre-Commit

https://pre-commit.com/

`pip install pre-commit`

create .pre-commit-config.yaml

```
repos:
- repo: git://github.com/dnephin/pre-commit-golang
  rev: master
  hooks:
  - id: go-fmt
  # - id: go-vet
  # - id: go-imports
  # - id: go-critic
  # - id: go-unit-tests
  - id: golangci-lint
  - id: go-mod-tidy
```

run `pre-commit install`

<!--- ################### Ghorg ################### --->

# Ghorg

https://github.com/gabrie30/ghorg setup for access token see link

#### Gitlab

* replace `yourProject` and `gitlab.xx.de` on self hosted gitlab instance

```
ghorg clone yourProject --protocol=ssh --branch=develop --scm=gitlab --base-url=https://your.hosted.gitlab.com --preserve-dir
```

on gitlab cloud

```
ghorg clone yourProject --protocol=ssh --branch=develop --scm=gitlab --base-url=https://gitlab.xx.de --preserve-dir
```

#### Github

* replace `yourOrganization`

```
ghorg clone yourOrganization --protocol=ssh --branch=develop --base-url=https://internal.github.com
```