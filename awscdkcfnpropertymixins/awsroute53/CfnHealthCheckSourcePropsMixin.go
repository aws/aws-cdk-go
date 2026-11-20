package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Route53::HealthCheckSource.
//
// The customer VPC, subnets and
// security groups from which Route 53 probes private endpoints (PEHC). Immutable after creation
// (no Update API).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnHealthCheckSourcePropsMixin := awscdkcfnpropertymixins.Aws_route53.NewCfnHealthCheckSourcePropsMixin(&CfnHealthCheckSourceMixinProps{
//   	HealthCheckSourceConfig: &HealthCheckSourceConfigProperty{
//   		SubnetConfigurationList: []interface{}{
//   			&SubnetConfigurationProperty{
//   				AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   				SecurityGroupIdList: []*string{
//   					jsii.String("securityGroupIdList"),
//   				},
//   				SubnetId: jsii.String("subnetId"),
//   			},
//   		},
//   		VpcId: jsii.String("vpcId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthchecksource.html
//
type CfnHealthCheckSourcePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnHealthCheckSourceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnHealthCheckSourcePropsMixin
type jsiiProxy_CfnHealthCheckSourcePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnHealthCheckSourcePropsMixin) Props() *CfnHealthCheckSourceMixinProps {
	var returns *CfnHealthCheckSourceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHealthCheckSourcePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53::HealthCheckSource`.
func NewCfnHealthCheckSourcePropsMixin(props *CfnHealthCheckSourceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnHealthCheckSourcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnHealthCheckSourcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnHealthCheckSourcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHealthCheckSourcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53::HealthCheckSource`.
func NewCfnHealthCheckSourcePropsMixin_Override(c CfnHealthCheckSourcePropsMixin, props *CfnHealthCheckSourceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHealthCheckSourcePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnHealthCheckSourcePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnHealthCheckSourcePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHealthCheckSourcePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHealthCheckSourcePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHealthCheckSourcePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHealthCheckSourcePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnHealthCheckSourcePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

