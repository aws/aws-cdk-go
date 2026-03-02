package previewawsentityresolutionmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnMatchingWorkflowWorkflowLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMatchingWorkflowWorkflowLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnMatchingWorkflowWorkflowLogsOutputFormat()
//
// Experimental.
type CfnMatchingWorkflowWorkflowLogsOutputFormat interface {
}

// The jsii proxy struct for CfnMatchingWorkflowWorkflowLogsOutputFormat
type jsiiProxy_CfnMatchingWorkflowWorkflowLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnMatchingWorkflowWorkflowLogsOutputFormat() CfnMatchingWorkflowWorkflowLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnMatchingWorkflowWorkflowLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowWorkflowLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMatchingWorkflowWorkflowLogsOutputFormat_Override(c CfnMatchingWorkflowWorkflowLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowWorkflowLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

