package previewawss3mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnBucketS3ServerAccessLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBucketS3ServerAccessLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnBucketS3ServerAccessLogsOutputFormat()
//
// Experimental.
type CfnBucketS3ServerAccessLogsOutputFormat interface {
}

// The jsii proxy struct for CfnBucketS3ServerAccessLogsOutputFormat
type jsiiProxy_CfnBucketS3ServerAccessLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnBucketS3ServerAccessLogsOutputFormat() CfnBucketS3ServerAccessLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnBucketS3ServerAccessLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketS3ServerAccessLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnBucketS3ServerAccessLogsOutputFormat_Override(c CfnBucketS3ServerAccessLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketS3ServerAccessLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

