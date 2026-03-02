package previewawss3mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBucketS3ServerAccessLogsLogGroupProps := &CfnBucketS3ServerAccessLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnBucketS3ServerAccessLogsOutputFormat.LogGroup_JSON,
//   	RecordFields: []CfnBucketS3ServerAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnBucketS3ServerAccessLogsRecordFields_BUCKET_NAME,
//   	},
//   }
//
// Experimental.
type CfnBucketS3ServerAccessLogsLogGroupProps struct {
	// Format for log output, options are json.
	// Experimental.
	OutputFormat CfnBucketS3ServerAccessLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnBucketS3ServerAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

