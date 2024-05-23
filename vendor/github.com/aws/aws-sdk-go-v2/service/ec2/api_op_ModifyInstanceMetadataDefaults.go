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

// Modifies the default instance metadata service (IMDS) settings at the account
// level in the specified Amazon Web Services  Region.
//
// To remove a parameter's account-level default setting, specify no-preference .
// If an account-level setting is cleared with no-preference , then the instance
// launch considers the other instance metadata settings. For more information, see
// [Order of precedence for instance metadata options]in the Amazon EC2 User Guide.
//
// [Order of precedence for instance metadata options]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-options.html#instance-metadata-options-order-of-precedence
func (c *Client) ModifyInstanceMetadataDefaults(ctx context.Context, params *ModifyInstanceMetadataDefaultsInput, optFns ...func(*Options)) (*ModifyInstanceMetadataDefaultsOutput, error) {
	if params == nil {
		params = &ModifyInstanceMetadataDefaultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyInstanceMetadataDefaults", params, optFns, c.addOperationModifyInstanceMetadataDefaultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyInstanceMetadataDefaultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyInstanceMetadataDefaultsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Enables or disables the IMDS endpoint on an instance. When disabled, the
	// instance metadata can't be accessed.
	HttpEndpoint types.DefaultInstanceMetadataEndpointState

	// The maximum number of hops that the metadata token can travel. To indicate no
	// preference, specify -1 .
	//
	// Possible values: Integers from 1 to 64 , and -1 to indicate no preference
	HttpPutResponseHopLimit *int32

	// Indicates whether IMDSv2 is required.
	//
	//   - optional – IMDSv2 is optional, which means that you can use either IMDSv2 or
	//   IMDSv1.
	//
	//   - required – IMDSv2 is required, which means that IMDSv1 is disabled, and you
	//   must use IMDSv2.
	HttpTokens types.MetadataDefaultHttpTokensState

	// Enables or disables access to an instance's tags from the instance metadata.
	// For more information, see [Work with instance tags using the instance metadata]in the Amazon EC2 User Guide.
	//
	// [Work with instance tags using the instance metadata]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#work-with-tags-in-IMDS
	InstanceMetadataTags types.DefaultInstanceMetadataTagsState

	noSmithyDocumentSerde
}

type ModifyInstanceMetadataDefaultsOutput struct {

	// If the request succeeds, the response returns true . If the request fails, no
	// response is returned, and instead an error message is returned.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyInstanceMetadataDefaultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyInstanceMetadataDefaults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyInstanceMetadataDefaults{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyInstanceMetadataDefaults"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyInstanceMetadataDefaults(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyInstanceMetadataDefaults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyInstanceMetadataDefaults",
	}
}
