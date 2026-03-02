package previewawsiotfleetwisemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVehicleIotFleetwiseLogsFirehoseProps := &CfnVehicleIotFleetwiseLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnVehicleIotFleetwiseLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnVehicleIotFleetwiseLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnVehicleIotFleetwiseLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

