package mixinsawstransfer

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awstransfer/mixinsawstransfer/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Allows you to create a workflow with specified steps and step details the workflow invokes after file transfer completes.
//
// After creating a workflow, you can associate the workflow created with any transfer servers by specifying the `workflow-details` field in `CreateServer` and `UpdateServer` operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var copyStepDetails interface{}
//   var customStepDetails interface{}
//   var deleteStepDetails interface{}
//   var tagStepDetails interface{}
//
//   cfnWorkflowPropsMixin := awscdkmixinspreview.Mixins.NewCfnWorkflowPropsMixin(&CfnWorkflowMixinProps{
//   	Description: jsii.String("description"),
//   	OnExceptionSteps: []interface{}{
//   		&WorkflowStepProperty{
//   			CopyStepDetails: copyStepDetails,
//   			CustomStepDetails: customStepDetails,
//   			DecryptStepDetails: &DecryptStepDetailsProperty{
//   				DestinationFileLocation: &InputFileLocationProperty{
//   					EfsFileLocation: &EfsInputFileLocationProperty{
//   						FileSystemId: jsii.String("fileSystemId"),
//   						Path: jsii.String("path"),
//   					},
//   					S3FileLocation: &S3InputFileLocationProperty{
//   						Bucket: jsii.String("bucket"),
//   						Key: jsii.String("key"),
//   					},
//   				},
//   				Name: jsii.String("name"),
//   				OverwriteExisting: jsii.String("overwriteExisting"),
//   				SourceFileLocation: jsii.String("sourceFileLocation"),
//   				Type: jsii.String("type"),
//   			},
//   			DeleteStepDetails: deleteStepDetails,
//   			TagStepDetails: tagStepDetails,
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Steps: []interface{}{
//   		&WorkflowStepProperty{
//   			CopyStepDetails: copyStepDetails,
//   			CustomStepDetails: customStepDetails,
//   			DecryptStepDetails: &DecryptStepDetailsProperty{
//   				DestinationFileLocation: &InputFileLocationProperty{
//   					EfsFileLocation: &EfsInputFileLocationProperty{
//   						FileSystemId: jsii.String("fileSystemId"),
//   						Path: jsii.String("path"),
//   					},
//   					S3FileLocation: &S3InputFileLocationProperty{
//   						Bucket: jsii.String("bucket"),
//   						Key: jsii.String("key"),
//   					},
//   				},
//   				Name: jsii.String("name"),
//   				OverwriteExisting: jsii.String("overwriteExisting"),
//   				SourceFileLocation: jsii.String("sourceFileLocation"),
//   				Type: jsii.String("type"),
//   			},
//   			DeleteStepDetails: deleteStepDetails,
//   			TagStepDetails: tagStepDetails,
//   			Type: jsii.String("type"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-workflow.html
//
type CfnWorkflowPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnWorkflowMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkflowPropsMixin
type jsiiProxy_CfnWorkflowPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnWorkflowPropsMixin) Props() *CfnWorkflowMixinProps {
	var returns *CfnWorkflowMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflowPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Transfer::Workflow`.
func NewCfnWorkflowPropsMixin(props *CfnWorkflowMixinProps, options *mixins.CfnPropertyMixinOptions) CfnWorkflowPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkflowPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkflowPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnWorkflowPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Transfer::Workflow`.
func NewCfnWorkflowPropsMixin_Override(c CfnWorkflowPropsMixin, props *CfnWorkflowMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnWorkflowPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnWorkflowPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkflowPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnWorkflowPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkflowPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnWorkflowPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkflowPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnWorkflowPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

