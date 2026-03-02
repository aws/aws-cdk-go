package previewawsnetworkfirewallmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallAlertLogsFirehoseProps := &CfnFirewallAlertLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnFirewallAlertLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnFirewallAlertLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnFirewallAlertLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

