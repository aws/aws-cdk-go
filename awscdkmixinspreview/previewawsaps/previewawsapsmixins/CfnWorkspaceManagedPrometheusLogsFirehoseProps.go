package previewawsapsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkspaceManagedPrometheusLogsFirehoseProps := &CfnWorkspaceManagedPrometheusLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnWorkspaceManagedPrometheusLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnWorkspaceManagedPrometheusLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnWorkspaceManagedPrometheusLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

