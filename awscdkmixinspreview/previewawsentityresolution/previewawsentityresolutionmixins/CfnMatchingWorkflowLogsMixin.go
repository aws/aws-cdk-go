package previewawsentityresolutionmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsentityresolution/previewawsentityresolutionmixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a matching workflow that defines the configuration for a data processing job.
//
// The workflow name must be unique. To modify an existing workflow, use `UpdateMatchingWorkflow` .
//
// > For workflows where `resolutionType` is `ML_MATCHING` or `PROVIDER` , incremental processing is not supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnMatchingWorkflowLogsMixin := awscdkmixinspreview.Mixins.NewCfnMatchingWorkflowLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html
//
type CfnMatchingWorkflowLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMatchingWorkflowLogsMixin
type jsiiProxy_CfnMatchingWorkflowLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMatchingWorkflowLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchingWorkflowLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::EntityResolution::MatchingWorkflow`.
func NewCfnMatchingWorkflowLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnMatchingWorkflowLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnMatchingWorkflowLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMatchingWorkflowLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::EntityResolution::MatchingWorkflow`.
func NewCfnMatchingWorkflowLogsMixin_Override(c CfnMatchingWorkflowLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMatchingWorkflowLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMatchingWorkflowLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMatchingWorkflowLogsMixin_WORKFLOW_LOGS() CfnMatchingWorkflowWorkflowLogs {
	_init_.Initialize()
	var returns CfnMatchingWorkflowWorkflowLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowLogsMixin",
		"WORKFLOW_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMatchingWorkflowLogsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMatchingWorkflowLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

