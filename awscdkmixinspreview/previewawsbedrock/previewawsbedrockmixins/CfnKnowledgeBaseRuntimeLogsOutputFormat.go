package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnKnowledgeBaseRuntimeLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBaseRuntimeLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnKnowledgeBaseRuntimeLogsOutputFormat()
//
// Experimental.
type CfnKnowledgeBaseRuntimeLogsOutputFormat interface {
}

// The jsii proxy struct for CfnKnowledgeBaseRuntimeLogsOutputFormat
type jsiiProxy_CfnKnowledgeBaseRuntimeLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnKnowledgeBaseRuntimeLogsOutputFormat() CfnKnowledgeBaseRuntimeLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnKnowledgeBaseRuntimeLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseRuntimeLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnKnowledgeBaseRuntimeLogsOutputFormat_Override(c CfnKnowledgeBaseRuntimeLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseRuntimeLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

