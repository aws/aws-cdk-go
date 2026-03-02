package previewawsm2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationConsoleLogsFirehoseProps := &CfnApplicationConsoleLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnApplicationConsoleLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnApplicationConsoleLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnApplicationConsoleLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

