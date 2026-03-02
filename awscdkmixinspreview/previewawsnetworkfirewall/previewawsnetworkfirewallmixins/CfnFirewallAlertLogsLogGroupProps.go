package previewawsnetworkfirewallmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallAlertLogsLogGroupProps := &CfnFirewallAlertLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnFirewallAlertLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnFirewallAlertLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnFirewallAlertLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

