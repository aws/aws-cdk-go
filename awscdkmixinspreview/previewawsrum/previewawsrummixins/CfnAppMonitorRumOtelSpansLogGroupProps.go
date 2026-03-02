package previewawsrummixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAppMonitorRumOtelSpansLogGroupProps := &CfnAppMonitorRumOtelSpansLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnAppMonitorRumOtelSpansOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnAppMonitorRumOtelSpansLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnAppMonitorRumOtelSpansOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

