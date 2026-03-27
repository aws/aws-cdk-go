package previewawsroute53globalresolvermixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGlobalResolverGlobalResolverLogsFirehoseProps := &CfnGlobalResolverGlobalResolverLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnGlobalResolverGlobalResolverLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnGlobalResolverGlobalResolverLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnGlobalResolverGlobalResolverLogsRecordFields_ACTION_NAME,
//   	},
//   }
//
// Experimental.
type CfnGlobalResolverGlobalResolverLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnGlobalResolverGlobalResolverLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnGlobalResolverGlobalResolverLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

