package previewawsmskmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterBrokerLogsFirehoseProps := &CfnClusterBrokerLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterBrokerLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnClusterBrokerLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnClusterBrokerLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

