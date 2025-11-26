package previewawsstepfunctionsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsstepfunctions/previewawsstepfunctionsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Provisions a state machine.
//
// A state machine consists of a collection of states that can do work ( `Task` states), determine to which states to transition next ( `Choice` states), stop an execution with an error ( `Fail` states), and so on. State machines are specified using a JSON-based, structured language.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var definition interface{}
//
//   cfnStateMachinePropsMixin := awscdkmixinspreview.Mixins.NewCfnStateMachinePropsMixin(&CfnStateMachineMixinProps{
//   	Definition: definition,
//   	DefinitionS3Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   		Version: jsii.String("version"),
//   	},
//   	DefinitionString: jsii.String("definitionString"),
//   	DefinitionSubstitutions: map[string]*string{
//   		"definitionSubstitutionsKey": jsii.String("definitionSubstitutions"),
//   	},
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsDataKeyReusePeriodSeconds: jsii.Number(123),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		Type: jsii.String("type"),
//   	},
//   	LoggingConfiguration: &LoggingConfigurationProperty{
//   		Destinations: []interface{}{
//   			&LogDestinationProperty{
//   				CloudWatchLogsLogGroup: &CloudWatchLogsLogGroupProperty{
//   					LogGroupArn: jsii.String("logGroupArn"),
//   				},
//   			},
//   		},
//   		IncludeExecutionData: jsii.Boolean(false),
//   		Level: jsii.String("level"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	StateMachineName: jsii.String("stateMachineName"),
//   	StateMachineType: jsii.String("stateMachineType"),
//   	Tags: []TagsEntryProperty{
//   		&TagsEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TracingConfiguration: &TracingConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html
//
type CfnStateMachinePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnStateMachineMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStateMachinePropsMixin
type jsiiProxy_CfnStateMachinePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnStateMachinePropsMixin) Props() *CfnStateMachineMixinProps {
	var returns *CfnStateMachineMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachinePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::StepFunctions::StateMachine`.
func NewCfnStateMachinePropsMixin(props *CfnStateMachineMixinProps, options *mixins.CfnPropertyMixinOptions) CfnStateMachinePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnStateMachinePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStateMachinePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachinePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::StepFunctions::StateMachine`.
func NewCfnStateMachinePropsMixin_Override(c CfnStateMachinePropsMixin, props *CfnStateMachineMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachinePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnStateMachinePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStateMachinePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachinePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStateMachinePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachinePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStateMachinePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnStateMachinePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

