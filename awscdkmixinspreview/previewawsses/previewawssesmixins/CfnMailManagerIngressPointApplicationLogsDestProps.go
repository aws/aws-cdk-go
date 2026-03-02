package previewawssesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerIngressPointApplicationLogsDestProps := &CfnMailManagerIngressPointApplicationLogsDestProps{
//   	RecordFields: []CfnMailManagerIngressPointApplicationLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnMailManagerIngressPointApplicationLogsRecordFields_INGRESS_POINT_TYPE,
//   	},
//   }
//
// Experimental.
type CfnMailManagerIngressPointApplicationLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMailManagerIngressPointApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

