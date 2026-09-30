package previewawspcsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsScalingLogsDestProps := &CfnClusterPcsScalingLogsDestProps{
//   	RecordFields: []CfnClusterPcsScalingLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnClusterPcsScalingLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnClusterPcsScalingLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterPcsScalingLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

