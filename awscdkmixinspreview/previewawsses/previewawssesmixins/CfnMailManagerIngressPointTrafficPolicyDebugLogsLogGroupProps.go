package previewawssesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerIngressPointTrafficPolicyDebugLogsLogGroupProps := &CfnMailManagerIngressPointTrafficPolicyDebugLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_INGRESS_POINT_TYPE,
//   	},
//   }
//
// Experimental.
type CfnMailManagerIngressPointTrafficPolicyDebugLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

