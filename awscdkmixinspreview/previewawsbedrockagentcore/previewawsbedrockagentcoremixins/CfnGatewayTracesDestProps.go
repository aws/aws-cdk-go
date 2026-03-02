package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGatewayTracesDestProps := &CfnGatewayTracesDestProps{
//   	RecordFields: []tRACE{
//   		awscdkmixinspreview.Mixins.CfnGatewayTracesRecordFields_*tRACE,
//   	},
//   }
//
// Experimental.
type CfnGatewayTracesDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnGatewayTracesRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

