package previewawsroute53globalresolvermixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGlobalResolverGlobalResolverLogsLogGroupProps := &CfnGlobalResolverGlobalResolverLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnGlobalResolverGlobalResolverLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnGlobalResolverGlobalResolverLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnGlobalResolverGlobalResolverLogsRecordFields_ACTION_NAME,
//   	},
//   }
//
// Experimental.
type CfnGlobalResolverGlobalResolverLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnGlobalResolverGlobalResolverLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnGlobalResolverGlobalResolverLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

