package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPNConnectionEventLogsDestProps := &CfnVPNConnectionEventLogsDestProps{
//   	RecordFields: []CfnVPNConnectionEventLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnVPNConnectionEventLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnVPNConnectionEventLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnVPNConnectionEventLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

