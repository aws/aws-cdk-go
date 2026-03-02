package previewawsnetworkfirewallmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallTlsLogsLogGroupProps := &CfnFirewallTlsLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnFirewallTlsLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnFirewallTlsLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnFirewallTlsLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

