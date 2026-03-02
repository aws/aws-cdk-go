package previewawspipesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPipeExecutionLogsLogGroupProps := &CfnPipeExecutionLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnPipeExecutionLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnPipeExecutionLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnPipeExecutionLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

