package previewawsnetworkfirewallmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallAllowLogsLogGroupProps := &CfnFirewallAllowLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnFirewallAllowLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnFirewallAllowLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnFirewallAllowLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnFirewallAllowLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnFirewallAllowLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnFirewallAllowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

