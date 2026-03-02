package previewawsshieldmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProtectionFlowLogsLogGroupProps := &CfnProtectionFlowLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnProtectionFlowLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnProtectionFlowLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnProtectionFlowLogsRecordFields_SRCADDR,
//   	},
//   }
//
// Experimental.
type CfnProtectionFlowLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnProtectionFlowLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnProtectionFlowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

