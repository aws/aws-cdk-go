package previewawsconnectmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceAmazonConnectFlowLogsFirehoseProps := &CfnInstanceAmazonConnectFlowLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnInstanceAmazonConnectFlowLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnInstanceAmazonConnectFlowLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnInstanceAmazonConnectFlowLogsRecordFields_INSTANCEARN,
//   	},
//   }
//
// Experimental.
type CfnInstanceAmazonConnectFlowLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnInstanceAmazonConnectFlowLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnInstanceAmazonConnectFlowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

