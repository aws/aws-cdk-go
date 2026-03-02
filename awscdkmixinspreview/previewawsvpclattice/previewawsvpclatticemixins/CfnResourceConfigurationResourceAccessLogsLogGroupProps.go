package previewawsvpclatticemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceConfigurationResourceAccessLogsLogGroupProps := &CfnResourceConfigurationResourceAccessLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnResourceConfigurationResourceAccessLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnResourceConfigurationResourceAccessLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnResourceConfigurationResourceAccessLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

