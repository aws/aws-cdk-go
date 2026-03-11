package previewawselasticloadbalancingv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerAlbHealthCheckLogsDestProps := &CfnLoadBalancerAlbHealthCheckLogsDestProps{
//   	RecordFields: []CfnLoadBalancerAlbHealthCheckLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnLoadBalancerAlbHealthCheckLogsRecordFields_TYPE,
//   	},
//   }
//
// Experimental.
type CfnLoadBalancerAlbHealthCheckLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLoadBalancerAlbHealthCheckLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

