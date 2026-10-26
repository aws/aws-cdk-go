package previewawsgameliftmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnGameSessionQueueQueuePlacementLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGameSessionQueueQueuePlacementLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnGameSessionQueueQueuePlacementLogsOutputFormat()
//
// Experimental.
type CfnGameSessionQueueQueuePlacementLogsOutputFormat interface {
}

// The jsii proxy struct for CfnGameSessionQueueQueuePlacementLogsOutputFormat
type jsiiProxy_CfnGameSessionQueueQueuePlacementLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnGameSessionQueueQueuePlacementLogsOutputFormat() CfnGameSessionQueueQueuePlacementLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnGameSessionQueueQueuePlacementLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueQueuePlacementLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnGameSessionQueueQueuePlacementLogsOutputFormat_Override(c CfnGameSessionQueueQueuePlacementLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_gamelift.mixins.CfnGameSessionQueueQueuePlacementLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

