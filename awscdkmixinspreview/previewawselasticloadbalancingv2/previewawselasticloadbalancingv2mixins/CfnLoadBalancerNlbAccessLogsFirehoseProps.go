package previewawselasticloadbalancingv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerNlbAccessLogsFirehoseProps := &CfnLoadBalancerNlbAccessLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnLoadBalancerNlbAccessLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnLoadBalancerNlbAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnLoadBalancerNlbAccessLogsRecordFields_TYPE,
//   	},
//   }
//
// Experimental.
type CfnLoadBalancerNlbAccessLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnLoadBalancerNlbAccessLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLoadBalancerNlbAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

