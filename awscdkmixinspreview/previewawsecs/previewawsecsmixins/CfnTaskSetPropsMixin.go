package previewawsecsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsecs/previewawsecsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a task set in the specified cluster and service.
//
// This is used when a service uses the `EXTERNAL` deployment controller type. For more information, see [Amazon ECS deployment types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// > On March 21, 2024, a change was made to resolve the task definition revision before authorization. When a task definition revision is not specified, authorization will occur using the latest revision of a task definition.
//
// For information about the maximum number of task sets and other quotas, see [Amazon ECS service quotas](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-quotas.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTaskSetPropsMixin := awscdkmixinspreview.Mixins.NewCfnTaskSetPropsMixin(&CfnTaskSetMixinProps{
//   	CapacityProviderStrategy: []interface{}{
//   		&CapacityProviderStrategyItemProperty{
//   			Base: jsii.Number(123),
//   			CapacityProvider: jsii.String("capacityProvider"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	Cluster: jsii.String("cluster"),
//   	ExternalId: jsii.String("externalId"),
//   	LaunchType: jsii.String("launchType"),
//   	LoadBalancers: []interface{}{
//   		&LoadBalancerProperty{
//   			ContainerName: jsii.String("containerName"),
//   			ContainerPort: jsii.Number(123),
//   			TargetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		AwsVpcConfiguration: &AwsVpcConfigurationProperty{
//   			AssignPublicIp: jsii.String("assignPublicIp"),
//   			SecurityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	PlatformVersion: jsii.String("platformVersion"),
//   	Scale: &ScaleProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   	Service: jsii.String("service"),
//   	ServiceRegistries: []interface{}{
//   		&ServiceRegistryProperty{
//   			ContainerName: jsii.String("containerName"),
//   			ContainerPort: jsii.Number(123),
//   			Port: jsii.Number(123),
//   			RegistryArn: jsii.String("registryArn"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskDefinition: jsii.String("taskDefinition"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-taskset.html
//
type CfnTaskSetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTaskSetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTaskSetPropsMixin
type jsiiProxy_CfnTaskSetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTaskSetPropsMixin) Props() *CfnTaskSetMixinProps {
	var returns *CfnTaskSetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTaskSetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ECS::TaskSet`.
func NewCfnTaskSetPropsMixin(props *CfnTaskSetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTaskSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTaskSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTaskSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ECS::TaskSet`.
func NewCfnTaskSetPropsMixin_Override(c CfnTaskSetPropsMixin, props *CfnTaskSetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTaskSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTaskSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskSetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTaskSetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnTaskSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTaskSetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnTaskSetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

