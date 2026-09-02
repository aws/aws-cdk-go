package previewawssecurityhubmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnHubSecurityFindingLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHubSecurityFindingLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnHubSecurityFindingLogsOutputFormat()
//
// Experimental.
type CfnHubSecurityFindingLogsOutputFormat interface {
}

// The jsii proxy struct for CfnHubSecurityFindingLogsOutputFormat
type jsiiProxy_CfnHubSecurityFindingLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnHubSecurityFindingLogsOutputFormat() CfnHubSecurityFindingLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnHubSecurityFindingLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubSecurityFindingLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnHubSecurityFindingLogsOutputFormat_Override(c CfnHubSecurityFindingLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubSecurityFindingLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

