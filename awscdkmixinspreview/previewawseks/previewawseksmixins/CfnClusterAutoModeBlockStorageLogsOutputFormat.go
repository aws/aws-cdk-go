package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnClusterAutoModeBlockStorageLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeBlockStorageLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnClusterAutoModeBlockStorageLogsOutputFormat()
//
// Experimental.
type CfnClusterAutoModeBlockStorageLogsOutputFormat interface {
}

// The jsii proxy struct for CfnClusterAutoModeBlockStorageLogsOutputFormat
type jsiiProxy_CfnClusterAutoModeBlockStorageLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnClusterAutoModeBlockStorageLogsOutputFormat() CfnClusterAutoModeBlockStorageLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterAutoModeBlockStorageLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeBlockStorageLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnClusterAutoModeBlockStorageLogsOutputFormat_Override(c CfnClusterAutoModeBlockStorageLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeBlockStorageLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

