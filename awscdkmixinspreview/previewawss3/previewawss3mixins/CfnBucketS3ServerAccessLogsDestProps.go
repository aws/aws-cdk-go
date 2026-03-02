package previewawss3mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBucketS3ServerAccessLogsDestProps := &CfnBucketS3ServerAccessLogsDestProps{
//   	RecordFields: []CfnBucketS3ServerAccessLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnBucketS3ServerAccessLogsRecordFields_BUCKET_NAME,
//   	},
//   }
//
// Experimental.
type CfnBucketS3ServerAccessLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnBucketS3ServerAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

