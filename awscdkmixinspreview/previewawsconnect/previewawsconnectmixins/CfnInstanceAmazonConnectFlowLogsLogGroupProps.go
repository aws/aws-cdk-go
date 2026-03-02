package previewawsconnectmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceAmazonConnectFlowLogsLogGroupProps := &CfnInstanceAmazonConnectFlowLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnInstanceAmazonConnectFlowLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnInstanceAmazonConnectFlowLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnInstanceAmazonConnectFlowLogsRecordFields_INSTANCEARN,
//   	},
//   }
//
// Experimental.
type CfnInstanceAmazonConnectFlowLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnInstanceAmazonConnectFlowLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnInstanceAmazonConnectFlowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

