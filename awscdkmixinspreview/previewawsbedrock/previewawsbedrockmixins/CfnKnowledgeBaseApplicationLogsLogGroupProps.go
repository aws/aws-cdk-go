package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBaseApplicationLogsLogGroupProps := &CfnKnowledgeBaseApplicationLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnKnowledgeBaseApplicationLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnKnowledgeBaseApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnKnowledgeBaseApplicationLogsRecordFields_EVENT_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnKnowledgeBaseApplicationLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnKnowledgeBaseApplicationLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnKnowledgeBaseApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

