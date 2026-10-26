package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnCapabilityEksCapabilityArgocdNotificationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat()
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat
type jsiiProxy_CfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat() CfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat_Override(c CfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

