package previewawselasticloadbalancingv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerNlbAccessLogsDestProps := &CfnLoadBalancerNlbAccessLogsDestProps{
//   	RecordFields: []CfnLoadBalancerNlbAccessLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnLoadBalancerNlbAccessLogsRecordFields_TYPE,
//   	},
//   }
//
// Experimental.
type CfnLoadBalancerNlbAccessLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLoadBalancerNlbAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

