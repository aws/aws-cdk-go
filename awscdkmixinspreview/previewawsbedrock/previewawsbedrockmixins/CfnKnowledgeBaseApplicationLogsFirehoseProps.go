package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBaseApplicationLogsFirehoseProps := &CfnKnowledgeBaseApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnKnowledgeBaseApplicationLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnKnowledgeBaseApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnKnowledgeBaseApplicationLogsRecordFields_EVENT_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnKnowledgeBaseApplicationLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnKnowledgeBaseApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnKnowledgeBaseApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

