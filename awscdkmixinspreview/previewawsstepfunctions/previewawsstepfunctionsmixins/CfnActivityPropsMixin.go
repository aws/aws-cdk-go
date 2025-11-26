package previewawsstepfunctionsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsstepfunctions/previewawsstepfunctionsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An activity is a task that you write in any programming language and host on any machine that has access to AWS Step Functions .
//
// Activities must poll Step Functions using the `GetActivityTask` API action and respond using `SendTask*` API actions. This function makes Step Functions aware of your activity and returns an identifier for use in a state machine and when polling from the activity.
//
// For information about creating an activity, see [Creating an Activity State Machine](https://docs.aws.amazon.com/step-functions/latest/dg/tutorial-creating-activity-state-machine.html) in the *AWS Step Functions Developer Guide* and [CreateActivity](https://docs.aws.amazon.com/step-functions/latest/apireference/API_CreateActivity.html) in the *AWS Step Functions API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnActivityPropsMixin := awscdkmixinspreview.Mixins.NewCfnActivityPropsMixin(&CfnActivityMixinProps{
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsDataKeyReusePeriodSeconds: jsii.Number(123),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		Type: jsii.String("type"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []TagsEntryProperty{
//   		&TagsEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-activity.html
//
type CfnActivityPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnActivityMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnActivityPropsMixin
type jsiiProxy_CfnActivityPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnActivityPropsMixin) Props() *CfnActivityMixinProps {
	var returns *CfnActivityMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnActivityPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::StepFunctions::Activity`.
func NewCfnActivityPropsMixin(props *CfnActivityMixinProps, options *mixins.CfnPropertyMixinOptions) CfnActivityPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnActivityPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnActivityPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnActivityPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::StepFunctions::Activity`.
func NewCfnActivityPropsMixin_Override(c CfnActivityPropsMixin, props *CfnActivityMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnActivityPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnActivityPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnActivityPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnActivityPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnActivityPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnActivityPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnActivityPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnActivityPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

