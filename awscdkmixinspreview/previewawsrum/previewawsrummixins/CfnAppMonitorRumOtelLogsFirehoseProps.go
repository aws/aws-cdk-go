package previewawsrummixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppMonitorRumOtelLogsFirehoseProps := &CfnAppMonitorRumOtelLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnAppMonitorRumOtelLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnAppMonitorRumOtelLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnAppMonitorRumOtelLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

