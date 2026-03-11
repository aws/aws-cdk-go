package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Anycast static IP list.
//
// For more information, see [Request Anycast static IPs to use for allowlisting](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/request-static-ips.html) in the *Amazon CloudFront Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAnycastIpListPropsMixin := awscdkcfnpropertymixins.Aws_cloudfront.NewCfnAnycastIpListPropsMixin(&CfnAnycastIpListMixinProps{
//   	IpAddressType: jsii.String("ipAddressType"),
//   	IpamCidrConfigs: []interface{}{
//   		&IpamCidrConfigProperty{
//   			Cidr: jsii.String("cidr"),
//   			IpamPoolArn: jsii.String("ipamPoolArn"),
//   		},
//   	},
//   	IpCount: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Tags: &TagsProperty{
//   		Items: []interface{}{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-anycastiplist.html
//
type CfnAnycastIpListPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAnycastIpListMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAnycastIpListPropsMixin
type jsiiProxy_CfnAnycastIpListPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAnycastIpListPropsMixin) Props() *CfnAnycastIpListMixinProps {
	var returns *CfnAnycastIpListMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAnycastIpListPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFront::AnycastIpList`.
func NewCfnAnycastIpListPropsMixin(props *CfnAnycastIpListMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAnycastIpListPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAnycastIpListPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAnycastIpListPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnAnycastIpListPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFront::AnycastIpList`.
func NewCfnAnycastIpListPropsMixin_Override(c CfnAnycastIpListPropsMixin, props *CfnAnycastIpListMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnAnycastIpListPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAnycastIpListPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAnycastIpListPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnAnycastIpListPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAnycastIpListPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnAnycastIpListPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAnycastIpListPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAnycastIpListPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

