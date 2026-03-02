package previewawsm2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationConfigLogsFirehoseProps := &CfnApplicationConfigLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnApplicationConfigLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnApplicationConfigLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnApplicationConfigLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

