package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeLoadBalancingLogsLogGroupProps := &CfnClusterAutoModeLoadBalancingLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterAutoModeLoadBalancingLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnClusterAutoModeLoadBalancingLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterAutoModeLoadBalancingLogsRecordFields_LEVEL,
//   	},
//   }
//
// Experimental.
type CfnClusterAutoModeLoadBalancingLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnClusterAutoModeLoadBalancingLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterAutoModeLoadBalancingLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

