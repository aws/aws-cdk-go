package previewawsvpclatticemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceAccessLogsFirehoseProps := &CfnServiceAccessLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnServiceAccessLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnServiceAccessLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnServiceAccessLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

