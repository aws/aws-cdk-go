package previewawsvpclatticemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceAccessLogsLogGroupProps := &CfnServiceAccessLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnServiceAccessLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnServiceAccessLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnServiceAccessLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

