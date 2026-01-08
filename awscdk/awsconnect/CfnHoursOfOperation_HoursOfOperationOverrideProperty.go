package awsconnect


// Information about the hours of operations override.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hoursOfOperationOverrideProperty := &HoursOfOperationOverrideProperty{
//   	EffectiveFrom: jsii.String("effectiveFrom"),
//   	EffectiveTill: jsii.String("effectiveTill"),
//   	OverrideConfig: []interface{}{
//   		&HoursOfOperationOverrideConfigProperty{
//   			Day: jsii.String("day"),
//   			EndTime: &OverrideTimeSliceProperty{
//   				Hours: jsii.Number(123),
//   				Minutes: jsii.Number(123),
//   			},
//   			StartTime: &OverrideTimeSliceProperty{
//   				Hours: jsii.Number(123),
//   				Minutes: jsii.Number(123),
//   			},
//   		},
//   	},
//   	OverrideName: jsii.String("overrideName"),
//
//   	// the properties below are optional
//   	HoursOfOperationOverrideId: jsii.String("hoursOfOperationOverrideId"),
//   	OverrideDescription: jsii.String("overrideDescription"),
//   	OverrideType: jsii.String("overrideType"),
//   	RecurrenceConfig: &RecurrenceConfigProperty{
//   		RecurrencePattern: &RecurrencePatternProperty{
//   			ByMonth: []interface{}{
//   				jsii.Number(123),
//   			},
//   			ByMonthDay: []interface{}{
//   				jsii.Number(123),
//   			},
//   			ByWeekdayOccurrence: []interface{}{
//   				jsii.Number(123),
//   			},
//   			Frequency: jsii.String("frequency"),
//   			Interval: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html
//
type CfnHoursOfOperation_HoursOfOperationOverrideProperty struct {
	// The date from which the hours of operation override would be effective.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html#cfn-connect-hoursofoperation-hoursofoperationoverride-effectivefrom
	//
	EffectiveFrom *string `field:"required" json:"effectiveFrom" yaml:"effectiveFrom"`
	// The date until the hours of operation override is effective.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html#cfn-connect-hoursofoperation-hoursofoperationoverride-effectivetill
	//
	EffectiveTill *string `field:"required" json:"effectiveTill" yaml:"effectiveTill"`
	// Configuration information for the hours of operation override: day, start time, and end time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html#cfn-connect-hoursofoperation-hoursofoperationoverride-overrideconfig
	//
	OverrideConfig interface{} `field:"required" json:"overrideConfig" yaml:"overrideConfig"`
	// The name of the hours of operation override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html#cfn-connect-hoursofoperation-hoursofoperationoverride-overridename
	//
	OverrideName *string `field:"required" json:"overrideName" yaml:"overrideName"`
	// The identifier for the hours of operation override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html#cfn-connect-hoursofoperation-hoursofoperationoverride-hoursofoperationoverrideid
	//
	HoursOfOperationOverrideId *string `field:"optional" json:"hoursOfOperationOverrideId" yaml:"hoursOfOperationOverrideId"`
	// The description of the hours of operation override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html#cfn-connect-hoursofoperation-hoursofoperationoverride-overridedescription
	//
	OverrideDescription *string `field:"optional" json:"overrideDescription" yaml:"overrideDescription"`
	// The type of hours of operation override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html#cfn-connect-hoursofoperation-hoursofoperationoverride-overridetype
	//
	OverrideType *string `field:"optional" json:"overrideType" yaml:"overrideType"`
	// Configuration for recurring hours of operation overrides.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html#cfn-connect-hoursofoperation-hoursofoperationoverride-recurrenceconfig
	//
	RecurrenceConfig interface{} `field:"optional" json:"recurrenceConfig" yaml:"recurrenceConfig"`
}

