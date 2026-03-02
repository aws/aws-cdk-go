package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBaseRuntimeLogsFirehoseProps := &CfnKnowledgeBaseRuntimeLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnKnowledgeBaseRuntimeLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnKnowledgeBaseRuntimeLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnKnowledgeBaseRuntimeLogsRecordFields_EVENT_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnKnowledgeBaseRuntimeLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnKnowledgeBaseRuntimeLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnKnowledgeBaseRuntimeLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

