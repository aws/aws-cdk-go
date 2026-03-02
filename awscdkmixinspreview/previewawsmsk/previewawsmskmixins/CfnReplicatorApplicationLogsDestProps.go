package previewawsmskmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnReplicatorApplicationLogsDestProps := &CfnReplicatorApplicationLogsDestProps{
//   	RecordFields: []CfnReplicatorApplicationLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnReplicatorApplicationLogsRecordFields_REPLICATOR_NAME,
//   	},
//   }
//
// Experimental.
type CfnReplicatorApplicationLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnReplicatorApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

