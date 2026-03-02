package previewawsconnectmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceAmazonConnectFlowLogsDestProps := &CfnInstanceAmazonConnectFlowLogsDestProps{
//   	RecordFields: []CfnInstanceAmazonConnectFlowLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnInstanceAmazonConnectFlowLogsRecordFields_INSTANCEARN,
//   	},
//   }
//
// Experimental.
type CfnInstanceAmazonConnectFlowLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnInstanceAmazonConnectFlowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

