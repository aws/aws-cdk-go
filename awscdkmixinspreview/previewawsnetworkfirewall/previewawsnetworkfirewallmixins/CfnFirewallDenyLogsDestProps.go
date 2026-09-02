package previewawsnetworkfirewallmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallDenyLogsDestProps := &CfnFirewallDenyLogsDestProps{
//   	RecordFields: []CfnFirewallDenyLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnFirewallDenyLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnFirewallDenyLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnFirewallDenyLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

