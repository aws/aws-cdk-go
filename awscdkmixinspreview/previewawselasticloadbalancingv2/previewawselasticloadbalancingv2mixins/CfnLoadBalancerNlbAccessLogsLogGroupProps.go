package previewawselasticloadbalancingv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerNlbAccessLogsLogGroupProps := &CfnLoadBalancerNlbAccessLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnLoadBalancerNlbAccessLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnLoadBalancerNlbAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnLoadBalancerNlbAccessLogsRecordFields_TYPE,
//   	},
//   }
//
// Experimental.
type CfnLoadBalancerNlbAccessLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnLoadBalancerNlbAccessLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLoadBalancerNlbAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

