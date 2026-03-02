package previewawscloudfrontmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDistributionConnectionLogsDestProps := &CfnDistributionConnectionLogsDestProps{
//   	RecordFields: []CfnDistributionConnectionLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnDistributionConnectionLogsRecordFields_CONNECTIONSTATUS,
//   	},
//   }
//
// Experimental.
type CfnDistributionConnectionLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnDistributionConnectionLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

