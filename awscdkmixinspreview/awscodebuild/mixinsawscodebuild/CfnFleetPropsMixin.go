package mixinsawscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awscodebuild/mixinsawscodebuild/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CodeBuild::Fleet` resource configures a compute fleet, a set of dedicated instances for your build environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFleetPropsMixin := awscdkmixinspreview.Mixins.NewCfnFleetPropsMixin(&CfnFleetMixinProps{
//   	BaseCapacity: jsii.Number(123),
//   	ComputeConfiguration: &ComputeConfigurationProperty{
//   		Disk: jsii.Number(123),
//   		InstanceType: jsii.String("instanceType"),
//   		MachineType: jsii.String("machineType"),
//   		Memory: jsii.Number(123),
//   		VCpu: jsii.Number(123),
//   	},
//   	ComputeType: jsii.String("computeType"),
//   	EnvironmentType: jsii.String("environmentType"),
//   	FleetProxyConfiguration: &ProxyConfigurationProperty{
//   		DefaultBehavior: jsii.String("defaultBehavior"),
//   		OrderedProxyRules: []interface{}{
//   			&FleetProxyRuleProperty{
//   				Effect: jsii.String("effect"),
//   				Entities: []*string{
//   					jsii.String("entities"),
//   				},
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	FleetServiceRole: jsii.String("fleetServiceRole"),
//   	FleetVpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//   	ImageId: jsii.String("imageId"),
//   	Name: jsii.String("name"),
//   	OverflowBehavior: jsii.String("overflowBehavior"),
//   	ScalingConfiguration: &ScalingConfigurationInputProperty{
//   		MaxCapacity: jsii.Number(123),
//   		ScalingType: jsii.String("scalingType"),
//   		TargetTrackingScalingConfigs: []interface{}{
//   			&TargetTrackingScalingConfigurationProperty{
//   				MetricType: jsii.String("metricType"),
//   				TargetValue: jsii.Number(123),
//   			},
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-fleet.html
//
type CfnFleetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFleetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFleetPropsMixin
type jsiiProxy_CfnFleetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFleetPropsMixin) Props() *CfnFleetMixinProps {
	var returns *CfnFleetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFleetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CodeBuild::Fleet`.
func NewCfnFleetPropsMixin(props *CfnFleetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFleetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFleetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFleetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnFleetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CodeBuild::Fleet`.
func NewCfnFleetPropsMixin_Override(c CfnFleetPropsMixin, props *CfnFleetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnFleetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFleetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFleetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnFleetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFleetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnFleetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFleetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnFleetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

