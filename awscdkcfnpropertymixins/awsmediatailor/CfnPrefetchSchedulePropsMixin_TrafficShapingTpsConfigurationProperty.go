package awsmediatailor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   trafficShapingTpsConfigurationProperty := &TrafficShapingTpsConfigurationProperty{
//   	PeakConcurrentUsers: jsii.Number(123),
//   	PeakTps: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-trafficshapingtpsconfiguration.html
//
type CfnPrefetchSchedulePropsMixin_TrafficShapingTpsConfigurationProperty struct {
	// The expected peak number of concurrent viewers for your content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-trafficshapingtpsconfiguration.html#cfn-mediatailor-prefetchschedule-trafficshapingtpsconfiguration-peakconcurrentusers
	//
	PeakConcurrentUsers *float64 `field:"optional" json:"peakConcurrentUsers" yaml:"peakConcurrentUsers"`
	// The maximum number of transactions per second (TPS) that your ad decision server (ADS) can handle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-trafficshapingtpsconfiguration.html#cfn-mediatailor-prefetchschedule-trafficshapingtpsconfiguration-peaktps
	//
	PeakTps *float64 `field:"optional" json:"peakTps" yaml:"peakTps"`
}

