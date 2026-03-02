package previewawstransfermixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServerTransferLogsFirehoseProps := &CfnServerTransferLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnServerTransferLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnServerTransferLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnServerTransferLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

