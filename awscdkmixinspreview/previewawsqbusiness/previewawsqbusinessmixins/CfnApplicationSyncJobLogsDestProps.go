package previewawsqbusinessmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationSyncJobLogsDestProps := &CfnApplicationSyncJobLogsDestProps{
//   	RecordFields: []CfnApplicationSyncJobLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnApplicationSyncJobLogsRecordFields_AWSACCOUNTID,
//   	},
//   }
//
// Experimental.
type CfnApplicationSyncJobLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnApplicationSyncJobLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

