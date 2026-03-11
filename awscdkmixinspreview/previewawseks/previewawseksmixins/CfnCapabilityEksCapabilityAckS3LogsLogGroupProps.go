package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityAckS3LogsLogGroupProps := &CfnCapabilityEksCapabilityAckS3LogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityAckS3LogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnCapabilityEksCapabilityAckS3LogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCapabilityEksCapabilityAckS3LogsRecordFields_STREAM,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityAckS3LogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnCapabilityEksCapabilityAckS3LogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityAckS3LogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

