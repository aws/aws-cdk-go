package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGatewayApplicationLogsDestProps := &CfnGatewayApplicationLogsDestProps{
//   	RecordFields: []CfnGatewayApplicationLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnGatewayApplicationLogsRecordFields_BODY,
//   	},
//   }
//
// Experimental.
type CfnGatewayApplicationLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnGatewayApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

