# argo-rollouts-api

This module contains just the Kubernetes API resources and types for https://github.com/argoproj/argo-rollouts (from pkg/apis and pkg/client), minus pkg/apis/rollout/validation. This allows us to import and use these types without having to transitively import k8s.io/kubernetes, which argoproj/argo-rollouts depends on.

## How to update

To update the API definitions (e.g. from a newer release of Argo Rollouts), you can do the following:
1. Checkout https://github.com/argoproj/argo-rollouts locally at the release tag you want to copy from
2. Copy the relevant packages over:
    ```shell
    $ cp -fr $argo_rollouts_folder/pkg/apis ./pkg
    $ cp -fr $argo_rollouts_folder/pkg/client ./pkg
    ```
3. Delete the `validation` package as this has a dependency on `argo-rollouts` and we don't need it just for API object usages:
    ```shell
    $ rm -fr ./pkg/apis/rollouts/validation
    ```
4. Raise a PR with the changes to monzo/argo-rollouts-api, get it merged
5. In all **other projects** where you want to add this as a dependency, add a go mod replace directive:
   ```shell
   go mod edit -replace github.com/argoproj/argo-rollouts=github.com/monzo/argo-rollouts-api@$merge-commit-sha
   ```



## Why is this a problem, and why can't we just use `argoproj/argo-rollouts`?

The `argoproj/argo-rollouts` repository contains these API types, but also all of the code for the Argo Rollouts controllers. Many places in this controller code import code from `k8s.io/kubernetes`, which is the main holding repository for the Kubernetes project (https://github.com/kubernetes/kubernetes).

This repository contains the staging directories for many published Kubernetes packages. For example, you can find the original source for `k8s.io/api` at [kubernetes/kubernetes/staging/src/k8s.io/api](https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/api). To make it possible to test these staging packages with the rest of the codebase, the kubernetes/kubernetes `go.mod` makes heavy use of `replace` directives to redirect imports of various `k8s.io` packages to the respective staging directories. For example for `k8s.io/api`:
```
...

require (
   ...
   k8s.io/api v0.0.0
   ...
)

replace (
   ...
   k8s.io/api => ./staging/src/k8s.io/api
   ...
)
```

Go modules don't propagate `replace` directives to their downstream dependencies - this would cause chaos if they did! Instead, if a project imports `k8s.io/kubernetes`, it just gets the `require` versions - in this case, `k8s.io/api = v0.0.0`. Since this isn't a real version, go mod dependency resolution fails.

The workaround for this is to replicate the `replace` directives, but instead target actual versions, e.g.:
```

replace (
   ...
   k8s.io/api => k8s.io/api v0.24.2
   ...
)
```

This is a fair amount of complexity, which we simply don't need. We just want to use the API types, not the controllers. So we've extracted the API types into this repository, which is a much simpler dependency to manage.


See https://github.com/kubernetes/kubernetes/issues/90358#issuecomment-617859364
