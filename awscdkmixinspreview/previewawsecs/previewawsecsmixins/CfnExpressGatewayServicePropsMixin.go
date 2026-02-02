package previewawsecsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsecs/previewawsecsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Express service that simplifies deploying containerized web applications on Amazon ECS with managed AWS infrastructure.
//
// This operation provisions and configures Application Load Balancers, target groups, security groups, and auto-scaling policies automatically.
//
// Specify a primary container configuration with your application image and basic settings. Amazon ECS creates the necessary AWS resources for traffic distribution, health monitoring, network access control, and capacity management.
//
// Provide an execution role for task operations and an infrastructure role for managing AWS resources on your behalf.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnExpressGatewayServicePropsMixin := awscdkmixinspreview.Mixins.NewCfnExpressGatewayServicePropsMixin(&CfnExpressGatewayServiceMixinProps{
//   	Cluster: jsii.String("cluster"),
//   	Cpu: jsii.String("cpu"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	HealthCheckPath: jsii.String("healthCheckPath"),
//   	InfrastructureRoleArn: jsii.String("infrastructureRoleArn"),
//   	Memory: jsii.String("memory"),
//   	NetworkConfiguration: &ExpressGatewayServiceNetworkConfigurationProperty{
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   	PrimaryContainer: &ExpressGatewayContainerProperty{
//   		AwsLogsConfiguration: &ExpressGatewayServiceAwsLogsConfigurationProperty{
//   			LogGroup: jsii.String("logGroup"),
//   			LogStreamPrefix: jsii.String("logStreamPrefix"),
//   		},
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		ContainerPort: jsii.Number(123),
//   		Environment: []interface{}{
//   			&KeyValuePairProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Image: jsii.String("image"),
//   		RepositoryCredentials: &ExpressGatewayRepositoryCredentialsProperty{
//   			CredentialsParameter: jsii.String("credentialsParameter"),
//   		},
//   		Secrets: []interface{}{
//   			&SecretProperty{
//   				Name: jsii.String("name"),
//   				ValueFrom: jsii.String("valueFrom"),
//   			},
//   		},
//   	},
//   	ScalingTarget: &ExpressGatewayScalingTargetProperty{
//   		AutoScalingMetric: jsii.String("autoScalingMetric"),
//   		AutoScalingTargetValue: jsii.Number(123),
//   		MaxTaskCount: jsii.Number(123),
//   		MinTaskCount: jsii.Number(123),
//   	},
//   	ServiceName: jsii.String("serviceName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskRoleArn: jsii.String("taskRoleArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-expressgatewayservice.html
//
type CfnExpressGatewayServicePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnExpressGatewayServiceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnExpressGatewayServicePropsMixin
type jsiiProxy_CfnExpressGatewayServicePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnExpressGatewayServicePropsMixin) Props() *CfnExpressGatewayServiceMixinProps {
	var returns *CfnExpressGatewayServiceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExpressGatewayServicePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ECS::ExpressGatewayService`.
func NewCfnExpressGatewayServicePropsMixin(props *CfnExpressGatewayServiceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnExpressGatewayServicePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnExpressGatewayServicePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnExpressGatewayServicePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ECS::ExpressGatewayService`.
func NewCfnExpressGatewayServicePropsMixin_Override(c CfnExpressGatewayServicePropsMixin, props *CfnExpressGatewayServiceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnExpressGatewayServicePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnExpressGatewayServicePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnExpressGatewayServicePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnExpressGatewayServicePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnExpressGatewayServicePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnExpressGatewayServicePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

