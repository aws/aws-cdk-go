package previewawscognitomixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnUserPoolApplicationLogsLogGroupProps := &CfnUserPoolApplicationLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnUserPoolApplicationLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnUserPoolApplicationLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnUserPoolApplicationLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

