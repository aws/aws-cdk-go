package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Flex Cost Allocation (FCA) policy defines how shared network costs (data transfer, data processing) are allocated and reported for network resources such as NAT gateways, EC2 instances, and VPC peering connections.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnFcaPolicyPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnFcaPolicyPropsMixin(&CfnFcaPolicyMixinProps{
//   	Description: jsii.String("description"),
//   	InstanceFcaPolicyDocument: &InstanceFcaPolicyDocumentProperty{
//   		CidrRangeSpecifications: &CidrRangeSpecificationsProperty{
//   			Prefixes: []interface{}{
//   				&CidrMappingProperty{
//   					CidrLabel: jsii.String("cidrLabel"),
//   					CidrRange: jsii.String("cidrRange"),
//   				},
//   			},
//   			PrefixesType: jsii.String("prefixesType"),
//   		},
//   		Dimensions: []*string{
//   			jsii.String("dimensions"),
//   		},
//   		ReportingGranularity: []*string{
//   			jsii.String("reportingGranularity"),
//   		},
//   	},
//   	NatFcaPolicyDocument: &NatFcaPolicyDocumentProperty{
//   		CidrRangeSpecifications: &CidrRangeSpecificationsProperty{
//   			Prefixes: []interface{}{
//   				&CidrMappingProperty{
//   					CidrLabel: jsii.String("cidrLabel"),
//   					CidrRange: jsii.String("cidrRange"),
//   				},
//   			},
//   			PrefixesType: jsii.String("prefixesType"),
//   		},
//   		Dimensions: []*string{
//   			jsii.String("dimensions"),
//   		},
//   		ReportingGranularity: []*string{
//   			jsii.String("reportingGranularity"),
//   		},
//   	},
//   	NetworkResourceType: jsii.String("networkResourceType"),
//   	PcxFcaPolicyDocument: &PcxFcaPolicyDocumentProperty{
//   		CidrRangeSpecifications: &CidrRangeSpecificationsProperty{
//   			Prefixes: []interface{}{
//   				&CidrMappingProperty{
//   					CidrLabel: jsii.String("cidrLabel"),
//   					CidrRange: jsii.String("cidrRange"),
//   				},
//   			},
//   			PrefixesType: jsii.String("prefixesType"),
//   		},
//   		Dimensions: []*string{
//   			jsii.String("dimensions"),
//   		},
//   		ReportingGranularity: []*string{
//   			jsii.String("reportingGranularity"),
//   		},
//   	},
//   	ShowBackEnabled: jsii.Boolean(false),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-fcapolicy.html
//
type CfnFcaPolicyPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnFcaPolicyMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFcaPolicyPropsMixin
type jsiiProxy_CfnFcaPolicyPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnFcaPolicyPropsMixin) Props() *CfnFcaPolicyMixinProps {
	var returns *CfnFcaPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFcaPolicyPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::FcaPolicy`.
func NewCfnFcaPolicyPropsMixin(props *CfnFcaPolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnFcaPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFcaPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFcaPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnFcaPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::FcaPolicy`.
func NewCfnFcaPolicyPropsMixin_Override(c CfnFcaPolicyPropsMixin, props *CfnFcaPolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnFcaPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnFcaPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFcaPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnFcaPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFcaPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnFcaPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFcaPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFcaPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

