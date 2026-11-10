package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS ECS Gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnGatewayPropsMixin := awscdkcfnpropertymixins.Aws_ecs.NewCfnGatewayPropsMixin(&CfnGatewayMixinProps{
//   	Configuration: &GatewayConfigurationProperty{
//   		HostSharedAlb: &HostSharedAlbConfigurationProperty{
//   			AccessLogs: &GatewayElbAccessLogConfigurationProperty{
//   				Destination: &GatewayElbAccessLogDestinationProperty{
//   					S3: &GatewayElbAccessLogS3DestinationProperty{
//   						Bucket: jsii.String("bucket"),
//   						Prefix: jsii.String("prefix"),
//   					},
//   				},
//   			},
//   			CurrentAlbCount: jsii.Number(123),
//   			HttpDesync: jsii.String("httpDesync"),
//   			MaxAlbCount: jsii.Number(123),
//   			NetworkConfiguration: &GatewayElbNetworkConfigurationProperty{
//   				InboundSecurityGroup: jsii.String("inboundSecurityGroup"),
//   				IpAddressType: jsii.String("ipAddressType"),
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   			Port: jsii.Number(123),
//   			Protocol: jsii.String("protocol"),
//   			Scheme: jsii.String("scheme"),
//   			TestTrafficRules: &GatewayAlbTestTrafficRulesProperty{
//   				Header: &GatewayTestTrafficHeaderProperty{
//   					Name: jsii.String("name"),
//   					Values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			WebAclArn: jsii.String("webAclArn"),
//   		},
//   	},
//   	DeletionProtection: jsii.Boolean(false),
//   	GatewayName: jsii.String("gatewayName"),
//   	InfrastructureRoleArn: jsii.String("infrastructureRoleArn"),
//   	PropagateTags: jsii.String("propagateTags"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-gateway.html
//
type CfnGatewayPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnGatewayMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGatewayPropsMixin
type jsiiProxy_CfnGatewayPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnGatewayPropsMixin) Props() *CfnGatewayMixinProps {
	var returns *CfnGatewayMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGatewayPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ECS::Gateway`.
func NewCfnGatewayPropsMixin(props *CfnGatewayMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnGatewayPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGatewayPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGatewayPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnGatewayPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ECS::Gateway`.
func NewCfnGatewayPropsMixin_Override(c CfnGatewayPropsMixin, props *CfnGatewayMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnGatewayPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnGatewayPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGatewayPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnGatewayPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGatewayPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ecs.CfnGatewayPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGatewayPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnGatewayPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

