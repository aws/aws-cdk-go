package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceConsoleLogsLogGroupProps := &CfnInstanceConsoleLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnInstanceConsoleLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnInstanceConsoleLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnInstanceConsoleLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnInstanceConsoleLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnInstanceConsoleLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnInstanceConsoleLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

