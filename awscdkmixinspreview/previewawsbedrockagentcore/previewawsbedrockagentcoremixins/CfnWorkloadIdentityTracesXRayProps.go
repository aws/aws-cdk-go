package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkloadIdentityTracesXRayProps := &CfnWorkloadIdentityTracesXRayProps{
//   	RecordFields: []CfnWorkloadIdentityTracesRecordFields{
//   		awscdkmixinspreview.Mixins.CfnWorkloadIdentityTracesRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnWorkloadIdentityTracesXRayProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnWorkloadIdentityTracesRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

