package previewawsconnectmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnHoursOfOperationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHoursOfOperationMixinProps := &CfnHoursOfOperationMixinProps{
//   	ChildHoursOfOperations: []interface{}{
//   		&HoursOfOperationsIdentifierProperty{
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Config: []interface{}{
//   		&HoursOfOperationConfigProperty{
//   			Day: jsii.String("day"),
//   			EndTime: &HoursOfOperationTimeSliceProperty{
//   				Hours: jsii.Number(123),
//   				Minutes: jsii.Number(123),
//   			},
//   			StartTime: &HoursOfOperationTimeSliceProperty{
//   				Hours: jsii.Number(123),
//   				Minutes: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	HoursOfOperationOverrides: []interface{}{
//   		&HoursOfOperationOverrideProperty{
//   			EffectiveFrom: jsii.String("effectiveFrom"),
//   			EffectiveTill: jsii.String("effectiveTill"),
//   			HoursOfOperationOverrideId: jsii.String("hoursOfOperationOverrideId"),
//   			OverrideConfig: []interface{}{
//   				&HoursOfOperationOverrideConfigProperty{
//   					Day: jsii.String("day"),
//   					EndTime: &OverrideTimeSliceProperty{
//   						Hours: jsii.Number(123),
//   						Minutes: jsii.Number(123),
//   					},
//   					StartTime: &OverrideTimeSliceProperty{
//   						Hours: jsii.Number(123),
//   						Minutes: jsii.Number(123),
//   					},
//   				},
//   			},
//   			OverrideDescription: jsii.String("overrideDescription"),
//   			OverrideName: jsii.String("overrideName"),
//   			OverrideType: jsii.String("overrideType"),
//   			RecurrenceConfig: &RecurrenceConfigProperty{
//   				RecurrencePattern: &RecurrencePatternProperty{
//   					ByMonth: []interface{}{
//   						jsii.Number(123),
//   					},
//   					ByMonthDay: []interface{}{
//   						jsii.Number(123),
//   					},
//   					ByWeekdayOccurrence: []interface{}{
//   						jsii.Number(123),
//   					},
//   					Frequency: jsii.String("frequency"),
//   					Interval: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	ParentHoursOfOperations: []interface{}{
//   		&HoursOfOperationsIdentifierProperty{
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeZone: jsii.String("timeZone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-hoursofoperation.html
//
type CfnHoursOfOperationMixinProps struct {
	// List of child hours of operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-hoursofoperation.html#cfn-connect-hoursofoperation-childhoursofoperations
	//
	ChildHoursOfOperations interface{} `field:"optional" json:"childHoursOfOperations" yaml:"childHoursOfOperations"`
	// Configuration information for the hours of operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-hoursofoperation.html#cfn-connect-hoursofoperation-config
	//
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// The description for the hours of operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-hoursofoperation.html#cfn-connect-hoursofoperation-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// One or more hours of operation overrides assigned to an hour of operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-hoursofoperation.html#cfn-connect-hoursofoperation-hoursofoperationoverrides
	//
	HoursOfOperationOverrides interface{} `field:"optional" json:"hoursOfOperationOverrides" yaml:"hoursOfOperationOverrides"`
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-hoursofoperation.html#cfn-connect-hoursofoperation-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// The name for the hours of operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-hoursofoperation.html#cfn-connect-hoursofoperation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// List of parent hours of operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-hoursofoperation.html#cfn-connect-hoursofoperation-parenthoursofoperations
	//
	ParentHoursOfOperations interface{} `field:"optional" json:"parentHoursOfOperations" yaml:"parentHoursOfOperations"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "Tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-hoursofoperation.html#cfn-connect-hoursofoperation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The time zone for the hours of operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-hoursofoperation.html#cfn-connect-hoursofoperation-timezone
	//
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

