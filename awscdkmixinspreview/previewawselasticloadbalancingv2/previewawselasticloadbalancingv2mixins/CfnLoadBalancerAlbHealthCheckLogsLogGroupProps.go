package previewawselasticloadbalancingv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerAlbHealthCheckLogsLogGroupProps := &CfnLoadBalancerAlbHealthCheckLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnLoadBalancerAlbHealthCheckLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnLoadBalancerAlbHealthCheckLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnLoadBalancerAlbHealthCheckLogsRecordFields_TYPE,
//   	},
//   }
//
// Experimental.
type CfnLoadBalancerAlbHealthCheckLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnLoadBalancerAlbHealthCheckLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLoadBalancerAlbHealthCheckLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

