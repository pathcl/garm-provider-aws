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

// Describes the specified EC2 Fleet or all of your EC2 Fleets.
//
// If a fleet is of type instant , you must specify the fleet ID in the request,
// otherwise the fleet does not appear in the response.
//
// For more information, see [Describe your EC2 Fleet] in the Amazon EC2 User Guide.
//
// [Describe your EC2 Fleet]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/manage-ec2-fleet.html#monitor-ec2-fleet
func (c *Client) DescribeFleets(ctx context.Context, params *DescribeFleetsInput, optFns ...func(*Options)) (*DescribeFleetsOutput, error) {
	if params == nil {
		params = &DescribeFleetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleets", params, optFns, c.addOperationDescribeFleetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The filters.
	//
	//   - activity-status - The progress of the EC2 Fleet ( error |
	//   pending-fulfillment | pending-termination | fulfilled ).
	//
	//   - excess-capacity-termination-policy - Indicates whether to terminate running
	//   instances if the target capacity is decreased below the current EC2 Fleet size (
	//   true | false ).
	//
	//   - fleet-state - The state of the EC2 Fleet ( submitted | active | deleted |
	//   failed | deleted-running | deleted-terminating | modifying ).
	//
	//   - replace-unhealthy-instances - Indicates whether EC2 Fleet should replace
	//   unhealthy instances ( true | false ).
	//
	//   - type - The type of request ( instant | request | maintain ).
	Filters []types.Filter

	// The IDs of the EC2 Fleets.
	//
	// If a fleet is of type instant , you must specify the fleet ID, otherwise it does
	// not appear in the response.
	FleetIds []string

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see [Pagination].
	//
	// [Pagination]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeFleetsOutput struct {

	// Information about the EC2 Fleets.
	Fleets []types.FleetData

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeFleets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeFleets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFleets"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleets(options.Region), middleware.Before); err != nil {
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

// DescribeFleetsAPIClient is a client that implements the DescribeFleets
// operation.
type DescribeFleetsAPIClient interface {
	DescribeFleets(context.Context, *DescribeFleetsInput, ...func(*Options)) (*DescribeFleetsOutput, error)
}

var _ DescribeFleetsAPIClient = (*Client)(nil)

// DescribeFleetsPaginatorOptions is the paginator options for DescribeFleets
type DescribeFleetsPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see [Pagination].
	//
	// [Pagination]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFleetsPaginator is a paginator for DescribeFleets
type DescribeFleetsPaginator struct {
	options   DescribeFleetsPaginatorOptions
	client    DescribeFleetsAPIClient
	params    *DescribeFleetsInput
	nextToken *string
	firstPage bool
}

// NewDescribeFleetsPaginator returns a new DescribeFleetsPaginator
func NewDescribeFleetsPaginator(client DescribeFleetsAPIClient, params *DescribeFleetsInput, optFns ...func(*DescribeFleetsPaginatorOptions)) *DescribeFleetsPaginator {
	if params == nil {
		params = &DescribeFleetsInput{}
	}

	options := DescribeFleetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFleetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFleetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFleets page.
func (p *DescribeFleetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFleetsOutput, error) {
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

	result, err := p.client.DescribeFleets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeFleets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFleets",
	}
}
