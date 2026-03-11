package previewawselasticloadbalancingv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerAlbConnectionLogsLogGroupProps := &CfnLoadBalancerAlbConnectionLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnLoadBalancerAlbConnectionLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnLoadBalancerAlbConnectionLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnLoadBalancerAlbConnectionLogsRecordFields_TIME,
//   	},
//   }
//
// Experimental.
type CfnLoadBalancerAlbConnectionLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnLoadBalancerAlbConnectionLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLoadBalancerAlbConnectionLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

