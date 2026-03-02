package previewawsmskmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterBrokerLogsLogGroupProps := &CfnClusterBrokerLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterBrokerLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnClusterBrokerLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnClusterBrokerLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

