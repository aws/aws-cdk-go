package previewawsm2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationDatasetImportLogsLogGroupProps := &CfnApplicationDatasetImportLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnApplicationDatasetImportLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnApplicationDatasetImportLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnApplicationDatasetImportLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

