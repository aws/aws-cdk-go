package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBaseRuntimeLogsLogGroupProps := &CfnKnowledgeBaseRuntimeLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnKnowledgeBaseRuntimeLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnKnowledgeBaseRuntimeLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnKnowledgeBaseRuntimeLogsRecordFields_EVENT_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnKnowledgeBaseRuntimeLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnKnowledgeBaseRuntimeLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnKnowledgeBaseRuntimeLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

