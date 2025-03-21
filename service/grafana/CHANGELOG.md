# v1.9.2 (2022-06-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.1 (2022-05-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.0 (2022-05-13)

* **Feature**: This release adds APIs for creating and deleting API keys in an Amazon Managed Grafana workspace.

# v1.8.1 (2022-04-25)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.0 (2022-03-31)

* **Feature**: This release adds tagging support to the Managed Grafana service. New APIs: TagResource, UntagResource and ListTagsForResource. Updates: add optional field tags to support tagging while calling CreateWorkspace.

# v1.7.3 (2022-03-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.2 (2022-03-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.1 (2022-03-23)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2022-03-08)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.0 (2022-02-24)

* **Feature**: API client updated
* **Feature**: Adds RetryMaxAttempts and RetryMod to API client Options. This allows the API clients' default Retryer to be configured from the shared configuration files or environment variables. Adding a new Retry mode of `Adaptive`. `Adaptive` retry mode is an experimental mode, adding client rate limiting when throttles reponses are received from an API. See [retry.AdaptiveMode](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/aws/retry#AdaptiveMode) for more details, and configuration options.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2022-01-14)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.0 (2022-01-07)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.0 (2021-12-21)

* **Feature**: API Paginators now support specifying the initial starting token, and support stopping on empty string tokens.

# v1.2.3 (2021-12-03)

* **Bug Fix**: Fixed an issue that prevent auto-filling of an API's idempotency parameters when not explictly provided by the caller.

# v1.2.2 (2021-12-02)

* **Bug Fix**: Fixes a bug that prevented aws.EndpointResolverWithOptions from being used by the service client. ([#1514](https://github.com/aws/aws-sdk-go-v2/pull/1514))
* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.1 (2021-11-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.0 (2021-11-06)

* **Feature**: The SDK now supports configuration of FIPS and DualStack endpoints using environment variables, shared configuration, or programmatically.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.0 (2021-10-21)

* **Feature**: Updated  to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.0 (2021-10-11)

* **Release**: New AWS service client module
* **Feature**: API client updated
* **Dependency Update**: Updated to the latest SDK module versions

