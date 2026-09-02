package previewawssecurityhubmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHubV2SecurityFindingLogsDestProps := &CfnHubV2SecurityFindingLogsDestProps{
//   	RecordFields: []eVENT{
//   		awscdkmixinspreview.Mixins.CfnHubV2SecurityFindingLogsRecordFields_*eVENT,
//   	},
//   }
//
// Experimental.
type CfnHubV2SecurityFindingLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnHubV2SecurityFindingLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

