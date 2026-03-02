package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnRuntimeTraces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuntimeTracesOutputFormat := awscdkmixinspreview.Mixins.NewCfnRuntimeTracesOutputFormat()
//
// Experimental.
type CfnRuntimeTracesOutputFormat interface {
}

// The jsii proxy struct for CfnRuntimeTracesOutputFormat
type jsiiProxy_CfnRuntimeTracesOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnRuntimeTracesOutputFormat() CfnRuntimeTracesOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnRuntimeTracesOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeTracesOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnRuntimeTracesOutputFormat_Override(c CfnRuntimeTracesOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeTracesOutputFormat",
		nil, // no parameters
		c,
	)
}

