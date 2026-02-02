package previewawsentityresolutionmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsentityresolution/previewawsentityresolutionmixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an `IdMappingWorkflow` object which stores the configuration of the data processing job to be run.
//
// Each `IdMappingWorkflow` must have a unique workflow name. To modify an existing workflow, use the UpdateIdMappingWorkflow API.
//
// > Incremental processing is not supported for ID mapping workflows.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnIdMappingWorkflowLogsMixin := awscdkmixinspreview.Mixins.NewCfnIdMappingWorkflowLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html
//
type CfnIdMappingWorkflowLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIdMappingWorkflowLogsMixin
type jsiiProxy_CfnIdMappingWorkflowLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIdMappingWorkflowLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdMappingWorkflowLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::EntityResolution::IdMappingWorkflow`.
func NewCfnIdMappingWorkflowLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnIdMappingWorkflowLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnIdMappingWorkflowLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIdMappingWorkflowLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::EntityResolution::IdMappingWorkflow`.
func NewCfnIdMappingWorkflowLogsMixin_Override(c CfnIdMappingWorkflowLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIdMappingWorkflowLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIdMappingWorkflowLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIdMappingWorkflowLogsMixin_WORKFLOW_LOGS() CfnIdMappingWorkflowWorkflowLogs {
	_init_.Initialize()
	var returns CfnIdMappingWorkflowWorkflowLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowLogsMixin",
		"WORKFLOW_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIdMappingWorkflowLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnIdMappingWorkflowLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

