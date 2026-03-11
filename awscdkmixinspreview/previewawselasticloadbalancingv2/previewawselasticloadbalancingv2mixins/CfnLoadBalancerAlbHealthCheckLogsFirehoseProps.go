package previewawselasticloadbalancingv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerAlbHealthCheckLogsFirehoseProps := &CfnLoadBalancerAlbHealthCheckLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnLoadBalancerAlbHealthCheckLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnLoadBalancerAlbHealthCheckLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnLoadBalancerAlbHealthCheckLogsRecordFields_TYPE,
//   	},
//   }
//
// Experimental.
type CfnLoadBalancerAlbHealthCheckLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnLoadBalancerAlbHealthCheckLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLoadBalancerAlbHealthCheckLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

