package previewawsm2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnApplicationBatchJobLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationBatchJobLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnApplicationBatchJobLogsOutputFormat()
//
// Experimental.
type CfnApplicationBatchJobLogsOutputFormat interface {
}

// The jsii proxy struct for CfnApplicationBatchJobLogsOutputFormat
type jsiiProxy_CfnApplicationBatchJobLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnApplicationBatchJobLogsOutputFormat() CfnApplicationBatchJobLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationBatchJobLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationBatchJobLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnApplicationBatchJobLogsOutputFormat_Override(c CfnApplicationBatchJobLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationBatchJobLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

