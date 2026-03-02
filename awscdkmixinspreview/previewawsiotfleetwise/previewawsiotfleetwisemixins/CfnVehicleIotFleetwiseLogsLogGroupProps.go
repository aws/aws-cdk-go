package previewawsiotfleetwisemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVehicleIotFleetwiseLogsLogGroupProps := &CfnVehicleIotFleetwiseLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnVehicleIotFleetwiseLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnVehicleIotFleetwiseLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnVehicleIotFleetwiseLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

