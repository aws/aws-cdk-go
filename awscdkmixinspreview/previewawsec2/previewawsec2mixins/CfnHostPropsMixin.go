package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Allocates a fully dedicated physical server for launching EC2 instances.
//
// Because the host is fully dedicated for your use, it can help you address compliance requirements and reduce costs by allowing you to use your existing server-bound software licenses. For more information, see [Dedicated Hosts](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-overview.html) in the *Amazon EC2 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHostPropsMixin := awscdkmixinspreview.Mixins.NewCfnHostPropsMixin(&CfnHostMixinProps{
//   	AssetId: jsii.String("assetId"),
//   	AutoPlacement: jsii.String("autoPlacement"),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	HostMaintenance: jsii.String("hostMaintenance"),
//   	HostRecovery: jsii.String("hostRecovery"),
//   	InstanceFamily: jsii.String("instanceFamily"),
//   	InstanceType: jsii.String("instanceType"),
//   	OutpostArn: jsii.String("outpostArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html
//
type CfnHostPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnHostMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnHostPropsMixin
type jsiiProxy_CfnHostPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnHostPropsMixin) Props() *CfnHostMixinProps {
	var returns *CfnHostMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHostPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::Host`.
func NewCfnHostPropsMixin(props *CfnHostMixinProps, options *mixins.CfnPropertyMixinOptions) CfnHostPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnHostPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnHostPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnHostPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::Host`.
func NewCfnHostPropsMixin_Override(c CfnHostPropsMixin, props *CfnHostMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnHostPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnHostPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnHostPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnHostPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHostPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnHostPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHostPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHostPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

