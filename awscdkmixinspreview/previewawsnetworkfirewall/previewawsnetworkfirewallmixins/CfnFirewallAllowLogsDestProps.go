package previewawsnetworkfirewallmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallAllowLogsDestProps := &CfnFirewallAllowLogsDestProps{
//   	RecordFields: []CfnFirewallAllowLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnFirewallAllowLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnFirewallAllowLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnFirewallAllowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

