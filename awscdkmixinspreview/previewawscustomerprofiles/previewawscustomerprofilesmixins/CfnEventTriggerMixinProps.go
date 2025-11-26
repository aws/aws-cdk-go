package previewawscustomerprofilesmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnEventTriggerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventTriggerMixinProps := &CfnEventTriggerMixinProps{
//   	Description: jsii.String("description"),
//   	DomainName: jsii.String("domainName"),
//   	EventTriggerConditions: []interface{}{
//   		&EventTriggerConditionProperty{
//   			EventTriggerDimensions: []interface{}{
//   				&EventTriggerDimensionProperty{
//   					ObjectAttributes: []interface{}{
//   						&ObjectAttributeProperty{
//   							ComparisonOperator: jsii.String("comparisonOperator"),
//   							FieldName: jsii.String("fieldName"),
//   							Source: jsii.String("source"),
//   							Values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			LogicalOperator: jsii.String("logicalOperator"),
//   		},
//   	},
//   	EventTriggerLimits: &EventTriggerLimitsProperty{
//   		EventExpiration: jsii.Number(123),
//   		Periods: []interface{}{
//   			&PeriodProperty{
//   				MaxInvocationsPerProfile: jsii.Number(123),
//   				Unit: jsii.String("unit"),
//   				Unlimited: jsii.Boolean(false),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	EventTriggerName: jsii.String("eventTriggerName"),
//   	ObjectTypeName: jsii.String("objectTypeName"),
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
type CfnEventTriggerMixinProps struct {
	// The description of the event trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The unique name of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// A list of conditions that determine when an event should trigger the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-eventtriggerconditions
	//
	EventTriggerConditions interface{} `field:"optional" json:"eventTriggerConditions" yaml:"eventTriggerConditions"`
	// Defines limits controlling whether an event triggers the destination, based on ingestion latency and the number of invocations per profile over specific time periods.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-eventtriggerlimits
	//
	EventTriggerLimits interface{} `field:"optional" json:"eventTriggerLimits" yaml:"eventTriggerLimits"`
	// The unique name of the event trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-eventtriggername
	//
	EventTriggerName *string `field:"optional" json:"eventTriggerName" yaml:"eventTriggerName"`
	// The unique name of the object type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-objecttypename
	//
	ObjectTypeName *string `field:"optional" json:"objectTypeName" yaml:"objectTypeName"`
	// The destination is triggered only for profiles that meet the criteria of a segment definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-segmentfilter
	//
	SegmentFilter *string `field:"optional" json:"segmentFilter" yaml:"segmentFilter"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-customerprofiles-eventtrigger.html#cfn-customerprofiles-eventtrigger-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

