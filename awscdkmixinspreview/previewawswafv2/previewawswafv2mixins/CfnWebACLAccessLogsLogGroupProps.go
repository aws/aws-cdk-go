package previewawswafv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWebACLAccessLogsLogGroupProps := &CfnWebACLAccessLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnWebACLAccessLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnWebACLAccessLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnWebACLAccessLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

