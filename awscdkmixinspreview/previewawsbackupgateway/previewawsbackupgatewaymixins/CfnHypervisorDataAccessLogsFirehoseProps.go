package previewawsbackupgatewaymixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHypervisorDataAccessLogsFirehoseProps := &CfnHypervisorDataAccessLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnHypervisorDataAccessLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnHypervisorDataAccessLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnHypervisorDataAccessLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

