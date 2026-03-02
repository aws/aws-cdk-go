package previewawsapsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnScraperApplicationLogsFirehoseProps := &CfnScraperApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnScraperApplicationLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnScraperApplicationLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnScraperApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

