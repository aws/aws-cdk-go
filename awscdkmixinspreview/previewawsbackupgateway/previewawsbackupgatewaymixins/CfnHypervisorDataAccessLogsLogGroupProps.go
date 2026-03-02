package previewawsbackupgatewaymixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHypervisorDataAccessLogsLogGroupProps := &CfnHypervisorDataAccessLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnHypervisorDataAccessLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnHypervisorDataAccessLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnHypervisorDataAccessLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

