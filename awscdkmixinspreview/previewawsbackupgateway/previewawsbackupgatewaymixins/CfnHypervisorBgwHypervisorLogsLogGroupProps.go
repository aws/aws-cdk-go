package previewawsbackupgatewaymixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHypervisorBgwHypervisorLogsLogGroupProps := &CfnHypervisorBgwHypervisorLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnHypervisorBgwHypervisorLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnHypervisorBgwHypervisorLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnHypervisorBgwHypervisorLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

