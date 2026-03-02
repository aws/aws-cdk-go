package previewawskafkaconnectmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConnectorApplicationLogsFirehoseProps := &CfnConnectorApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnConnectorApplicationLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnConnectorApplicationLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnConnectorApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

