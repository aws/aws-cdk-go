package previewawsmskmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnReplicatorApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnReplicatorApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnReplicatorApplicationLogsOutputFormat()
//
// Experimental.
type CfnReplicatorApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnReplicatorApplicationLogsOutputFormat
type jsiiProxy_CfnReplicatorApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnReplicatorApplicationLogsOutputFormat() CfnReplicatorApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnReplicatorApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnReplicatorApplicationLogsOutputFormat_Override(c CfnReplicatorApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnReplicatorApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

