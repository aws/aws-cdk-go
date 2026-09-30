package previewawspcsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsScalingLogsFirehoseProps := &CfnClusterPcsScalingLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterPcsScalingLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnClusterPcsScalingLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterPcsScalingLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnClusterPcsScalingLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnClusterPcsScalingLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterPcsScalingLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

