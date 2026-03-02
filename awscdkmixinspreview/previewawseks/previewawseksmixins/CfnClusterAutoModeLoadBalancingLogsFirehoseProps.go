package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeLoadBalancingLogsFirehoseProps := &CfnClusterAutoModeLoadBalancingLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterAutoModeLoadBalancingLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnClusterAutoModeLoadBalancingLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterAutoModeLoadBalancingLogsRecordFields_LEVEL,
//   	},
//   }
//
// Experimental.
type CfnClusterAutoModeLoadBalancingLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnClusterAutoModeLoadBalancingLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterAutoModeLoadBalancingLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

