package awsmediatailor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   trafficShapingRetrievalWindowProperty := &TrafficShapingRetrievalWindowProperty{
//   	RetrievalWindowDurationSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-trafficshapingretrievalwindow.html
//
type CfnPrefetchSchedulePropsMixin_TrafficShapingRetrievalWindowProperty struct {
	// The amount of time, in seconds, that MediaTailor spreads prefetch requests to the ADS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-trafficshapingretrievalwindow.html#cfn-mediatailor-prefetchschedule-trafficshapingretrievalwindow-retrievalwindowdurationseconds
	//
	RetrievalWindowDurationSeconds *float64 `field:"optional" json:"retrievalWindowDurationSeconds" yaml:"retrievalWindowDurationSeconds"`
}

