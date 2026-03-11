package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a path to analyze for reachability.
//
// VPC Reachability Analyzer enables you to analyze and debug network reachability between two resources in your virtual private cloud (VPC). For more information, see the [Reachability Analyzer User Guide](https://docs.aws.amazon.com/vpc/latest/reachability/what-is-reachability-analyzer.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnNetworkInsightsPathPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnNetworkInsightsPathPropsMixin(&CfnNetworkInsightsPathMixinProps{
//   	Destination: jsii.String("destination"),
//   	DestinationIp: jsii.String("destinationIp"),
//   	DestinationPort: jsii.Number(123),
//   	FilterAtDestination: &PathFilterProperty{
//   		DestinationAddress: jsii.String("destinationAddress"),
//   		DestinationPortRange: &FilterPortRangeProperty{
//   			FromPort: jsii.Number(123),
//   			ToPort: jsii.Number(123),
//   		},
//   		SourceAddress: jsii.String("sourceAddress"),
//   		SourcePortRange: &FilterPortRangeProperty{
//   			FromPort: jsii.Number(123),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	FilterAtSource: &PathFilterProperty{
//   		DestinationAddress: jsii.String("destinationAddress"),
//   		DestinationPortRange: &FilterPortRangeProperty{
//   			FromPort: jsii.Number(123),
//   			ToPort: jsii.Number(123),
//   		},
//   		SourceAddress: jsii.String("sourceAddress"),
//   		SourcePortRange: &FilterPortRangeProperty{
//   			FromPort: jsii.Number(123),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	Protocol: jsii.String("protocol"),
//   	Source: jsii.String("source"),
//   	SourceIp: jsii.String("sourceIp"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinsightspath.html
//
type CfnNetworkInsightsPathPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnNetworkInsightsPathMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNetworkInsightsPathPropsMixin
type jsiiProxy_CfnNetworkInsightsPathPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnNetworkInsightsPathPropsMixin) Props() *CfnNetworkInsightsPathMixinProps {
	var returns *CfnNetworkInsightsPathMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNetworkInsightsPathPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::NetworkInsightsPath`.
func NewCfnNetworkInsightsPathPropsMixin(props *CfnNetworkInsightsPathMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnNetworkInsightsPathPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNetworkInsightsPathPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNetworkInsightsPathPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnNetworkInsightsPathPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::NetworkInsightsPath`.
func NewCfnNetworkInsightsPathPropsMixin_Override(c CfnNetworkInsightsPathPropsMixin, props *CfnNetworkInsightsPathMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnNetworkInsightsPathPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnNetworkInsightsPathPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNetworkInsightsPathPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnNetworkInsightsPathPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNetworkInsightsPathPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnNetworkInsightsPathPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNetworkInsightsPathPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnNetworkInsightsPathPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

