package previewawsm2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationBatchJobLogsLogGroupProps := &CfnApplicationBatchJobLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnApplicationBatchJobLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnApplicationBatchJobLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnApplicationBatchJobLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

