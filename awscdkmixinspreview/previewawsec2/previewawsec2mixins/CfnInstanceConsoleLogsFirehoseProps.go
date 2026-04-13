package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceConsoleLogsFirehoseProps := &CfnInstanceConsoleLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnInstanceConsoleLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnInstanceConsoleLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnInstanceConsoleLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnInstanceConsoleLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnInstanceConsoleLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnInstanceConsoleLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

