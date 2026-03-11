package previewawselasticloadbalancingv2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerAlbAccessLogsFirehoseProps := &CfnLoadBalancerAlbAccessLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnLoadBalancerAlbAccessLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnLoadBalancerAlbAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnLoadBalancerAlbAccessLogsRecordFields_TYPE,
//   	},
//   }
//
// Experimental.
type CfnLoadBalancerAlbAccessLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnLoadBalancerAlbAccessLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLoadBalancerAlbAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

