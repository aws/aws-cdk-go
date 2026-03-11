package awsmediaconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a router network interface in AWS Elemental MediaConnect that is used to define a network boundary for router resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnRouterNetworkInterfacePropsMixin := awscdkcfnpropertymixins.Aws_mediaconnect.NewCfnRouterNetworkInterfacePropsMixin(&CfnRouterNetworkInterfaceMixinProps{
//   	Configuration: &RouterNetworkInterfaceConfigurationProperty{
//   		Public: &PublicRouterNetworkInterfaceConfigurationProperty{
//   			AllowRules: []interface{}{
//   				&PublicRouterNetworkInterfaceRuleProperty{
//   					Cidr: jsii.String("cidr"),
//   				},
//   			},
//   		},
//   		Vpc: &VpcRouterNetworkInterfaceConfigurationProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RegionName: jsii.String("regionName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routernetworkinterface.html
//
type CfnRouterNetworkInterfacePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRouterNetworkInterfaceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRouterNetworkInterfacePropsMixin
type jsiiProxy_CfnRouterNetworkInterfacePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRouterNetworkInterfacePropsMixin) Props() *CfnRouterNetworkInterfaceMixinProps {
	var returns *CfnRouterNetworkInterfaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouterNetworkInterfacePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaConnect::RouterNetworkInterface`.
func NewCfnRouterNetworkInterfacePropsMixin(props *CfnRouterNetworkInterfaceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRouterNetworkInterfacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRouterNetworkInterfacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRouterNetworkInterfacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterNetworkInterfacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaConnect::RouterNetworkInterface`.
func NewCfnRouterNetworkInterfacePropsMixin_Override(c CfnRouterNetworkInterfacePropsMixin, props *CfnRouterNetworkInterfaceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterNetworkInterfacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRouterNetworkInterfacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRouterNetworkInterfacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterNetworkInterfacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRouterNetworkInterfacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mediaconnect.CfnRouterNetworkInterfacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRouterNetworkInterfacePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRouterNetworkInterfacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

