package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceConsoleLogsDestProps := &CfnInstanceConsoleLogsDestProps{
//   	RecordFields: []CfnInstanceConsoleLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnInstanceConsoleLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnInstanceConsoleLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnInstanceConsoleLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

