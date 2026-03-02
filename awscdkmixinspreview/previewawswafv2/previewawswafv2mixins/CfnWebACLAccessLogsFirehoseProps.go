package previewawswafv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWebACLAccessLogsFirehoseProps := &CfnWebACLAccessLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnWebACLAccessLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnWebACLAccessLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnWebACLAccessLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

