// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates an embed URL that you can use to embed an Amazon QuickSight experience
// in your website. This action can be used for any type of user registered in an
// Amazon QuickSight account. Before you use this action, make sure that you have
// configured the relevant Amazon QuickSight resource and permissions. The
// following rules apply to the generated URL:
//
// * It contains a temporary bearer
// token. It is valid for 5 minutes after it is generated. Once redeemed within
// this period, it cannot be re-used again.
//
// * The URL validity period should not
// be confused with the actual session lifetime that can be customized using the
// SessionLifetimeInMinutes
// (https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GenerateEmbedUrlForRegisteredUser.html#QS-GenerateEmbedUrlForRegisteredUser-request-SessionLifetimeInMinutes)
// parameter. The resulting user session is valid for 15 minutes (minimum) to 10
// hours (maximum). The default session duration is 10 hours.
//
// * You are charged
// only when the URL is used or there is interaction with Amazon QuickSight.
//
// For
// more information, see Embedded Analytics
// (https://docs.aws.amazon.com/quicksight/latest/user/embedded-analytics.html) in
// the Amazon QuickSight User Guide. For more information about the high-level
// steps for embedding and for an interactive demo of the ways you can customize
// embedding, visit the Amazon QuickSight Developer Portal
// (https://docs.aws.amazon.com/quicksight/latest/user/quicksight-dev-portal.html).
func (c *Client) GenerateEmbedUrlForRegisteredUser(ctx context.Context, params *GenerateEmbedUrlForRegisteredUserInput, optFns ...func(*Options)) (*GenerateEmbedUrlForRegisteredUserOutput, error) {
	if params == nil {
		params = &GenerateEmbedUrlForRegisteredUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateEmbedUrlForRegisteredUser", params, optFns, c.addOperationGenerateEmbedUrlForRegisteredUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateEmbedUrlForRegisteredUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateEmbedUrlForRegisteredUserInput struct {

	// The ID for the Amazon Web Services account that contains the dashboard that
	// you're embedding.
	//
	// This member is required.
	AwsAccountId *string

	// The experience you are embedding. For registered users, you can embed Amazon
	// QuickSight dashboards or the entire Amazon QuickSight console.
	//
	// This member is required.
	ExperienceConfiguration *types.RegisteredUserEmbeddingExperienceConfiguration

	// The Amazon Resource Name for the registered user.
	//
	// This member is required.
	UserArn *string

	// How many minutes the session is valid. The session lifetime must be in [15-600]
	// minutes range.
	SessionLifetimeInMinutes *int64

	noSmithyDocumentSerde
}

type GenerateEmbedUrlForRegisteredUserOutput struct {

	// The embed URL for the Amazon QuickSight dashboard or console.
	//
	// This member is required.
	EmbedUrl *string

	// The Amazon Web Services request ID for this operation.
	//
	// This member is required.
	RequestId *string

	// The HTTP status of the request.
	//
	// This member is required.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateEmbedUrlForRegisteredUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGenerateEmbedUrlForRegisteredUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGenerateEmbedUrlForRegisteredUser{}, middleware.After)
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
	if err = addOpGenerateEmbedUrlForRegisteredUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateEmbedUrlForRegisteredUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGenerateEmbedUrlForRegisteredUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "GenerateEmbedUrlForRegisteredUser",
	}
}
