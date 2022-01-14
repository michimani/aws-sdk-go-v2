// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The operation to list queued messages.
func (c *Client) ListQueuedMessages(ctx context.Context, params *ListQueuedMessagesInput, optFns ...func(*Options)) (*ListQueuedMessagesOutput, error) {
	if params == nil {
		params = &ListQueuedMessagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQueuedMessages", params, optFns, c.addOperationListQueuedMessagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQueuedMessagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQueuedMessagesInput struct {

	// Id of a given wireless device which the downlink packets are targeted
	//
	// This member is required.
	Id *string

	// The maximum number of results to return in this operation.
	MaxResults int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	// The wireless device type, it is either Sidewalk or LoRaWAN.
	WirelessDeviceType types.WirelessDeviceType

	noSmithyDocumentSerde
}

type ListQueuedMessagesOutput struct {

	// The messages in downlink queue.
	DownlinkQueueMessagesList []types.DownlinkQueueMessage

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListQueuedMessagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListQueuedMessages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListQueuedMessages{}, middleware.After)
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
	if err = addOpListQueuedMessagesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQueuedMessages(options.Region), middleware.Before); err != nil {
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

// ListQueuedMessagesAPIClient is a client that implements the ListQueuedMessages
// operation.
type ListQueuedMessagesAPIClient interface {
	ListQueuedMessages(context.Context, *ListQueuedMessagesInput, ...func(*Options)) (*ListQueuedMessagesOutput, error)
}

var _ ListQueuedMessagesAPIClient = (*Client)(nil)

// ListQueuedMessagesPaginatorOptions is the paginator options for
// ListQueuedMessages
type ListQueuedMessagesPaginatorOptions struct {
	// The maximum number of results to return in this operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListQueuedMessagesPaginator is a paginator for ListQueuedMessages
type ListQueuedMessagesPaginator struct {
	options   ListQueuedMessagesPaginatorOptions
	client    ListQueuedMessagesAPIClient
	params    *ListQueuedMessagesInput
	nextToken *string
	firstPage bool
}

// NewListQueuedMessagesPaginator returns a new ListQueuedMessagesPaginator
func NewListQueuedMessagesPaginator(client ListQueuedMessagesAPIClient, params *ListQueuedMessagesInput, optFns ...func(*ListQueuedMessagesPaginatorOptions)) *ListQueuedMessagesPaginator {
	if params == nil {
		params = &ListQueuedMessagesInput{}
	}

	options := ListQueuedMessagesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListQueuedMessagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListQueuedMessagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListQueuedMessages page.
func (p *ListQueuedMessagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListQueuedMessagesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListQueuedMessages(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListQueuedMessages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "ListQueuedMessages",
	}
}
