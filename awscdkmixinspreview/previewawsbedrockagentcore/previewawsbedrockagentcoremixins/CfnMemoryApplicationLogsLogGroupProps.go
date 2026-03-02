package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMemoryApplicationLogsLogGroupProps := &CfnMemoryApplicationLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnMemoryApplicationLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnMemoryApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnMemoryApplicationLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnMemoryApplicationLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnMemoryApplicationLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMemoryApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

