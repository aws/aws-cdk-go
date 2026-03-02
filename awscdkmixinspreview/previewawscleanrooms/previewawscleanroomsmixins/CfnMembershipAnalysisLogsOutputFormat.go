package previewawscleanroomsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnMembershipAnalysisLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMembershipAnalysisLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnMembershipAnalysisLogsOutputFormat()
//
// Experimental.
type CfnMembershipAnalysisLogsOutputFormat interface {
}

// The jsii proxy struct for CfnMembershipAnalysisLogsOutputFormat
type jsiiProxy_CfnMembershipAnalysisLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnMembershipAnalysisLogsOutputFormat() CfnMembershipAnalysisLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnMembershipAnalysisLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipAnalysisLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMembershipAnalysisLogsOutputFormat_Override(c CfnMembershipAnalysisLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cleanrooms.mixins.CfnMembershipAnalysisLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

