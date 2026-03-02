package previewawswafv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWebACLTokenLogsFirehoseProps := &CfnWebACLTokenLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnWebACLTokenLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnWebACLTokenLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnWebACLTokenLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

