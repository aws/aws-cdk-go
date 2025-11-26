package previewawsschedulermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsscheduler/previewawsschedulermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A *schedule* is the main resource you create, configure, and manage using Amazon EventBridge Scheduler.
//
// Every schedule has a *schedule expression* that determines when, and with what frequency, the schedule runs. EventBridge Scheduler supports three types of schedules: rate, cron, and one-time schedules. For more information about different schedule types, see [Schedule types](https://docs.aws.amazon.com/scheduler/latest/UserGuide/schedule-types.html) in the *EventBridge Scheduler User Guide* .
//
// When you create a schedule, you configure a target for the schedule to invoke. A target is an API operation that EventBridge Scheduler calls on your behalf every time your schedule runs. EventBridge Scheduler supports two types of targets: *templated* targets invoke common API operations across a core groups of services, and customizeable *universal* targets that you can use to call more than 6,000 operations across over 270 services. For more information about configuring targets, see [Managing targets](https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-targets.html) in the *EventBridge Scheduler User Guide* .
//
// For more information about managing schedules, changing the schedule state, setting up flexible time windows, and configuring a dead-letter queue for a schedule, see [Managing a schedule](https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-schedule.html) in the *EventBridge Scheduler User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   cfnSchedulePropsMixin := awscdkmixinspreview.Mixins.NewCfnSchedulePropsMixin(&CfnScheduleMixinProps{
//   	Description: jsii.String("description"),
//   	EndDate: jsii.String("endDate"),
//   	FlexibleTimeWindow: &FlexibleTimeWindowProperty{
//   		MaximumWindowInMinutes: jsii.Number(123),
//   		Mode: jsii.String("mode"),
//   	},
//   	GroupName: jsii.String("groupName"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Name: jsii.String("name"),
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   	ScheduleExpressionTimezone: jsii.String("scheduleExpressionTimezone"),
//   	StartDate: jsii.String("startDate"),
//   	State: jsii.String("state"),
//   	Target: &TargetProperty{
//   		Arn: jsii.String("arn"),
//   		DeadLetterConfig: &DeadLetterConfigProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		EcsParameters: &EcsParametersProperty{
//   			CapacityProviderStrategy: []interface{}{
//   				&CapacityProviderStrategyItemProperty{
//   					Base: jsii.Number(123),
//   					CapacityProvider: jsii.String("capacityProvider"),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   			EnableEcsManagedTags: jsii.Boolean(false),
//   			EnableExecuteCommand: jsii.Boolean(false),
//   			Group: jsii.String("group"),
//   			LaunchType: jsii.String("launchType"),
//   			NetworkConfiguration: &NetworkConfigurationProperty{
//   				AwsvpcConfiguration: &AwsVpcConfigurationProperty{
//   					AssignPublicIp: jsii.String("assignPublicIp"),
//   					SecurityGroups: []*string{
//   						jsii.String("securityGroups"),
//   					},
//   					Subnets: []*string{
//   						jsii.String("subnets"),
//   					},
//   				},
//   			},
//   			PlacementConstraints: []interface{}{
//   				&PlacementConstraintProperty{
//   					Expression: jsii.String("expression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			PlacementStrategy: []interface{}{
//   				&PlacementStrategyProperty{
//   					Field: jsii.String("field"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			PlatformVersion: jsii.String("platformVersion"),
//   			PropagateTags: jsii.String("propagateTags"),
//   			ReferenceId: jsii.String("referenceId"),
//   			Tags: tags,
//   			TaskCount: jsii.Number(123),
//   			TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//   		},
//   		EventBridgeParameters: &EventBridgeParametersProperty{
//   			DetailType: jsii.String("detailType"),
//   			Source: jsii.String("source"),
//   		},
//   		Input: jsii.String("input"),
//   		KinesisParameters: &KinesisParametersProperty{
//   			PartitionKey: jsii.String("partitionKey"),
//   		},
//   		RetryPolicy: &RetryPolicyProperty{
//   			MaximumEventAgeInSeconds: jsii.Number(123),
//   			MaximumRetryAttempts: jsii.Number(123),
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   		SageMakerPipelineParameters: &SageMakerPipelineParametersProperty{
//   			PipelineParameterList: []interface{}{
//   				&SageMakerPipelineParameterProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		SqsParameters: &SqsParametersProperty{
//   			MessageGroupId: jsii.String("messageGroupId"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scheduler-schedule.html
//
type CfnSchedulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnScheduleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSchedulePropsMixin
type jsiiProxy_CfnSchedulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSchedulePropsMixin) Props() *CfnScheduleMixinProps {
	var returns *CfnScheduleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSchedulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Scheduler::Schedule`.
func NewCfnSchedulePropsMixin(props *CfnScheduleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSchedulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSchedulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSchedulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Scheduler::Schedule`.
func NewCfnSchedulePropsMixin_Override(c CfnSchedulePropsMixin, props *CfnScheduleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSchedulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSchedulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSchedulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSchedulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSchedulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

