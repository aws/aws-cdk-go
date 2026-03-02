package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPNConnectionConnectionLogsDestProps := &CfnVPNConnectionConnectionLogsDestProps{
//   	RecordFields: []CfnVPNConnectionConnectionLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnVPNConnectionConnectionLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnVPNConnectionConnectionLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnVPNConnectionConnectionLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

