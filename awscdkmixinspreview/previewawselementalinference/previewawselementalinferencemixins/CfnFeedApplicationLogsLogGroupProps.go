package previewawselementalinferencemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFeedApplicationLogsLogGroupProps := &CfnFeedApplicationLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnFeedApplicationLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnFeedApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnFeedApplicationLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnFeedApplicationLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnFeedApplicationLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnFeedApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

