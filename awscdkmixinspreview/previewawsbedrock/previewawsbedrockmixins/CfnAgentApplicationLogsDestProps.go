package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentApplicationLogsDestProps := &CfnAgentApplicationLogsDestProps{
//   	RecordFields: []CfnAgentApplicationLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnAgentApplicationLogsRecordFields_AGENT_ALIAS_ARN,
//   	},
//   }
//
// Experimental.
type CfnAgentApplicationLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnAgentApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

