package previewawsmskmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnClusterBrokerLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterBrokerLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnClusterBrokerLogsOutputFormat()
//
// Experimental.
type CfnClusterBrokerLogsOutputFormat interface {
}

// The jsii proxy struct for CfnClusterBrokerLogsOutputFormat
type jsiiProxy_CfnClusterBrokerLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnClusterBrokerLogsOutputFormat() CfnClusterBrokerLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterBrokerLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterBrokerLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnClusterBrokerLogsOutputFormat_Override(c CfnClusterBrokerLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_msk.mixins.CfnClusterBrokerLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

