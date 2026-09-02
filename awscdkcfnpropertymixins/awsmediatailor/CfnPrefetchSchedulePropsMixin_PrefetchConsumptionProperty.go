package awsmediatailor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   prefetchConsumptionProperty := &PrefetchConsumptionProperty{
//   	AvailMatchingCriteria: []interface{}{
//   		&AvailMatchingCriteriaProperty{
//   			DynamicVariable: jsii.String("dynamicVariable"),
//   			Operator: jsii.String("operator"),
//   		},
//   	},
//   	EndTime: jsii.String("endTime"),
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-prefetchconsumption.html
//
type CfnPrefetchSchedulePropsMixin_PrefetchConsumptionProperty struct {
	// If you only want MediaTailor to insert prefetched ads into avails that match specific dynamic variables, set the avail matching criteria.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-prefetchconsumption.html#cfn-mediatailor-prefetchschedule-prefetchconsumption-availmatchingcriteria
	//
	AvailMatchingCriteria interface{} `field:"optional" json:"availMatchingCriteria" yaml:"availMatchingCriteria"`
	// The time when MediaTailor no longer considers the prefetched ads for use in an ad break, as an ISO 8601 date-time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-prefetchconsumption.html#cfn-mediatailor-prefetchschedule-prefetchconsumption-endtime
	//
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The time when prefetched ads are considered for use in an ad break, as an ISO 8601 date-time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-prefetchschedule-prefetchconsumption.html#cfn-mediatailor-prefetchschedule-prefetchconsumption-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

