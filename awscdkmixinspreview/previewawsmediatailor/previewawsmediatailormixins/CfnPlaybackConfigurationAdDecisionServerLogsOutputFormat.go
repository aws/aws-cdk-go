package previewawsmediatailormixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnPlaybackConfigurationAdDecisionServerLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackConfigurationAdDecisionServerLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnPlaybackConfigurationAdDecisionServerLogsOutputFormat()
//
// Experimental.
type CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat interface {
}

// The jsii proxy struct for CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat
type jsiiProxy_CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnPlaybackConfigurationAdDecisionServerLogsOutputFormat() CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_Override(c CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

