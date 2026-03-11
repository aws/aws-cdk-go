package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityAckS3LogsDestProps := &CfnCapabilityEksCapabilityAckS3LogsDestProps{
//   	RecordFields: []CfnCapabilityEksCapabilityAckS3LogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityAckS3LogsRecordFields_STREAM,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityAckS3LogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityAckS3LogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

