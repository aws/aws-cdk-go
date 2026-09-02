package previewawssecurityhubmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnHubV2SecurityFindingLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHubV2SecurityFindingLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnHubV2SecurityFindingLogsOutputFormat()
//
// Experimental.
type CfnHubV2SecurityFindingLogsOutputFormat interface {
}

// The jsii proxy struct for CfnHubV2SecurityFindingLogsOutputFormat
type jsiiProxy_CfnHubV2SecurityFindingLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnHubV2SecurityFindingLogsOutputFormat() CfnHubV2SecurityFindingLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnHubV2SecurityFindingLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubV2SecurityFindingLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnHubV2SecurityFindingLogsOutputFormat_Override(c CfnHubV2SecurityFindingLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubV2SecurityFindingLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

