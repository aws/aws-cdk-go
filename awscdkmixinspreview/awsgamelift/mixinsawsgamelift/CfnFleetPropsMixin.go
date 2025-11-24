package mixinsawsgamelift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsgamelift/mixinsawsgamelift/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::GameLift::Fleet` resource creates an Amazon GameLift (GameLift) fleet to host custom game server or Realtime Servers.
//
// A fleet is a set of EC2 instances, configured with instructions to run game servers on each instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFleetPropsMixin := awscdkmixinspreview.Mixins.NewCfnFleetPropsMixin(&CfnFleetMixinProps{
//   	AnywhereConfiguration: &AnywhereConfigurationProperty{
//   		Cost: jsii.String("cost"),
//   	},
//   	ApplyCapacity: jsii.String("applyCapacity"),
//   	BuildId: jsii.String("buildId"),
//   	CertificateConfiguration: &CertificateConfigurationProperty{
//   		CertificateType: jsii.String("certificateType"),
//   	},
//   	ComputeType: jsii.String("computeType"),
//   	Description: jsii.String("description"),
//   	DesiredEc2Instances: jsii.Number(123),
//   	Ec2InboundPermissions: []interface{}{
//   		&IpPermissionProperty{
//   			FromPort: jsii.Number(123),
//   			IpRange: jsii.String("ipRange"),
//   			Protocol: jsii.String("protocol"),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	Ec2InstanceType: jsii.String("ec2InstanceType"),
//   	FleetType: jsii.String("fleetType"),
//   	InstanceRoleArn: jsii.String("instanceRoleArn"),
//   	InstanceRoleCredentialsProvider: jsii.String("instanceRoleCredentialsProvider"),
//   	Locations: []interface{}{
//   		&LocationConfigurationProperty{
//   			Location: jsii.String("location"),
//   			LocationCapacity: &LocationCapacityProperty{
//   				DesiredEc2Instances: jsii.Number(123),
//   				MaxSize: jsii.Number(123),
//   				MinSize: jsii.Number(123),
//   			},
//   		},
//   	},
//   	LogPaths: []*string{
//   		jsii.String("logPaths"),
//   	},
//   	MaxSize: jsii.Number(123),
//   	MetricGroups: []*string{
//   		jsii.String("metricGroups"),
//   	},
//   	MinSize: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	NewGameSessionProtectionPolicy: jsii.String("newGameSessionProtectionPolicy"),
//   	PeerVpcAwsAccountId: jsii.String("peerVpcAwsAccountId"),
//   	PeerVpcId: jsii.String("peerVpcId"),
//   	ResourceCreationLimitPolicy: &ResourceCreationLimitPolicyProperty{
//   		NewGameSessionsPerCreator: jsii.Number(123),
//   		PolicyPeriodInMinutes: jsii.Number(123),
//   	},
//   	RuntimeConfiguration: &RuntimeConfigurationProperty{
//   		GameSessionActivationTimeoutSeconds: jsii.Number(123),
//   		MaxConcurrentGameSessionActivations: jsii.Number(123),
//   		ServerProcesses: []interface{}{
//   			&ServerProcessProperty{
//   				ConcurrentExecutions: jsii.Number(123),
//   				LaunchPath: jsii.String("launchPath"),
//   				Parameters: jsii.String("parameters"),
//   			},
//   		},
//   	},
//   	ScalingPolicies: []interface{}{
//   		&ScalingPolicyProperty{
//   			ComparisonOperator: jsii.String("comparisonOperator"),
//   			EvaluationPeriods: jsii.Number(123),
//   			Location: jsii.String("location"),
//   			MetricName: jsii.String("metricName"),
//   			Name: jsii.String("name"),
//   			PolicyType: jsii.String("policyType"),
//   			ScalingAdjustment: jsii.Number(123),
//   			ScalingAdjustmentType: jsii.String("scalingAdjustmentType"),
//   			Status: jsii.String("status"),
//   			TargetConfiguration: &TargetConfigurationProperty{
//   				TargetValue: jsii.Number(123),
//   			},
//   			Threshold: jsii.Number(123),
//   			UpdateStatus: jsii.String("updateStatus"),
//   		},
//   	},
//   	ScriptId: jsii.String("scriptId"),
//   	ServerLaunchParameters: jsii.String("serverLaunchParameters"),
//   	ServerLaunchPath: jsii.String("serverLaunchPath"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html
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


// Create a mixin to apply properties to `AWS::GameLift::Fleet`.
func NewCfnFleetPropsMixin(props *CfnFleetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFleetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFleetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFleetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GameLift::Fleet`.
func NewCfnFleetPropsMixin_Override(c CfnFleetPropsMixin, props *CfnFleetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnFleetPropsMixin",
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

