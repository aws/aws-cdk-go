package previewawsgameliftmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgamelift/previewawsgameliftmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Describes an Amazon GameLift Servers managed container fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnContainerFleetPropsMixin := awscdkmixinspreview.Mixins.NewCfnContainerFleetPropsMixin(&CfnContainerFleetMixinProps{
//   	BillingType: jsii.String("billingType"),
//   	DeploymentConfiguration: &DeploymentConfigurationProperty{
//   		ImpairmentStrategy: jsii.String("impairmentStrategy"),
//   		MinimumHealthyPercentage: jsii.Number(123),
//   		ProtectionStrategy: jsii.String("protectionStrategy"),
//   	},
//   	Description: jsii.String("description"),
//   	FleetRoleArn: jsii.String("fleetRoleArn"),
//   	GameServerContainerGroupDefinitionName: jsii.String("gameServerContainerGroupDefinitionName"),
//   	GameServerContainerGroupsPerInstance: jsii.Number(123),
//   	GameSessionCreationLimitPolicy: &GameSessionCreationLimitPolicyProperty{
//   		NewGameSessionsPerCreator: jsii.Number(123),
//   		PolicyPeriodInMinutes: jsii.Number(123),
//   	},
//   	InstanceConnectionPortRange: &ConnectionPortRangeProperty{
//   		FromPort: jsii.Number(123),
//   		ToPort: jsii.Number(123),
//   	},
//   	InstanceInboundPermissions: []interface{}{
//   		&IpPermissionProperty{
//   			FromPort: jsii.Number(123),
//   			IpRange: jsii.String("ipRange"),
//   			Protocol: jsii.String("protocol"),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	InstanceType: jsii.String("instanceType"),
//   	Locations: []interface{}{
//   		&LocationConfigurationProperty{
//   			Location: jsii.String("location"),
//   			LocationCapacity: &LocationCapacityProperty{
//   				DesiredEc2Instances: jsii.Number(123),
//   				ManagedCapacityConfiguration: &ManagedCapacityConfigurationProperty{
//   					ScaleInAfterInactivityMinutes: jsii.Number(123),
//   					ZeroCapacityStrategy: jsii.String("zeroCapacityStrategy"),
//   				},
//   				MaxSize: jsii.Number(123),
//   				MinSize: jsii.Number(123),
//   			},
//   			StoppedActions: []*string{
//   				jsii.String("stoppedActions"),
//   			},
//   		},
//   	},
//   	LogConfiguration: &LogConfigurationProperty{
//   		LogDestination: jsii.String("logDestination"),
//   		LogGroupArn: jsii.String("logGroupArn"),
//   		S3BucketName: jsii.String("s3BucketName"),
//   	},
//   	MetricGroups: []*string{
//   		jsii.String("metricGroups"),
//   	},
//   	NewGameSessionProtectionPolicy: jsii.String("newGameSessionProtectionPolicy"),
//   	PerInstanceContainerGroupDefinitionName: jsii.String("perInstanceContainerGroupDefinitionName"),
//   	ScalingPolicies: []interface{}{
//   		&ScalingPolicyProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			EvaluationPeriods: jsii.Number(123),
//   			MetricName: jsii.String("metricName"),
//   			Name: jsii.String("name"),
//   			PolicyType: jsii.String("policyType"),
//   			ScalingAdjustment: jsii.Number(123),
//   			ScalingAdjustmentType: jsii.String("scalingAdjustmentType"),
//   			TargetConfiguration: &TargetConfigurationProperty{
//   				TargetValue: jsii.Number(123),
//   			},
//   			Threshold: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-containerfleet.html
//
type CfnContainerFleetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnContainerFleetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnContainerFleetPropsMixin
type jsiiProxy_CfnContainerFleetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnContainerFleetPropsMixin) Props() *CfnContainerFleetMixinProps {
	var returns *CfnContainerFleetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerFleetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GameLift::ContainerFleet`.
func NewCfnContainerFleetPropsMixin(props *CfnContainerFleetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnContainerFleetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnContainerFleetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnContainerFleetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GameLift::ContainerFleet`.
func NewCfnContainerFleetPropsMixin_Override(c CfnContainerFleetPropsMixin, props *CfnContainerFleetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnContainerFleetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerFleetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContainerFleetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnContainerFleetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContainerFleetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnContainerFleetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

