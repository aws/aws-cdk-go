package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnKnowledgeBaseTraces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBaseTracesOutputFormat := awscdkmixinspreview.Mixins.NewCfnKnowledgeBaseTracesOutputFormat()
//
// Experimental.
type CfnKnowledgeBaseTracesOutputFormat interface {
}

// The jsii proxy struct for CfnKnowledgeBaseTracesOutputFormat
type jsiiProxy_CfnKnowledgeBaseTracesOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnKnowledgeBaseTracesOutputFormat() CfnKnowledgeBaseTracesOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnKnowledgeBaseTracesOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseTracesOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnKnowledgeBaseTracesOutputFormat_Override(c CfnKnowledgeBaseTracesOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseTracesOutputFormat",
		nil, // no parameters
		c,
	)
}

