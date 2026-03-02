package previewawskafkaconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnConnectorApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConnectorApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnConnectorApplicationLogsOutputFormat()
//
// Experimental.
type CfnConnectorApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnConnectorApplicationLogsOutputFormat
type jsiiProxy_CfnConnectorApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnConnectorApplicationLogsOutputFormat() CfnConnectorApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnConnectorApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnConnectorApplicationLogsOutputFormat_Override(c CfnConnectorApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

