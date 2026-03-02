package previewawsivschatmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnRoomIvsChatLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRoomIvsChatLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnRoomIvsChatLogsOutputFormat()
//
// Experimental.
type CfnRoomIvsChatLogsOutputFormat interface {
}

// The jsii proxy struct for CfnRoomIvsChatLogsOutputFormat
type jsiiProxy_CfnRoomIvsChatLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnRoomIvsChatLogsOutputFormat() CfnRoomIvsChatLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnRoomIvsChatLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomIvsChatLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnRoomIvsChatLogsOutputFormat_Override(c CfnRoomIvsChatLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomIvsChatLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

