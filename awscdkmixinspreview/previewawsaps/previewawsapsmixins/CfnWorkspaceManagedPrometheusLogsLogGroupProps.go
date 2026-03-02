package previewawsapsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkspaceManagedPrometheusLogsLogGroupProps := &CfnWorkspaceManagedPrometheusLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnWorkspaceManagedPrometheusLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnWorkspaceManagedPrometheusLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnWorkspaceManagedPrometheusLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

