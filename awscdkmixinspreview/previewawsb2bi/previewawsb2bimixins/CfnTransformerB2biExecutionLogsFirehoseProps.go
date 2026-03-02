package previewawsb2bimixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransformerB2biExecutionLogsFirehoseProps := &CfnTransformerB2biExecutionLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnTransformerB2biExecutionLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnTransformerB2biExecutionLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnTransformerB2biExecutionLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

