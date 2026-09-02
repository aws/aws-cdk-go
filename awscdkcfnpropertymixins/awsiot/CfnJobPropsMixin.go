package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the AWS::IoT::Job resource to declare an AWS IoT job.
//
// A job can be used to define a set of remote operations that are sent to and run on one or more devices (things or thing groups) connected to AWS IoT.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnJobPropsMixin := awscdkcfnpropertymixins.Aws_iot.NewCfnJobPropsMixin(&CfnJobMixinProps{
//   	AbortConfig: &AbortConfigProperty{
//   		CriteriaList: []interface{}{
//   			&AbortCriteriaProperty{
//   				Action: jsii.String("action"),
//   				FailureType: jsii.String("failureType"),
//   				MinNumberOfExecutedThings: jsii.Number(123),
//   				ThresholdPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DestinationPackageVersions: []*string{
//   		jsii.String("destinationPackageVersions"),
//   	},
//   	Document: jsii.String("document"),
//   	DocumentParameters: map[string]*string{
//   		"documentParametersKey": jsii.String("documentParameters"),
//   	},
//   	DocumentSource: jsii.String("documentSource"),
//   	JobExecutionsRetryConfig: &JobExecutionsRetryConfigProperty{
//   		CriteriaList: []interface{}{
//   			&RetryCriteriaProperty{
//   				FailureType: jsii.String("failureType"),
//   				NumberOfRetries: jsii.Number(123),
//   			},
//   		},
//   	},
//   	JobExecutionsRolloutConfig: &JobExecutionsRolloutConfigProperty{
//   		ExponentialRate: &ExponentialRolloutRateProperty{
//   			BaseRatePerMinute: jsii.Number(123),
//   			IncrementFactor: jsii.Number(123),
//   			RateIncreaseCriteria: &RateIncreaseCriteriaProperty{
//   				NumberOfNotifiedThings: jsii.Number(123),
//   				NumberOfSucceededThings: jsii.Number(123),
//   			},
//   		},
//   		MaximumPerMinute: jsii.Number(123),
//   	},
//   	JobId: jsii.String("jobId"),
//   	JobTemplateArn: jsii.String("jobTemplateArn"),
//   	PresignedUrlConfig: &PresignedUrlConfigProperty{
//   		ExpiresInSec: jsii.Number(123),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	SchedulingConfig: &SchedulingConfigProperty{
//   		EndBehavior: jsii.String("endBehavior"),
//   		EndTime: jsii.String("endTime"),
//   		MaintenanceWindows: []interface{}{
//   			&MaintenanceWindowProperty{
//   				DurationInMinutes: jsii.Number(123),
//   				StartTime: jsii.String("startTime"),
//   			},
//   		},
//   		StartTime: jsii.String("startTime"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Targets: []*string{
//   		jsii.String("targets"),
//   	},
//   	TargetSelection: jsii.String("targetSelection"),
//   	TimeoutConfig: &TimeoutConfigProperty{
//   		InProgressTimeoutInMinutes: jsii.Number(123),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-job.html
//
type CfnJobPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnJobMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnJobPropsMixin
type jsiiProxy_CfnJobPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnJobPropsMixin) Props() *CfnJobMixinProps {
	var returns *CfnJobMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::Job`.
func NewCfnJobPropsMixin(props *CfnJobMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnJobPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnJobPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnJobPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnJobPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::Job`.
func NewCfnJobPropsMixin_Override(c CfnJobPropsMixin, props *CfnJobMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnJobPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnJobPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJobPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnJobPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJobPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnJobPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnJobPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnJobPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

