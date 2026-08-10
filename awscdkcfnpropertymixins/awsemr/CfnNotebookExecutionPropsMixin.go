package awsemr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Amazon EMR Notebook execution.
//
// A notebook execution is a specific instance that an Amazon EMR Notebook is run using the StartNotebookExecution action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnNotebookExecutionPropsMixin := awscdkcfnpropertymixins.Aws_emr.NewCfnNotebookExecutionPropsMixin(&CfnNotebookExecutionMixinProps{
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	ExecutionEngine: &ExecutionEngineConfigProperty{
//   		Id: jsii.String("id"),
//   		Type: jsii.String("type"),
//   	},
//   	NotebookExecutionName: jsii.String("notebookExecutionName"),
//   	NotebookParams: jsii.String("notebookParams"),
//   	NotebookS3Location: &NotebookS3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//   	OutputNotebookFormat: jsii.String("outputNotebookFormat"),
//   	OutputNotebookS3Location: &OutputNotebookS3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-notebookexecution.html
//
type CfnNotebookExecutionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnNotebookExecutionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNotebookExecutionPropsMixin
type jsiiProxy_CfnNotebookExecutionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnNotebookExecutionPropsMixin) Props() *CfnNotebookExecutionMixinProps {
	var returns *CfnNotebookExecutionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotebookExecutionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EMR::NotebookExecution`.
func NewCfnNotebookExecutionPropsMixin(props *CfnNotebookExecutionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnNotebookExecutionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNotebookExecutionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNotebookExecutionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnNotebookExecutionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EMR::NotebookExecution`.
func NewCfnNotebookExecutionPropsMixin_Override(c CfnNotebookExecutionPropsMixin, props *CfnNotebookExecutionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnNotebookExecutionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnNotebookExecutionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNotebookExecutionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnNotebookExecutionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNotebookExecutionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnNotebookExecutionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNotebookExecutionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnNotebookExecutionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

