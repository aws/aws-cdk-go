package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnCodeInterpreterCustomTraces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterCustomTracesOutputFormat := awscdkmixinspreview.Mixins.NewCfnCodeInterpreterCustomTracesOutputFormat()
//
// Experimental.
type CfnCodeInterpreterCustomTracesOutputFormat interface {
}

// The jsii proxy struct for CfnCodeInterpreterCustomTracesOutputFormat
type jsiiProxy_CfnCodeInterpreterCustomTracesOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnCodeInterpreterCustomTracesOutputFormat() CfnCodeInterpreterCustomTracesOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnCodeInterpreterCustomTracesOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomTracesOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCodeInterpreterCustomTracesOutputFormat_Override(c CfnCodeInterpreterCustomTracesOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomTracesOutputFormat",
		nil, // no parameters
		c,
	)
}

