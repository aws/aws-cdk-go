package awsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new VPC connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnVPCConnectionPropsMixin := awscdkcfnpropertymixins.Aws_quicksight.NewCfnVPCConnectionPropsMixin(&CfnVPCConnectionMixinProps{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	DnsResolvers: []*string{
//   		jsii.String("dnsResolvers"),
//   	},
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConnectionId: jsii.String("vpcConnectionId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-vpcconnection.html
//
type CfnVPCConnectionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnVPCConnectionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVPCConnectionPropsMixin
type jsiiProxy_CfnVPCConnectionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnVPCConnectionPropsMixin) Props() *CfnVPCConnectionMixinProps {
	var returns *CfnVPCConnectionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCConnectionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::QuickSight::VPCConnection`.
func NewCfnVPCConnectionPropsMixin(props *CfnVPCConnectionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnVPCConnectionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVPCConnectionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVPCConnectionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnVPCConnectionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::QuickSight::VPCConnection`.
func NewCfnVPCConnectionPropsMixin_Override(c CfnVPCConnectionPropsMixin, props *CfnVPCConnectionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnVPCConnectionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnVPCConnectionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVPCConnectionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnVPCConnectionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVPCConnectionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_quicksight.CfnVPCConnectionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVPCConnectionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnVPCConnectionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

