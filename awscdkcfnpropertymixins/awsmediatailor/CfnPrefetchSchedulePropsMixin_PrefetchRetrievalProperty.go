package awsmediatailor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   prefetchRetrievalProperty := &PrefetchRetrievalProperty{
//   	DynamicVariables: map[string]*string{
//   		"dynamicVariablesKey": jsii.String("dynamicVariables"),
//   	},
//   	EndTime: jsii.String("endTime"),
//   	StartTime: jsii.String("startTime"),
//   	TrafficShapingRetrievalWindow: &TrafficShapingRetrievalWindowProperty{
//   		RetrievalWindowDurationSeconds: jsii.Number(123),
//   	},
//   	TrafficShapingTpsConfiguration: &TrafficShapingTpsConfigurationProperty{
//   		PeakConcurrentUsers: jsii.Number(123),
//   		PeakTps: jsii.Number(123),
//   	},
//   	TrafficShapingType: jsii.String("trafficShapingType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-prefetchretrieval.html
//
type CfnPrefetchSchedulePropsMixin_PrefetchRetrievalProperty struct {
	// The dynamic variables to use for substitution during prefetch requests to the ad decision server (ADS).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-prefetchretrieval.html#cfn-mediatailor-prefetchschedule-prefetchretrieval-dynamicvariables
	//
	DynamicVariables interface{} `field:"optional" json:"dynamicVariables" yaml:"dynamicVariables"`
	// The time when prefetch retrieval ends for the ad break, as an ISO 8601 date-time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-prefetchretrieval.html#cfn-mediatailor-prefetchschedule-prefetchretrieval-endtime
	//
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The time when prefetch retrievals can start for this break, as an ISO 8601 date-time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-prefetchretrieval.html#cfn-mediatailor-prefetchschedule-prefetchretrieval-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-prefetchretrieval.html#cfn-mediatailor-prefetchschedule-prefetchretrieval-trafficshapingretrievalwindow
	//
	TrafficShapingRetrievalWindow interface{} `field:"optional" json:"trafficShapingRetrievalWindow" yaml:"trafficShapingRetrievalWindow"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-prefetchretrieval.html#cfn-mediatailor-prefetchschedule-prefetchretrieval-trafficshapingtpsconfiguration
	//
	TrafficShapingTpsConfiguration interface{} `field:"optional" json:"trafficShapingTpsConfiguration" yaml:"trafficShapingTpsConfiguration"`
	// Indicates the type of traffic shaping used to limit the number of requests to the ADS at one time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-prefetchretrieval.html#cfn-mediatailor-prefetchschedule-prefetchretrieval-trafficshapingtype
	//
	TrafficShapingType *string `field:"optional" json:"trafficShapingType" yaml:"trafficShapingType"`
}

