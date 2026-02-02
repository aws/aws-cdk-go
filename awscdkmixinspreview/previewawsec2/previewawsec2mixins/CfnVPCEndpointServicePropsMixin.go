package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a VPC endpoint service configuration to which service consumers ( AWS accounts, users, and IAM roles) can connect.
//
// To create an endpoint service configuration, you must first create one of the following for your service:
//
// - A [Network Load Balancer](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/introduction.html) . Service consumers connect to your service using an interface endpoint.
// - A [Gateway Load Balancer](https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/introduction.html) . Service consumers connect to your service using a Gateway Load Balancer endpoint.
//
// For more information, see the [AWS PrivateLink User Guide](https://docs.aws.amazon.com/vpc/latest/privatelink/) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPCEndpointServicePropsMixin := awscdkmixinspreview.Mixins.NewCfnVPCEndpointServicePropsMixin(&CfnVPCEndpointServiceMixinProps{
//   	AcceptanceRequired: jsii.Boolean(false),
//   	ContributorInsightsEnabled: jsii.Boolean(false),
//   	GatewayLoadBalancerArns: []*string{
//   		jsii.String("gatewayLoadBalancerArns"),
//   	},
//   	NetworkLoadBalancerArns: []*string{
//   		jsii.String("networkLoadBalancerArns"),
//   	},
//   	PayerResponsibility: jsii.String("payerResponsibility"),
//   	SupportedIpAddressTypes: []*string{
//   		jsii.String("supportedIpAddressTypes"),
//   	},
//   	SupportedRegions: []*string{
//   		jsii.String("supportedRegions"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservice.html
//
type CfnVPCEndpointServicePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnVPCEndpointServiceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVPCEndpointServicePropsMixin
type jsiiProxy_CfnVPCEndpointServicePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnVPCEndpointServicePropsMixin) Props() *CfnVPCEndpointServiceMixinProps {
	var returns *CfnVPCEndpointServiceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCEndpointServicePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::VPCEndpointService`.
func NewCfnVPCEndpointServicePropsMixin(props *CfnVPCEndpointServiceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnVPCEndpointServicePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVPCEndpointServicePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVPCEndpointServicePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCEndpointServicePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::VPCEndpointService`.
func NewCfnVPCEndpointServicePropsMixin_Override(c CfnVPCEndpointServicePropsMixin, props *CfnVPCEndpointServiceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCEndpointServicePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnVPCEndpointServicePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVPCEndpointServicePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCEndpointServicePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVPCEndpointServicePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCEndpointServicePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVPCEndpointServicePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnVPCEndpointServicePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

