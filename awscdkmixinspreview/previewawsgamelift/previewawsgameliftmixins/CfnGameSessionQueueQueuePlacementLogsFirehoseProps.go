package previewawsgameliftmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGameSessionQueueQueuePlacementLogsFirehoseProps := &CfnGameSessionQueueQueuePlacementLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnGameSessionQueueQueuePlacementLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnGameSessionQueueQueuePlacementLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnGameSessionQueueQueuePlacementLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnGameSessionQueueQueuePlacementLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnGameSessionQueueQueuePlacementLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnGameSessionQueueQueuePlacementLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

