package previewawssesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnMailManagerRuleSetApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerRuleSetApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnMailManagerRuleSetApplicationLogsOutputFormat()
//
// Experimental.
type CfnMailManagerRuleSetApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnMailManagerRuleSetApplicationLogsOutputFormat
type jsiiProxy_CfnMailManagerRuleSetApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnMailManagerRuleSetApplicationLogsOutputFormat() CfnMailManagerRuleSetApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnMailManagerRuleSetApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMailManagerRuleSetApplicationLogsOutputFormat_Override(c CfnMailManagerRuleSetApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

