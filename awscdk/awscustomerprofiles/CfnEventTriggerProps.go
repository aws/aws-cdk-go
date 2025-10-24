package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEventTrigger`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventTriggerProps := &CfnEventTriggerProps{
//   	DomainName: jsii.String("domainName"),
//   	EventTriggerConditions: []interface{}{
//   		&EventTriggerConditionProperty{
//   			EventTriggerDimensions: []interface{}{
//   				&EventTriggerDimensionProperty{
//   					ObjectAttributes: []interface{}{
//   						&ObjectAttributeProperty{
//   							ComparisonOperator: jsii.String("comparisonOperator"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//
//   							// the properties below are optional
//   							FieldName: jsii.String("fieldName"),
//   							Source: jsii.String("source"),
//   						},
//   					},
//   				},
//   			},
//   			LogicalOperator: jsii.String("logicalOperator"),
//   		},
//   	},
//   	EventTriggerName: jsii.String("eventTriggerName"),
//   	ObjectTypeName: jsii.String("objectTypeName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EventTriggerLimits: &EventTriggerLimitsProperty{
//   		EventExpiration: jsii.Number(123),
//   		Periods: []interface{}{
//   			&PeriodProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//
//   				// the properties below are optional
//   				MaxInvocationsPerProfile: jsii.Number(123),
//   				Unlimited: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	SegmentFilter: jsii.String("segmentFilter"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html
//
type CfnEventTriggerProps struct {
	// The unique name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// A list of conditions that determine when an event should trigger the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-eventtriggerconditions
	//
	EventTriggerConditions interface{} `field:"required" json:"eventTriggerConditions" yaml:"eventTriggerConditions"`
	// The unique name of the event trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-eventtriggername
	//
	EventTriggerName *string `field:"required" json:"eventTriggerName" yaml:"eventTriggerName"`
	// The unique name of the object type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-objecttypename
	//
	ObjectTypeName *string `field:"required" json:"objectTypeName" yaml:"objectTypeName"`
	// The description of the event trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Defines limits controlling whether an event triggers the destination, based on ingestion latency and the number of invocations per profile over specific time periods.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-eventtriggerlimits
	//
	EventTriggerLimits interface{} `field:"optional" json:"eventTriggerLimits" yaml:"eventTriggerLimits"`
	// The destination is triggered only for profiles that meet the criteria of a segment definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-segmentfilter
	//
	SegmentFilter *string `field:"optional" json:"segmentFilter" yaml:"segmentFilter"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

