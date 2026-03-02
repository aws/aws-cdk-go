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
type CfnDistributionConnectionLogsOutputFormat_LogGroup string

const (
	// Experimental.
	CfnDistributionConnectionLogsOutputFormat_LogGroup_PLAIN CfnDistributionConnectionLogsOutputFormat_LogGroup = "PLAIN"
	// Experimental.
	CfnDistributionConnectionLogsOutputFormat_LogGroup_JSON CfnDistributionConnectionLogsOutputFormat_LogGroup = "JSON"
)

