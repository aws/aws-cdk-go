package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnBrowserCustomTraces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserCustomTracesOutputFormat := awscdkmixinspreview.Mixins.NewCfnBrowserCustomTracesOutputFormat()
//
// Experimental.
type CfnBrowserCustomTracesOutputFormat interface {
}

// The jsii proxy struct for CfnBrowserCustomTracesOutputFormat
type jsiiProxy_CfnBrowserCustomTracesOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnBrowserCustomTracesOutputFormat() CfnBrowserCustomTracesOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnBrowserCustomTracesOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomTracesOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnBrowserCustomTracesOutputFormat_Override(c CfnBrowserCustomTracesOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomTracesOutputFormat",
		nil, // no parameters
		c,
	)
}

