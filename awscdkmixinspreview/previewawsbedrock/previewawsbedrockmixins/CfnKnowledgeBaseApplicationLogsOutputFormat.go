package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnKnowledgeBaseApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBaseApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnKnowledgeBaseApplicationLogsOutputFormat()
//
// Experimental.
type CfnKnowledgeBaseApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnKnowledgeBaseApplicationLogsOutputFormat
type jsiiProxy_CfnKnowledgeBaseApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnKnowledgeBaseApplicationLogsOutputFormat() CfnKnowledgeBaseApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnKnowledgeBaseApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnKnowledgeBaseApplicationLogsOutputFormat_Override(c CfnKnowledgeBaseApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

