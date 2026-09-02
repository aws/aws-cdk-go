package previewawsnetworkfirewallmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallDenyLogsFirehoseProps := &CfnFirewallDenyLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnFirewallDenyLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnFirewallDenyLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnFirewallDenyLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnFirewallDenyLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnFirewallDenyLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnFirewallDenyLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

