package previewawsmediatailormixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnPlaybackConfigurationManifestServiceLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackConfigurationManifestServiceLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnPlaybackConfigurationManifestServiceLogsOutputFormat()
//
// Experimental.
type CfnPlaybackConfigurationManifestServiceLogsOutputFormat interface {
}

// The jsii proxy struct for CfnPlaybackConfigurationManifestServiceLogsOutputFormat
type jsiiProxy_CfnPlaybackConfigurationManifestServiceLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnPlaybackConfigurationManifestServiceLogsOutputFormat() CfnPlaybackConfigurationManifestServiceLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnPlaybackConfigurationManifestServiceLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnPlaybackConfigurationManifestServiceLogsOutputFormat_Override(c CfnPlaybackConfigurationManifestServiceLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationManifestServiceLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

