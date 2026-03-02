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
type CfnDistributionConnectionLogsRecordFields string

const (
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_CONNECTIONSTATUS CfnDistributionConnectionLogsRecordFields = "CONNECTIONSTATUS"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_CLIENTIP CfnDistributionConnectionLogsRecordFields = "CLIENTIP"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_CLIENTPORT CfnDistributionConnectionLogsRecordFields = "CLIENTPORT"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_SERVERIP CfnDistributionConnectionLogsRecordFields = "SERVERIP"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_DISTRIBUTIONTENANTID CfnDistributionConnectionLogsRecordFields = "DISTRIBUTIONTENANTID"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_TLSPROTOCOL CfnDistributionConnectionLogsRecordFields = "TLSPROTOCOL"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_TLSCIPHER CfnDistributionConnectionLogsRecordFields = "TLSCIPHER"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_TLSHANDSHAKEDURATION CfnDistributionConnectionLogsRecordFields = "TLSHANDSHAKEDURATION"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_TLSSNI CfnDistributionConnectionLogsRecordFields = "TLSSNI"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_CLIENTLEAFCERTSERIALNUMBER CfnDistributionConnectionLogsRecordFields = "CLIENTLEAFCERTSERIALNUMBER"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_CLIENTLEAFCERTSUBJECT CfnDistributionConnectionLogsRecordFields = "CLIENTLEAFCERTSUBJECT"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_CLIENTLEAFCERTISSUER CfnDistributionConnectionLogsRecordFields = "CLIENTLEAFCERTISSUER"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_CLIENTLEAFCERTVALIDITY CfnDistributionConnectionLogsRecordFields = "CLIENTLEAFCERTVALIDITY"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_CONNECTIONLOGCUSTOMDATA CfnDistributionConnectionLogsRecordFields = "CONNECTIONLOGCUSTOMDATA"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_EVENTTIMESTAMP CfnDistributionConnectionLogsRecordFields = "EVENTTIMESTAMP"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_CONNECTIONID CfnDistributionConnectionLogsRecordFields = "CONNECTIONID"
	// Experimental.
	CfnDistributionConnectionLogsRecordFields_DISTRIBUTIONID CfnDistributionConnectionLogsRecordFields = "DISTRIBUTIONID"
)

