package previewawswafv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWebACLTokenLogsLogGroupProps := &CfnWebACLTokenLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnWebACLTokenLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnWebACLTokenLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnWebACLTokenLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

