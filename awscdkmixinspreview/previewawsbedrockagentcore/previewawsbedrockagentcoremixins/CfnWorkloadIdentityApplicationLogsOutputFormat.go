package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnWorkloadIdentityApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkloadIdentityApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnWorkloadIdentityApplicationLogsOutputFormat()
//
// Experimental.
type CfnWorkloadIdentityApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnWorkloadIdentityApplicationLogsOutputFormat
type jsiiProxy_CfnWorkloadIdentityApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnWorkloadIdentityApplicationLogsOutputFormat() CfnWorkloadIdentityApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnWorkloadIdentityApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnWorkloadIdentityApplicationLogsOutputFormat_Override(c CfnWorkloadIdentityApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

