package previewawspcsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsSchedulerLogsDestProps := &CfnClusterPcsSchedulerLogsDestProps{
//   	RecordFields: []CfnClusterPcsSchedulerLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnClusterPcsSchedulerLogsRecordFields_RESOURCE_ID,
//   	},
//   }
//
// Experimental.
type CfnClusterPcsSchedulerLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterPcsSchedulerLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

