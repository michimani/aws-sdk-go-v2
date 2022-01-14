// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a Resolver object.
func (c *Client) UpdateResolver(ctx context.Context, params *UpdateResolverInput, optFns ...func(*Options)) (*UpdateResolverOutput, error) {
	if params == nil {
		params = &UpdateResolverInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateResolver", params, optFns, c.addOperationUpdateResolverMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateResolverOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResolverInput struct {

	// The API ID.
	//
	// This member is required.
	ApiId *string

	// The new field name.
	//
	// This member is required.
	FieldName *string

	// The new type name.
	//
	// This member is required.
	TypeName *string

	// The caching configuration for the resolver.
	CachingConfig *types.CachingConfig

	// The new data source name.
	DataSourceName *string

	// The resolver type.
	//
	// * UNIT: A UNIT resolver type. A UNIT resolver is the default
	// resolver type. You can use a UNIT resolver to run a GraphQL query against a
	// single data source.
	//
	// * PIPELINE: A PIPELINE resolver type. You can use a
	// PIPELINE resolver to invoke a series of Function objects in a serial manner. You
	// can use a pipeline resolver to run a GraphQL query against multiple data
	// sources.
	Kind types.ResolverKind

	// The maximum batching size for a resolver.
	MaxBatchSize int32

	// The PipelineConfig.
	PipelineConfig *types.PipelineConfig

	// The new request mapping template. A resolver uses a request mapping template to
	// convert a GraphQL expression into a format that a data source can understand.
	// Mapping templates are written in Apache Velocity Template Language (VTL). VTL
	// request mapping templates are optional when using an Lambda data source. For all
	// other data sources, VTL request and response mapping templates are required.
	RequestMappingTemplate *string

	// The new response mapping template.
	ResponseMappingTemplate *string

	// The SyncConfig for a resolver attached to a versioned data source.
	SyncConfig *types.SyncConfig

	noSmithyDocumentSerde
}

type UpdateResolverOutput struct {

	// The updated Resolver object.
	Resolver *types.Resolver

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateResolverMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateResolver{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateResolver{}, middleware.After)
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
	if err = addOpUpdateResolverValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResolver(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateResolver(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "UpdateResolver",
	}
}
