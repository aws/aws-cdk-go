package previewawsrummixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppMonitorRumTelemetryLogsLogGroupProps := &CfnAppMonitorRumTelemetryLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnAppMonitorRumTelemetryLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnAppMonitorRumTelemetryLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnAppMonitorRumTelemetryLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

