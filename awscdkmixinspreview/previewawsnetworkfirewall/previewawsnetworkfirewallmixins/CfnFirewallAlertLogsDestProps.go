package previewawsnetworkfirewallmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallAlertLogsDestProps := &CfnFirewallAlertLogsDestProps{
//   	RecordFields: []CfnFirewallAlertLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnFirewallAlertLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnFirewallAlertLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnFirewallAlertLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

