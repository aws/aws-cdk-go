package previewawsmediatailormixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnPlaybackConfigurationTranscodeLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackConfigurationTranscodeLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnPlaybackConfigurationTranscodeLogsOutputFormat()
//
// Experimental.
type CfnPlaybackConfigurationTranscodeLogsOutputFormat interface {
}

// The jsii proxy struct for CfnPlaybackConfigurationTranscodeLogsOutputFormat
type jsiiProxy_CfnPlaybackConfigurationTranscodeLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnPlaybackConfigurationTranscodeLogsOutputFormat() CfnPlaybackConfigurationTranscodeLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnPlaybackConfigurationTranscodeLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationTranscodeLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnPlaybackConfigurationTranscodeLogsOutputFormat_Override(c CfnPlaybackConfigurationTranscodeLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationTranscodeLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

