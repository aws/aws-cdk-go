package previewawscloudfrontmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnDistributionConnectionLogs.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   // Create CloudFront distribution
//   var origin IBucket
//
//   distribution := cloudfront.NewDistribution(scope, jsii.String("Distribution"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.S3BucketOrigin_WithOriginAccessControl(origin),
//   	},
//   })
//
//   // Create log destination
//   logGroup := logs.NewLogGroup(scope, jsii.String("DeliveryLogGroup"))
//
//   // Configure log delivery using the mixin
//   distribution.With(cloudfrontMixins.CfnDistributionLogsMixin_CONNECTION_LOGS().ToLogGroup(logGroup, &CfnDistributionConnectionLogsLogGroupProps{
//   	OutputFormat: cloudfrontMixins.CfnDistributionConnectionLogsOutputFormat.LogGroup_JSON,
//   	RecordFields: []CfnDistributionConnectionLogsRecordFields{
//   		cloudfrontMixins.CfnDistributionConnectionLogsRecordFields_CONNECTIONSTATUS,
//   		cloudfrontMixins.CfnDistributionConnectionLogsRecordFields_CLIENTIP,
//   		cloudfrontMixins.CfnDistributionConnectionLogsRecordFields_SERVERIP,
//   		cloudfrontMixins.CfnDistributionConnectionLogsRecordFields_TLSPROTOCOL,
//   	},
//   }))
//
// Experimental.
type CfnDistributionConnectionLogsOutputFormat interface {
}

// The jsii proxy struct for CfnDistributionConnectionLogsOutputFormat
type jsiiProxy_CfnDistributionConnectionLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnDistributionConnectionLogsOutputFormat() CfnDistributionConnectionLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnDistributionConnectionLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionConnectionLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnDistributionConnectionLogsOutputFormat_Override(c CfnDistributionConnectionLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionConnectionLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

