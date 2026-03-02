package previewawsm2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationBatchJobLogsFirehoseProps := &CfnApplicationBatchJobLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnApplicationBatchJobLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnApplicationBatchJobLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnApplicationBatchJobLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

