package previewawssecurityhubmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHubSecurityFindingLogsLogGroupProps := &CfnHubSecurityFindingLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnHubSecurityFindingLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []eVENT{
//   		awscdkmixinspreview.*Mixins.CfnHubSecurityFindingLogsRecordFields_*eVENT,
//   	},
//   }
//
// Experimental.
type CfnHubSecurityFindingLogsLogGroupProps struct {
	// Format for log output, options are plain.
	// Experimental.
	OutputFormat CfnHubSecurityFindingLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnHubSecurityFindingLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

