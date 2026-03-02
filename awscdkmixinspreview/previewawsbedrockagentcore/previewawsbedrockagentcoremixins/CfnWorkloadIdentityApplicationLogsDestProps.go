package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkloadIdentityApplicationLogsDestProps := &CfnWorkloadIdentityApplicationLogsDestProps{
//   	RecordFields: []CfnWorkloadIdentityApplicationLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnWorkloadIdentityApplicationLogsRecordFields_REQUEST_ID,
//   	},
//   }
//
// Experimental.
type CfnWorkloadIdentityApplicationLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnWorkloadIdentityApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

