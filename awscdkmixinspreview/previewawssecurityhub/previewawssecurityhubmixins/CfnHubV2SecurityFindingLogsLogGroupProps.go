package previewawssecurityhubmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHubV2SecurityFindingLogsLogGroupProps := &CfnHubV2SecurityFindingLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnHubV2SecurityFindingLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []eVENT{
//   		awscdkmixinspreview.*Mixins.CfnHubV2SecurityFindingLogsRecordFields_*eVENT,
//   	},
//   }
//
// Experimental.
type CfnHubV2SecurityFindingLogsLogGroupProps struct {
	// Format for log output, options are plain.
	// Experimental.
	OutputFormat CfnHubV2SecurityFindingLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnHubV2SecurityFindingLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

