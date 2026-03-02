package previewawsnetworkfirewallmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallFlowLogsLogGroupProps := &CfnFirewallFlowLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnFirewallFlowLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnFirewallFlowLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnFirewallFlowLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

