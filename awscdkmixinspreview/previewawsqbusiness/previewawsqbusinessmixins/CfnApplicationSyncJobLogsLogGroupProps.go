package previewawsqbusinessmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationSyncJobLogsLogGroupProps := &CfnApplicationSyncJobLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnApplicationSyncJobLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnApplicationSyncJobLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnApplicationSyncJobLogsRecordFields_AWSACCOUNTID,
//   	},
//   }
//
// Experimental.
type CfnApplicationSyncJobLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnApplicationSyncJobLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnApplicationSyncJobLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

