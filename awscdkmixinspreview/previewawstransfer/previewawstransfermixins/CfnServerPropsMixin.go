package previewawstransfermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawstransfer/previewawstransfermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Instantiates an auto-scaling virtual server based on the selected file transfer protocol in AWS .
//
// When you make updates to your file transfer protocol-enabled server or when you work with users, use the service-generated `ServerId` property that is assigned to the newly created server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServerPropsMixin := awscdkmixinspreview.Mixins.NewCfnServerPropsMixin(&CfnServerMixinProps{
//   	Certificate: jsii.String("certificate"),
//   	Domain: jsii.String("domain"),
//   	EndpointDetails: &EndpointDetailsProperty{
//   		AddressAllocationIds: []*string{
//   			jsii.String("addressAllocationIds"),
//   		},
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		VpcEndpointId: jsii.String("vpcEndpointId"),
//   		VpcId: jsii.String("vpcId"),
//   	},
//   	EndpointType: jsii.String("endpointType"),
//   	IdentityProviderDetails: &IdentityProviderDetailsProperty{
//   		DirectoryId: jsii.String("directoryId"),
//   		Function: jsii.String("function"),
//   		InvocationRole: jsii.String("invocationRole"),
//   		SftpAuthenticationMethods: jsii.String("sftpAuthenticationMethods"),
//   		Url: jsii.String("url"),
//   	},
//   	IdentityProviderType: jsii.String("identityProviderType"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	LoggingRole: jsii.String("loggingRole"),
//   	PostAuthenticationLoginBanner: jsii.String("postAuthenticationLoginBanner"),
//   	PreAuthenticationLoginBanner: jsii.String("preAuthenticationLoginBanner"),
//   	ProtocolDetails: &ProtocolDetailsProperty{
//   		As2Transports: []*string{
//   			jsii.String("as2Transports"),
//   		},
//   		PassiveIp: jsii.String("passiveIp"),
//   		SetStatOption: jsii.String("setStatOption"),
//   		TlsSessionResumptionMode: jsii.String("tlsSessionResumptionMode"),
//   	},
//   	Protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   	S3StorageOptions: &S3StorageOptionsProperty{
//   		DirectoryListingOptimization: jsii.String("directoryListingOptimization"),
//   	},
//   	SecurityPolicyName: jsii.String("securityPolicyName"),
//   	StructuredLogDestinations: []*string{
//   		jsii.String("structuredLogDestinations"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkflowDetails: &WorkflowDetailsProperty{
//   		OnPartialUpload: []interface{}{
//   			&WorkflowDetailProperty{
//   				ExecutionRole: jsii.String("executionRole"),
//   				WorkflowId: jsii.String("workflowId"),
//   			},
//   		},
//   		OnUpload: []interface{}{
//   			&WorkflowDetailProperty{
//   				ExecutionRole: jsii.String("executionRole"),
//   				WorkflowId: jsii.String("workflowId"),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html
//
type CfnServerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnServerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServerPropsMixin
type jsiiProxy_CfnServerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnServerPropsMixin) Props() *CfnServerMixinProps {
	var returns *CfnServerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Transfer::Server`.
func NewCfnServerPropsMixin(props *CfnServerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnServerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Transfer::Server`.
func NewCfnServerPropsMixin_Override(c CfnServerPropsMixin, props *CfnServerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnServerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnServerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnServerPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

