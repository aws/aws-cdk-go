package previewawsrummixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppMonitorRumOtelLogsLogGroupProps := &CfnAppMonitorRumOtelLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnAppMonitorRumOtelLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnAppMonitorRumOtelLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnAppMonitorRumOtelLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

