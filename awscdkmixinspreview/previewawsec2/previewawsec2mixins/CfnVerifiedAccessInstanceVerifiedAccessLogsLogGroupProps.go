package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVerifiedAccessInstanceVerifiedAccessLogsLogGroupProps := &CfnVerifiedAccessInstanceVerifiedAccessLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnVerifiedAccessInstanceVerifiedAccessLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

