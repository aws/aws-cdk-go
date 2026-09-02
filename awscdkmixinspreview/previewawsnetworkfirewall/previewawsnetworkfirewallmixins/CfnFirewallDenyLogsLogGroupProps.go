package previewawsnetworkfirewallmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallDenyLogsLogGroupProps := &CfnFirewallDenyLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnFirewallDenyLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnFirewallDenyLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnFirewallDenyLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnFirewallDenyLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnFirewallDenyLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnFirewallDenyLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

