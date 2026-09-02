package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnWorkloadIdentityTraces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkloadIdentityTracesOutputFormat := awscdkmixinspreview.Mixins.NewCfnWorkloadIdentityTracesOutputFormat()
//
// Experimental.
type CfnWorkloadIdentityTracesOutputFormat interface {
}

// The jsii proxy struct for CfnWorkloadIdentityTracesOutputFormat
type jsiiProxy_CfnWorkloadIdentityTracesOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnWorkloadIdentityTracesOutputFormat() CfnWorkloadIdentityTracesOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnWorkloadIdentityTracesOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityTracesOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnWorkloadIdentityTracesOutputFormat_Override(c CfnWorkloadIdentityTracesOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityTracesOutputFormat",
		nil, // no parameters
		c,
	)
}

