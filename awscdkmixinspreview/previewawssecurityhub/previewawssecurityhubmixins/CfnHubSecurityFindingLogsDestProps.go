package previewawssecurityhubmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHubSecurityFindingLogsDestProps := &CfnHubSecurityFindingLogsDestProps{
//   	RecordFields: []eVENT{
//   		awscdkmixinspreview.Mixins.CfnHubSecurityFindingLogsRecordFields_*eVENT,
//   	},
//   }
//
// Experimental.
type CfnHubSecurityFindingLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnHubSecurityFindingLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

