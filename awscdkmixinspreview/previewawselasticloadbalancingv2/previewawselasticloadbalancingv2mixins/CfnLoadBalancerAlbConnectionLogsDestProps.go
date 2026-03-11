package previewawselasticloadbalancingv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerAlbConnectionLogsDestProps := &CfnLoadBalancerAlbConnectionLogsDestProps{
//   	RecordFields: []CfnLoadBalancerAlbConnectionLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnLoadBalancerAlbConnectionLogsRecordFields_TIME,
//   	},
//   }
//
// Experimental.
type CfnLoadBalancerAlbConnectionLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLoadBalancerAlbConnectionLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

