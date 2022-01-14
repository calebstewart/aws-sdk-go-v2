// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the following prices for either all the TLDs supported by Route 53, or the
// specified TLD:
//
// * Registration
//
// * Transfer
//
// * Owner change
//
// * Domain renewal
//
// *
// Domain restoration
func (c *Client) ListPrices(ctx context.Context, params *ListPricesInput, optFns ...func(*Options)) (*ListPricesOutput, error) {
	if params == nil {
		params = &ListPricesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPrices", params, optFns, c.addOperationListPricesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPricesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPricesInput struct {

	// For an initial request for a list of prices, omit this element. If the number of
	// prices that are not yet complete is greater than the value that you specified
	// for MaxItems, you can use Marker to return additional prices. Get the value of
	// NextPageMarker from the previous response, and submit another request that
	// includes the value of NextPageMarker in the Marker element. Used only for all
	// TLDs. If you specify a TLD, don't specify a Marker.
	Marker *string

	// Number of Prices to be returned. Used only for all TLDs. If you specify a TLD,
	// don't specify a MaxItems.
	MaxItems *int32

	// The TLD for which you want to receive the pricing information. For example.
	// .net. If a Tld value is not provided, a list of prices for all TLDs supported by
	// Route 53 is returned.
	Tld *string

	noSmithyDocumentSerde
}

type ListPricesOutput struct {

	// A complex type that includes all the pricing information. If you specify a TLD,
	// this array contains only the pricing for that TLD.
	//
	// This member is required.
	Prices []types.DomainPrice

	// If there are more prices than you specified for MaxItems in the request, submit
	// another request and include the value of NextPageMarker in the value of Marker.
	// Used only for all TLDs. If you specify a TLD, don't specify a NextPageMarker.
	NextPageMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPricesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPrices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPrices{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPrices(options.Region), middleware.Before); err != nil {
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

// ListPricesAPIClient is a client that implements the ListPrices operation.
type ListPricesAPIClient interface {
	ListPrices(context.Context, *ListPricesInput, ...func(*Options)) (*ListPricesOutput, error)
}

var _ ListPricesAPIClient = (*Client)(nil)

// ListPricesPaginatorOptions is the paginator options for ListPrices
type ListPricesPaginatorOptions struct {
	// Number of Prices to be returned. Used only for all TLDs. If you specify a TLD,
	// don't specify a MaxItems.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPricesPaginator is a paginator for ListPrices
type ListPricesPaginator struct {
	options   ListPricesPaginatorOptions
	client    ListPricesAPIClient
	params    *ListPricesInput
	nextToken *string
	firstPage bool
}

// NewListPricesPaginator returns a new ListPricesPaginator
func NewListPricesPaginator(client ListPricesAPIClient, params *ListPricesInput, optFns ...func(*ListPricesPaginatorOptions)) *ListPricesPaginator {
	if params == nil {
		params = &ListPricesInput{}
	}

	options := ListPricesPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPricesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPricesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPrices page.
func (p *ListPricesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPricesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListPrices(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListPrices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "ListPrices",
	}
}