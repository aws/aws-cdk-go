package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A verification token is an AWS-generated random value that you can use to prove ownership of an external resource.
//
// For example, you can use a verification token to validate that you control a public IP address range when you bring an IP address range to AWS (BYOIP).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnIpamExternalResourceVerificationTokenPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnIpamExternalResourceVerificationTokenPropsMixin(&CfnIpamExternalResourceVerificationTokenMixinProps{
//   	IpamId: jsii.String("ipamId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamexternalresourceverificationtoken.html
//
type CfnIpamExternalResourceVerificationTokenPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnIpamExternalResourceVerificationTokenMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIpamExternalResourceVerificationTokenPropsMixin
type jsiiProxy_CfnIpamExternalResourceVerificationTokenPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnIpamExternalResourceVerificationTokenPropsMixin) Props() *CfnIpamExternalResourceVerificationTokenMixinProps {
	var returns *CfnIpamExternalResourceVerificationTokenMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIpamExternalResourceVerificationTokenPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::IpamExternalResourceVerificationToken`.
func NewCfnIpamExternalResourceVerificationTokenPropsMixin(props *CfnIpamExternalResourceVerificationTokenMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnIpamExternalResourceVerificationTokenPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIpamExternalResourceVerificationTokenPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIpamExternalResourceVerificationTokenPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnIpamExternalResourceVerificationTokenPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::IpamExternalResourceVerificationToken`.
func NewCfnIpamExternalResourceVerificationTokenPropsMixin_Override(c CfnIpamExternalResourceVerificationTokenPropsMixin, props *CfnIpamExternalResourceVerificationTokenMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnIpamExternalResourceVerificationTokenPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnIpamExternalResourceVerificationTokenPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIpamExternalResourceVerificationTokenPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnIpamExternalResourceVerificationTokenPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIpamExternalResourceVerificationTokenPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnIpamExternalResourceVerificationTokenPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIpamExternalResourceVerificationTokenPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnIpamExternalResourceVerificationTokenPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

