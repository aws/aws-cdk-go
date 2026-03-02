package previewawssagemakermixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkteamActivityLogsLogGroupProps := &CfnWorkteamActivityLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnWorkteamActivityLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnWorkteamActivityLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnWorkteamActivityLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

