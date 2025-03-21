# v1.10.1 (2022-06-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.10.0 (2022-06-02)

* **Feature**: Added a new attribute ServerSideEncryptionUpdateDetails to Domain and DomainSummary.

# v1.9.0 (2022-05-25)

* **Feature**: VoiceID will now automatically expire Speakers if they haven't been accessed for Enrollment, Re-enrollment or Successful Auth for three years. The Speaker APIs now return a "LastAccessedAt" time for Speakers, and the EvaluateSession API returns "SPEAKER_EXPIRED" Auth Decision for EXPIRED Speakers.

# v1.8.5 (2022-05-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.4 (2022-04-25)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.3 (2022-03-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.2 (2022-03-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.1 (2022-03-23)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.0 (2022-03-08)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2022-02-24)

* **Feature**: API client updated
* **Feature**: Adds RetryMaxAttempts and RetryMod to API client Options. This allows the API clients' default Retryer to be configured from the shared configuration files or environment variables. Adding a new Retry mode of `Adaptive`. `Adaptive` retry mode is an experimental mode, adding client rate limiting when throttles reponses are received from an API. See [retry.AdaptiveMode](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/aws/retry#AdaptiveMode) for more details, and configuration options.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.6.0 (2022-01-14)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2022-01-07)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.0 (2021-12-21)

* **Feature**: API Paginators now support specifying the initial starting token, and support stopping on empty string tokens.

# v1.3.2 (2021-12-02)

* **Bug Fix**: Fixes a bug that prevented aws.EndpointResolverWithOptions from being used by the service client. ([#1514](https://github.com/aws/aws-sdk-go-v2/pull/1514))
* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.1 (2021-11-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.0 (2021-11-12)

* **Feature**: Service clients now support custom endpoints that have an initial URI path defined.

# v1.2.0 (2021-11-06)

* **Feature**: The SDK now supports configuration of FIPS and DualStack endpoints using environment variables, shared configuration, or programmatically.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.0 (2021-10-21)

* **Feature**: Updated  to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.1 (2021-10-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.0 (2021-09-30)

* **Release**: New AWS service client module
* **Feature**: API client updated

