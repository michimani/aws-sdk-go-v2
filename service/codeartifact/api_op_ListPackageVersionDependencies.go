// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the direct dependencies for a package version. The dependencies are
// returned as PackageDependency
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageDependency.html)
// objects. CodeArtifact extracts the dependencies for a package version from the
// metadata file for the package format (for example, the package.json file for npm
// packages and the pom.xml file for Maven). Any package version dependencies that
// are not listed in the configuration file are not returned.
func (c *Client) ListPackageVersionDependencies(ctx context.Context, params *ListPackageVersionDependenciesInput, optFns ...func(*Options)) (*ListPackageVersionDependenciesOutput, error) {
	if params == nil {
		params = &ListPackageVersionDependenciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPackageVersionDependencies", params, optFns, c.addOperationListPackageVersionDependenciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPackageVersionDependenciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackageVersionDependenciesInput struct {

	// The name of the domain that contains the repository that contains the requested
	// package version dependencies.
	//
	// This member is required.
	Domain *string

	// The format of the package with the requested dependencies.
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the package versions' package.
	//
	// This member is required.
	Package *string

	// A string that contains the package version (for example, 3.5.2).
	//
	// This member is required.
	PackageVersion *string

	// The name of the repository that contains the requested package version.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	// * The namespace of a Maven package is its
	// groupId.
	//
	// * The namespace of an npm package is its scope.
	//
	// * A Python package
	// does not contain a corresponding component, so Python packages do not have a
	// namespace.
	Namespace *string

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPackageVersionDependenciesOutput struct {

	// The returned list of PackageDependency
	// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageDependency.html)
	// objects.
	Dependencies []types.PackageDependency

	// A format that specifies the type of the package that contains the returned
	// dependencies.
	Format types.PackageFormat

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	// * The namespace of a Maven package is its
	// groupId.
	//
	// * The namespace of an npm package is its scope.
	//
	// * A Python package
	// does not contain a corresponding component, so Python packages do not have a
	// namespace.
	Namespace *string

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The name of the package that contains the returned package versions
	// dependencies.
	Package *string

	// The version of the package that is specified in the request.
	Version *string

	// The current revision associated with the package version.
	VersionRevision *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPackageVersionDependenciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPackageVersionDependencies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackageVersionDependencies{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListPackageVersionDependenciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPackageVersionDependencies(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListPackageVersionDependencies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "ListPackageVersionDependencies",
	}
}
