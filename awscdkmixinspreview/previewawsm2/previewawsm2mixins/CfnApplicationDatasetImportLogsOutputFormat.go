package previewawsm2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnApplicationDatasetImportLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationDatasetImportLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnApplicationDatasetImportLogsOutputFormat()
//
// Experimental.
type CfnApplicationDatasetImportLogsOutputFormat interface {
}

// The jsii proxy struct for CfnApplicationDatasetImportLogsOutputFormat
type jsiiProxy_CfnApplicationDatasetImportLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnApplicationDatasetImportLogsOutputFormat() CfnApplicationDatasetImportLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnApplicationDatasetImportLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationDatasetImportLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnApplicationDatasetImportLogsOutputFormat_Override(c CfnApplicationDatasetImportLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_m2.mixins.CfnApplicationDatasetImportLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

