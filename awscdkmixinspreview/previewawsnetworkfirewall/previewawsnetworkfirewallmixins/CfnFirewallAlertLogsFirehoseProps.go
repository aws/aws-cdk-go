package previewawsnetworkfirewallmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallAlertLogsFirehoseProps := &CfnFirewallAlertLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnFirewallAlertLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnFirewallAlertLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnFirewallAlertLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnFirewallAlertLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnFirewallAlertLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnFirewallAlertLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

