package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentAliasApplicationLogsDestProps := &CfnAgentAliasApplicationLogsDestProps{
//   	RecordFields: []CfnAgentAliasApplicationLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnAgentAliasApplicationLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnAgentAliasApplicationLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnAgentAliasApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

