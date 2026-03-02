package previewawsb2bimixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransformerB2biExecutionLogsLogGroupProps := &CfnTransformerB2biExecutionLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnTransformerB2biExecutionLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnTransformerB2biExecutionLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnTransformerB2biExecutionLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

