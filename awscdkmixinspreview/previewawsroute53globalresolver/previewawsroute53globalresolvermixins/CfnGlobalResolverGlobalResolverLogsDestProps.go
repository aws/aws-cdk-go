package previewawsroute53globalresolvermixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGlobalResolverGlobalResolverLogsDestProps := &CfnGlobalResolverGlobalResolverLogsDestProps{
//   	RecordFields: []CfnGlobalResolverGlobalResolverLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnGlobalResolverGlobalResolverLogsRecordFields_ACTION_NAME,
//   	},
//   }
//
// Experimental.
type CfnGlobalResolverGlobalResolverLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnGlobalResolverGlobalResolverLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

