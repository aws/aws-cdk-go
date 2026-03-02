package previewawspcsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnClusterPcsSchedulerLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsSchedulerLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnClusterPcsSchedulerLogsOutputFormat()
//
// Experimental.
type CfnClusterPcsSchedulerLogsOutputFormat interface {
}

// The jsii proxy struct for CfnClusterPcsSchedulerLogsOutputFormat
type jsiiProxy_CfnClusterPcsSchedulerLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnClusterPcsSchedulerLogsOutputFormat() CfnClusterPcsSchedulerLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterPcsSchedulerLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnClusterPcsSchedulerLogsOutputFormat_Override(c CfnClusterPcsSchedulerLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsSchedulerLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

