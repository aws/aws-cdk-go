package previewawscloudfrontmixins


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
type CfnDistributionConnectionLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnDistributionConnectionLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnDistributionConnectionLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

