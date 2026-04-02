package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnCapabilityEksCapabilityArgocdCommitserverLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat()
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat interface {
}

// The jsii proxy struct for CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat
type jsiiProxy_CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat() CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat_Override(c CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

