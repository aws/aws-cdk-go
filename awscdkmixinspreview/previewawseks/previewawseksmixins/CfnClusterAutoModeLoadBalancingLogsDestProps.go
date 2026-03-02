package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeLoadBalancingLogsDestProps := &CfnClusterAutoModeLoadBalancingLogsDestProps{
//   	RecordFields: []CfnClusterAutoModeLoadBalancingLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnClusterAutoModeLoadBalancingLogsRecordFields_LEVEL,
//   	},
//   }
//
// Experimental.
type CfnClusterAutoModeLoadBalancingLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterAutoModeLoadBalancingLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

