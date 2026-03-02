package previewawsnetworkfirewallmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFirewallTlsLogsFirehoseProps := &CfnFirewallTlsLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnFirewallTlsLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnFirewallTlsLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnFirewallTlsLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

