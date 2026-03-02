package previewawsentityresolutionmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnIdMappingWorkflowWorkflowLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdMappingWorkflowWorkflowLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnIdMappingWorkflowWorkflowLogsOutputFormat()
//
// Experimental.
type CfnIdMappingWorkflowWorkflowLogsOutputFormat interface {
}

// The jsii proxy struct for CfnIdMappingWorkflowWorkflowLogsOutputFormat
type jsiiProxy_CfnIdMappingWorkflowWorkflowLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnIdMappingWorkflowWorkflowLogsOutputFormat() CfnIdMappingWorkflowWorkflowLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnIdMappingWorkflowWorkflowLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowWorkflowLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnIdMappingWorkflowWorkflowLogsOutputFormat_Override(c CfnIdMappingWorkflowWorkflowLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowWorkflowLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

