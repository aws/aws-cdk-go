package awsmediatailor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPrefetchSchedulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnPrefetchScheduleMixinProps := &CfnPrefetchScheduleMixinProps{
//   	Consumption: &PrefetchConsumptionProperty{
//   		AvailMatchingCriteria: []interface{}{
//   			&AvailMatchingCriteriaProperty{
//   				DynamicVariable: jsii.String("dynamicVariable"),
//   				Operator: jsii.String("operator"),
//   			},
//   		},
//   		EndTime: jsii.String("endTime"),
//   		StartTime: jsii.String("startTime"),
//   	},
//   	Name: jsii.String("name"),
//   	PlaybackConfigurationName: jsii.String("playbackConfigurationName"),
//   	RecurringPrefetchConfiguration: &RecurringPrefetchConfigurationProperty{
//   		EndTime: jsii.String("endTime"),
//   		RecurringConsumption: &RecurringConsumptionProperty{
//   			AvailMatchingCriteria: []interface{}{
//   				&AvailMatchingCriteriaProperty{
//   					DynamicVariable: jsii.String("dynamicVariable"),
//   					Operator: jsii.String("operator"),
//   				},
//   			},
//   			RetrievedAdExpirationSeconds: jsii.Number(123),
//   		},
//   		RecurringRetrieval: &RecurringRetrievalProperty{
//   			DelayAfterAvailEndSeconds: jsii.Number(123),
//   			DynamicVariables: map[string]*string{
//   				"dynamicVariablesKey": jsii.String("dynamicVariables"),
//   			},
//   			TrafficShapingRetrievalWindow: &TrafficShapingRetrievalWindowProperty{
//   				RetrievalWindowDurationSeconds: jsii.Number(123),
//   			},
//   			TrafficShapingTpsConfiguration: &TrafficShapingTpsConfigurationProperty{
//   				PeakConcurrentUsers: jsii.Number(123),
//   				PeakTps: jsii.Number(123),
//   			},
//   			TrafficShapingType: jsii.String("trafficShapingType"),
//   		},
//   		StartTime: jsii.String("startTime"),
//   	},
//   	Retrieval: &PrefetchRetrievalProperty{
//   		DynamicVariables: map[string]*string{
//   			"dynamicVariablesKey": jsii.String("dynamicVariables"),
//   		},
//   		EndTime: jsii.String("endTime"),
//   		StartTime: jsii.String("startTime"),
//   		TrafficShapingRetrievalWindow: &TrafficShapingRetrievalWindowProperty{
//   			RetrievalWindowDurationSeconds: jsii.Number(123),
//   		},
//   		TrafficShapingTpsConfiguration: &TrafficShapingTpsConfigurationProperty{
//   			PeakConcurrentUsers: jsii.Number(123),
//   			PeakTps: jsii.Number(123),
//   		},
//   		TrafficShapingType: jsii.String("trafficShapingType"),
//   	},
//   	ScheduleType: jsii.String("scheduleType"),
//   	StreamId: jsii.String("streamId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-prefetchschedule.html
//
type CfnPrefetchScheduleMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-prefetchschedule.html#cfn-mediatailor-prefetchschedule-consumption
	//
	Consumption interface{} `field:"optional" json:"consumption" yaml:"consumption"`
	// The name to assign to the prefetch schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-prefetchschedule.html#cfn-mediatailor-prefetchschedule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of the playback configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-prefetchschedule.html#cfn-mediatailor-prefetchschedule-playbackconfigurationname
	//
	PlaybackConfigurationName *string `field:"optional" json:"playbackConfigurationName" yaml:"playbackConfigurationName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-prefetchschedule.html#cfn-mediatailor-prefetchschedule-recurringprefetchconfiguration
	//
	RecurringPrefetchConfiguration interface{} `field:"optional" json:"recurringPrefetchConfiguration" yaml:"recurringPrefetchConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-prefetchschedule.html#cfn-mediatailor-prefetchschedule-retrieval
	//
	Retrieval interface{} `field:"optional" json:"retrieval" yaml:"retrieval"`
	// The frequency that MediaTailor creates prefetch schedules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-prefetchschedule.html#cfn-mediatailor-prefetchschedule-scheduletype
	//
	ScheduleType *string `field:"optional" json:"scheduleType" yaml:"scheduleType"`
	// An optional stream identifier that MediaTailor uses to prefetch ads for multiple streams that use the same playback configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-prefetchschedule.html#cfn-mediatailor-prefetchschedule-streamid
	//
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// The tags assigned to the prefetch schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-prefetchschedule.html#cfn-mediatailor-prefetchschedule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

