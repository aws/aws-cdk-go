package previewawsosismixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPipelinePipelineLogsLogGroupProps := &CfnPipelinePipelineLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnPipelinePipelineLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnPipelinePipelineLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnPipelinePipelineLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

