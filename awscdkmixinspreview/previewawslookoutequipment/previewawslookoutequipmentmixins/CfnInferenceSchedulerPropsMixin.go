package previewawslookoutequipmentmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslookoutequipment/previewawslookoutequipmentmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a scheduled inference.
//
// Scheduling an inference is setting up a continuous real-time inference plan to analyze new measurement data. When setting up the schedule, you provide an Amazon S3 bucket location for the input data, assign it a delimiter between separate entries in the data, set an offset delay if desired, and set the frequency of inferencing. You must also provide an Amazon S3 bucket location for the output data.
//
// > Updating some properties below (for example, InferenceSchedulerName and ServerSideKmsKeyId) triggers a resource replacement, which requires a new model. To replace such a property using CloudFormation , but without creating a completely new stack, you must replace ModelName. If you need to replace the property, but want to use the same model, delete the current stack and create a new one with the updated properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var dataInputConfiguration interface{}
//   var dataOutputConfiguration interface{}
//
//   cfnInferenceSchedulerPropsMixin := awscdkmixinspreview.Mixins.NewCfnInferenceSchedulerPropsMixin(&CfnInferenceSchedulerMixinProps{
//   	DataDelayOffsetInMinutes: jsii.Number(123),
//   	DataInputConfiguration: dataInputConfiguration,
//   	DataOutputConfiguration: dataOutputConfiguration,
//   	DataUploadFrequency: jsii.String("dataUploadFrequency"),
//   	InferenceSchedulerName: jsii.String("inferenceSchedulerName"),
//   	ModelName: jsii.String("modelName"),
//   	RoleArn: jsii.String("roleArn"),
//   	ServerSideKmsKeyId: jsii.String("serverSideKmsKeyId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutequipment-inferencescheduler.html
//
type CfnInferenceSchedulerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnInferenceSchedulerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInferenceSchedulerPropsMixin
type jsiiProxy_CfnInferenceSchedulerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnInferenceSchedulerPropsMixin) Props() *CfnInferenceSchedulerMixinProps {
	var returns *CfnInferenceSchedulerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInferenceSchedulerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::LookoutEquipment::InferenceScheduler`.
func NewCfnInferenceSchedulerPropsMixin(props *CfnInferenceSchedulerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInferenceSchedulerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInferenceSchedulerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInferenceSchedulerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lookoutequipment.mixins.CfnInferenceSchedulerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::LookoutEquipment::InferenceScheduler`.
func NewCfnInferenceSchedulerPropsMixin_Override(c CfnInferenceSchedulerPropsMixin, props *CfnInferenceSchedulerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lookoutequipment.mixins.CfnInferenceSchedulerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnInferenceSchedulerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInferenceSchedulerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lookoutequipment.mixins.CfnInferenceSchedulerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInferenceSchedulerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lookoutequipment.mixins.CfnInferenceSchedulerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInferenceSchedulerPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnInferenceSchedulerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

