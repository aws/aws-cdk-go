package previewawsnetworkfirewallmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallFlowLogsFirehoseProps := &CfnFirewallFlowLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnFirewallFlowLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnFirewallFlowLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnFirewallFlowLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

