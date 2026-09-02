package awsmediatailor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   recurringRetrievalProperty := &RecurringRetrievalProperty{
//   	DelayAfterAvailEndSeconds: jsii.Number(123),
//   	DynamicVariables: map[string]*string{
//   		"dynamicVariablesKey": jsii.String("dynamicVariables"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-recurringretrieval.html
//
type CfnPrefetchSchedulePropsMixin_RecurringRetrievalProperty struct {
	// The number of seconds that MediaTailor waits after an ad avail before prefetching ads for the next avail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-recurringretrieval.html#cfn-mediatailor-prefetchschedule-recurringretrieval-delayafteravailendseconds
	//
	DelayAfterAvailEndSeconds *float64 `field:"optional" json:"delayAfterAvailEndSeconds" yaml:"delayAfterAvailEndSeconds"`
	// The dynamic variables to use for substitution during prefetch requests to the ADS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-recurringretrieval.html#cfn-mediatailor-prefetchschedule-recurringretrieval-dynamicvariables
	//
	DynamicVariables interface{} `field:"optional" json:"dynamicVariables" yaml:"dynamicVariables"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-recurringretrieval.html#cfn-mediatailor-prefetchschedule-recurringretrieval-trafficshapingretrievalwindow
	//
	TrafficShapingRetrievalWindow interface{} `field:"optional" json:"trafficShapingRetrievalWindow" yaml:"trafficShapingRetrievalWindow"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-recurringretrieval.html#cfn-mediatailor-prefetchschedule-recurringretrieval-trafficshapingtpsconfiguration
	//
	TrafficShapingTpsConfiguration interface{} `field:"optional" json:"trafficShapingTpsConfiguration" yaml:"trafficShapingTpsConfiguration"`
	// Indicates the type of traffic shaping used to limit the number of requests to the ADS at one time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-recurringretrieval.html#cfn-mediatailor-prefetchschedule-recurringretrieval-trafficshapingtype
	//
	TrafficShapingType *string `field:"optional" json:"trafficShapingType" yaml:"trafficShapingType"`
}

