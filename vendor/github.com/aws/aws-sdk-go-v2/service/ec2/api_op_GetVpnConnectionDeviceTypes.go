// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Obtain a list of customer gateway devices for which sample configuration files
// can be provided. The request has no additional parameters. You can also see the
// list of device types with sample configuration files available under [Your customer gateway device]in the
// Amazon Web Services Site-to-Site VPN User Guide.
//
// [Your customer gateway device]: https://docs.aws.amazon.com/vpn/latest/s2svpn/your-cgw.html
func (c *Client) GetVpnConnectionDeviceTypes(ctx context.Context, params *GetVpnConnectionDeviceTypesInput, optFns ...func(*Options)) (*GetVpnConnectionDeviceTypesOutput, error) {
	if params == nil {
		params = &GetVpnConnectionDeviceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVpnConnectionDeviceTypes", params, optFns, c.addOperationGetVpnConnectionDeviceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVpnConnectionDeviceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVpnConnectionDeviceTypesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The maximum number of results returned by GetVpnConnectionDeviceTypes in
	// paginated output. When this parameter is used, GetVpnConnectionDeviceTypes only
	// returns MaxResults results in a single page along with a NextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another GetVpnConnectionDeviceTypes request with the returned NextToken value.
	// This value can be between 200 and 1000. If this parameter is not used, then
	// GetVpnConnectionDeviceTypes returns all results.
	MaxResults *int32

	// The NextToken value returned from a previous paginated
	// GetVpnConnectionDeviceTypes request where MaxResults was used and the results
	// exceeded the value of that parameter. Pagination continues from the end of the
	// previous results that returned the NextToken value. This value is null when
	// there are no more results to return.
	NextToken *string

	noSmithyDocumentSerde
}

type GetVpnConnectionDeviceTypesOutput struct {

	// The NextToken value to include in a future GetVpnConnectionDeviceTypes request.
	// When the results of a GetVpnConnectionDeviceTypes request exceed MaxResults ,
	// this value can be used to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string

	// List of customer gateway devices that have a sample configuration file
	// available for use.
	VpnConnectionDeviceTypes []types.VpnConnectionDeviceType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetVpnConnectionDeviceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetVpnConnectionDeviceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetVpnConnectionDeviceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetVpnConnectionDeviceTypes"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVpnConnectionDeviceTypes(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// GetVpnConnectionDeviceTypesAPIClient is a client that implements the
// GetVpnConnectionDeviceTypes operation.
type GetVpnConnectionDeviceTypesAPIClient interface {
	GetVpnConnectionDeviceTypes(context.Context, *GetVpnConnectionDeviceTypesInput, ...func(*Options)) (*GetVpnConnectionDeviceTypesOutput, error)
}

var _ GetVpnConnectionDeviceTypesAPIClient = (*Client)(nil)

// GetVpnConnectionDeviceTypesPaginatorOptions is the paginator options for
// GetVpnConnectionDeviceTypes
type GetVpnConnectionDeviceTypesPaginatorOptions struct {
	// The maximum number of results returned by GetVpnConnectionDeviceTypes in
	// paginated output. When this parameter is used, GetVpnConnectionDeviceTypes only
	// returns MaxResults results in a single page along with a NextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another GetVpnConnectionDeviceTypes request with the returned NextToken value.
	// This value can be between 200 and 1000. If this parameter is not used, then
	// GetVpnConnectionDeviceTypes returns all results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetVpnConnectionDeviceTypesPaginator is a paginator for
// GetVpnConnectionDeviceTypes
type GetVpnConnectionDeviceTypesPaginator struct {
	options   GetVpnConnectionDeviceTypesPaginatorOptions
	client    GetVpnConnectionDeviceTypesAPIClient
	params    *GetVpnConnectionDeviceTypesInput
	nextToken *string
	firstPage bool
}

// NewGetVpnConnectionDeviceTypesPaginator returns a new
// GetVpnConnectionDeviceTypesPaginator
func NewGetVpnConnectionDeviceTypesPaginator(client GetVpnConnectionDeviceTypesAPIClient, params *GetVpnConnectionDeviceTypesInput, optFns ...func(*GetVpnConnectionDeviceTypesPaginatorOptions)) *GetVpnConnectionDeviceTypesPaginator {
	if params == nil {
		params = &GetVpnConnectionDeviceTypesInput{}
	}

	options := GetVpnConnectionDeviceTypesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetVpnConnectionDeviceTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetVpnConnectionDeviceTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetVpnConnectionDeviceTypes page.
func (p *GetVpnConnectionDeviceTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetVpnConnectionDeviceTypesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetVpnConnectionDeviceTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetVpnConnectionDeviceTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetVpnConnectionDeviceTypes",
	}
}
