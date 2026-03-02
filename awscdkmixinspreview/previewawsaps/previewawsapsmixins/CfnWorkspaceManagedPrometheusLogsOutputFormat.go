package previewawsapsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnWorkspaceManagedPrometheusLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkspaceManagedPrometheusLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnWorkspaceManagedPrometheusLogsOutputFormat()
//
// Experimental.
type CfnWorkspaceManagedPrometheusLogsOutputFormat interface {
}

// The jsii proxy struct for CfnWorkspaceManagedPrometheusLogsOutputFormat
type jsiiProxy_CfnWorkspaceManagedPrometheusLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnWorkspaceManagedPrometheusLogsOutputFormat() CfnWorkspaceManagedPrometheusLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnWorkspaceManagedPrometheusLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspaceManagedPrometheusLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnWorkspaceManagedPrometheusLogsOutputFormat_Override(c CfnWorkspaceManagedPrometheusLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnWorkspaceManagedPrometheusLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

