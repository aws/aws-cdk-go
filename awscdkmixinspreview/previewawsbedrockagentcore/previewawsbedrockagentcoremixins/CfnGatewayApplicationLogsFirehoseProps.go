package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGatewayApplicationLogsFirehoseProps := &CfnGatewayApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnGatewayApplicationLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnGatewayApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnGatewayApplicationLogsRecordFields_BODY,
//   	},
//   }
//
// Experimental.
type CfnGatewayApplicationLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnGatewayApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnGatewayApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

