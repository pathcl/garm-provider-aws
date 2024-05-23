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

// A trust provider is a third-party entity that creates, maintains, and manages
// identity information for users and devices. When an application request is made,
// the identity information sent by the trust provider is evaluated by Verified
// Access before allowing or denying the application request.
func (c *Client) CreateVerifiedAccessTrustProvider(ctx context.Context, params *CreateVerifiedAccessTrustProviderInput, optFns ...func(*Options)) (*CreateVerifiedAccessTrustProviderOutput, error) {
	if params == nil {
		params = &CreateVerifiedAccessTrustProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVerifiedAccessTrustProvider", params, optFns, c.addOperationCreateVerifiedAccessTrustProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVerifiedAccessTrustProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVerifiedAccessTrustProviderInput struct {

	// The identifier to be used when working with policy rules.
	//
	// This member is required.
	PolicyReferenceName *string

	// The type of trust provider.
	//
	// This member is required.
	TrustProviderType types.TrustProviderType

	// A unique, case-sensitive token that you provide to ensure idempotency of your
	// modification request. For more information, see [Ensuring Idempotency].
	//
	// [Ensuring Idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	ClientToken *string

	// A description for the Verified Access trust provider.
	Description *string

	// The options for a device-based trust provider. This parameter is required when
	// the provider type is device .
	DeviceOptions *types.CreateVerifiedAccessTrustProviderDeviceOptions

	// The type of device-based trust provider. This parameter is required when the
	// provider type is device .
	DeviceTrustProviderType types.DeviceTrustProviderType

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The options for a OpenID Connect-compatible user-identity trust provider. This
	// parameter is required when the provider type is user .
	OidcOptions *types.CreateVerifiedAccessTrustProviderOidcOptions

	// The options for server side encryption.
	SseSpecification *types.VerifiedAccessSseSpecificationRequest

	// The tags to assign to the Verified Access trust provider.
	TagSpecifications []types.TagSpecification

	// The type of user-based trust provider. This parameter is required when the
	// provider type is user .
	UserTrustProviderType types.UserTrustProviderType

	noSmithyDocumentSerde
}

type CreateVerifiedAccessTrustProviderOutput struct {

	// Details about the Verified Access trust provider.
	VerifiedAccessTrustProvider *types.VerifiedAccessTrustProvider

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVerifiedAccessTrustProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateVerifiedAccessTrustProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateVerifiedAccessTrustProvider{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateVerifiedAccessTrustProvider"); err != nil {
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
	if err = addIdempotencyToken_opCreateVerifiedAccessTrustProviderMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateVerifiedAccessTrustProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVerifiedAccessTrustProvider(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateVerifiedAccessTrustProvider struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateVerifiedAccessTrustProvider) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateVerifiedAccessTrustProvider) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateVerifiedAccessTrustProviderInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateVerifiedAccessTrustProviderInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateVerifiedAccessTrustProviderMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateVerifiedAccessTrustProvider{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateVerifiedAccessTrustProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateVerifiedAccessTrustProvider",
	}
}
