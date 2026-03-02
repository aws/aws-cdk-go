package previewawssesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerIngressPointTrafficPolicyDebugLogsFirehoseProps := &CfnMailManagerIngressPointTrafficPolicyDebugLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_INGRESS_POINT_TYPE,
//   	},
//   }
//
// Experimental.
type CfnMailManagerIngressPointTrafficPolicyDebugLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

