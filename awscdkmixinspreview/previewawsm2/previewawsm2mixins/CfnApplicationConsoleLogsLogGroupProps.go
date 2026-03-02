package previewawsm2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationConsoleLogsLogGroupProps := &CfnApplicationConsoleLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnApplicationConsoleLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnApplicationConsoleLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnApplicationConsoleLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

