package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnMemoryTraces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMemoryTracesOutputFormat := awscdkmixinspreview.Mixins.NewCfnMemoryTracesOutputFormat()
//
// Experimental.
type CfnMemoryTracesOutputFormat interface {
}

// The jsii proxy struct for CfnMemoryTracesOutputFormat
type jsiiProxy_CfnMemoryTracesOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnMemoryTracesOutputFormat() CfnMemoryTracesOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnMemoryTracesOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryTracesOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMemoryTracesOutputFormat_Override(c CfnMemoryTracesOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryTracesOutputFormat",
		nil, // no parameters
		c,
	)
}

