package previewawselasticloadbalancingv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerAlbAccessLogsLogGroupProps := &CfnLoadBalancerAlbAccessLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnLoadBalancerAlbAccessLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnLoadBalancerAlbAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnLoadBalancerAlbAccessLogsRecordFields_TYPE,
//   	},
//   }
//
// Experimental.
type CfnLoadBalancerAlbAccessLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnLoadBalancerAlbAccessLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLoadBalancerAlbAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

