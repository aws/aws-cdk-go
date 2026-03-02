package previewawssesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerIngressPointTrafficPolicyDebugLogsDestProps := &CfnMailManagerIngressPointTrafficPolicyDebugLogsDestProps{
//   	RecordFields: []CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_INGRESS_POINT_TYPE,
//   	},
//   }
//
// Experimental.
type CfnMailManagerIngressPointTrafficPolicyDebugLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

