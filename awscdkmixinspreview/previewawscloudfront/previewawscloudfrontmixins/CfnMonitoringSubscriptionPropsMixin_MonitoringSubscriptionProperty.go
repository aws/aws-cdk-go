package previewawscloudfrontmixins


// A monitoring subscription.
//
// This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   monitoringSubscriptionProperty := &MonitoringSubscriptionProperty{
//   	RealtimeMetricsSubscriptionConfig: &RealtimeMetricsSubscriptionConfigProperty{
//   		RealtimeMetricsSubscriptionStatus: jsii.String("realtimeMetricsSubscriptionStatus"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-monitoringsubscription-monitoringsubscription.html
//
type CfnMonitoringSubscriptionPropsMixin_MonitoringSubscriptionProperty struct {
	// A subscription configuration for additional CloudWatch metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-monitoringsubscription-monitoringsubscription.html#cfn-cloudfront-monitoringsubscription-monitoringsubscription-realtimemetricssubscriptionconfig
	//
	RealtimeMetricsSubscriptionConfig interface{} `field:"optional" json:"realtimeMetricsSubscriptionConfig" yaml:"realtimeMetricsSubscriptionConfig"`
}

