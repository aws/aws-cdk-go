package previewawsbackupgatewaymixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHypervisorBgwHypervisorLogsFirehoseProps := &CfnHypervisorBgwHypervisorLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnHypervisorBgwHypervisorLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnHypervisorBgwHypervisorLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnHypervisorBgwHypervisorLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

