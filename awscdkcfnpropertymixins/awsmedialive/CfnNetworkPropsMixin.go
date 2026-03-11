package awsmedialive

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::MediaLive::Network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnNetworkPropsMixin := awscdkcfnpropertymixins.Aws_medialive.NewCfnNetworkPropsMixin(&CfnNetworkMixinProps{
//   	IpPools: []interface{}{
//   		&IpPoolProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Routes: []interface{}{
//   		&RouteProperty{
//   			Cidr: jsii.String("cidr"),
//   			Gateway: jsii.String("gateway"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-network.html
//
type CfnNetworkPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnNetworkMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNetworkPropsMixin
type jsiiProxy_CfnNetworkPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnNetworkPropsMixin) Props() *CfnNetworkMixinProps {
	var returns *CfnNetworkMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNetworkPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaLive::Network`.
func NewCfnNetworkPropsMixin(props *CfnNetworkMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnNetworkPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNetworkPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNetworkPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_medialive.CfnNetworkPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaLive::Network`.
func NewCfnNetworkPropsMixin_Override(c CfnNetworkPropsMixin, props *CfnNetworkMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_medialive.CfnNetworkPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnNetworkPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNetworkPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_medialive.CfnNetworkPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNetworkPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_medialive.CfnNetworkPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNetworkPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnNetworkPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

