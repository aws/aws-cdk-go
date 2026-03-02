package previewawss3mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBucketS3ServerAccessLogsFirehoseProps := &CfnBucketS3ServerAccessLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnBucketS3ServerAccessLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnBucketS3ServerAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnBucketS3ServerAccessLogsRecordFields_BUCKET_NAME,
//   	},
//   }
//
// Experimental.
type CfnBucketS3ServerAccessLogsFirehoseProps struct {
	// Format for log output, options are json.
	// Experimental.
	OutputFormat CfnBucketS3ServerAccessLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnBucketS3ServerAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

