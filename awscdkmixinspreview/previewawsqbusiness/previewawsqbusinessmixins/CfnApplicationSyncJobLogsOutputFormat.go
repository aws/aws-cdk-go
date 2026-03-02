package previewawsqbusinessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnApplicationSyncJobLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationSyncJobLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnApplicationSyncJobLogsOutputFormat()
//
// Experimental.
type CfnApplicationSyncJobLogsOutputFormat interface {
}

// The jsii proxy struct for CfnApplicationSyncJobLogsOutputFormat
type jsiiProxy_CfnApplicationSyncJobLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnApplicationSyncJobLogsOutputFormat() CfnApplicationSyncJobLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationSyncJobLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationSyncJobLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnApplicationSyncJobLogsOutputFormat_Override(c CfnApplicationSyncJobLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_qbusiness.mixins.CfnApplicationSyncJobLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

