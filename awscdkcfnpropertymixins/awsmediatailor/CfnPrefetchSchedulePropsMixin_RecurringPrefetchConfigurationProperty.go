package awsmediatailor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   recurringPrefetchConfigurationProperty := &RecurringPrefetchConfigurationProperty{
//   	EndTime: jsii.String("endTime"),
//   	RecurringConsumption: &RecurringConsumptionProperty{
//   		AvailMatchingCriteria: []interface{}{
//   			&AvailMatchingCriteriaProperty{
//   				DynamicVariable: jsii.String("dynamicVariable"),
//   				Operator: jsii.String("operator"),
//   			},
//   		},
//   		RetrievedAdExpirationSeconds: jsii.Number(123),
//   	},
//   	RecurringRetrieval: &RecurringRetrievalProperty{
//   		DelayAfterAvailEndSeconds: jsii.Number(123),
//   		DynamicVariables: map[string]*string{
//   			"dynamicVariablesKey": jsii.String("dynamicVariables"),
//   		},
//   		TrafficShapingRetrievalWindow: &TrafficShapingRetrievalWindowProperty{
//   			RetrievalWindowDurationSeconds: jsii.Number(123),
//   		},
//   		TrafficShapingTpsConfiguration: &TrafficShapingTpsConfigurationProperty{
//   			PeakConcurrentUsers: jsii.Number(123),
//   			PeakTps: jsii.Number(123),
//   		},
//   		TrafficShapingType: jsii.String("trafficShapingType"),
//   	},
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-recurringprefetchconfiguration.html
//
type CfnPrefetchSchedulePropsMixin_RecurringPrefetchConfigurationProperty struct {
	// The end time for the window that MediaTailor prefetches and inserts ads in a live event, as an ISO 8601 date-time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-recurringprefetchconfiguration.html#cfn-mediatailor-prefetchschedule-recurringprefetchconfiguration-endtime
	//
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-recurringprefetchconfiguration.html#cfn-mediatailor-prefetchschedule-recurringprefetchconfiguration-recurringconsumption
	//
	RecurringConsumption interface{} `field:"optional" json:"recurringConsumption" yaml:"recurringConsumption"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-recurringprefetchconfiguration.html#cfn-mediatailor-prefetchschedule-recurringprefetchconfiguration-recurringretrieval
	//
	RecurringRetrieval interface{} `field:"optional" json:"recurringRetrieval" yaml:"recurringRetrieval"`
	// The start time for the window that MediaTailor prefetches and inserts ads in a live event, as an ISO 8601 date-time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-recurringprefetchconfiguration.html#cfn-mediatailor-prefetchschedule-recurringprefetchconfiguration-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

