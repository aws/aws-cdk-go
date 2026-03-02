package previewawscloudfrontmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDistributionAccessLogsDestProps := &CfnDistributionAccessLogsDestProps{
//   	RecordFields: []CfnDistributionAccessLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnDistributionAccessLogsRecordFields_DATE,
//   	},
//   }
//
// Experimental.
type CfnDistributionAccessLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnDistributionAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

