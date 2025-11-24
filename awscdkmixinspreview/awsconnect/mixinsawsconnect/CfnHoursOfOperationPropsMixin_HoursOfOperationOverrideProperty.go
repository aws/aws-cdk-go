package mixinsawsconnect


// Information about the hours of operations override.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hoursOfOperationOverrideProperty := &HoursOfOperationOverrideProperty{
//   	EffectiveFrom: jsii.String("effectiveFrom"),
//   	EffectiveTill: jsii.String("effectiveTill"),
//   	HoursOfOperationOverrideId: jsii.String("hoursOfOperationOverrideId"),
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
//   	OverrideDescription: jsii.String("overrideDescription"),
//   	OverrideName: jsii.String("overrideName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html
//
type CfnHoursOfOperationPropsMixin_HoursOfOperationOverrideProperty struct {
	// The date from which the hours of operation override would be effective.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html#cfn-connect-hoursofoperation-hoursofoperationoverride-effectivefrom
	//
	EffectiveFrom *string `field:"optional" json:"effectiveFrom" yaml:"effectiveFrom"`
	// The date until the hours of operation override is effective.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html#cfn-connect-hoursofoperation-hoursofoperationoverride-effectivetill
	//
	EffectiveTill *string `field:"optional" json:"effectiveTill" yaml:"effectiveTill"`
	// The identifier for the hours of operation override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html#cfn-connect-hoursofoperation-hoursofoperationoverride-hoursofoperationoverrideid
	//
	HoursOfOperationOverrideId *string `field:"optional" json:"hoursOfOperationOverrideId" yaml:"hoursOfOperationOverrideId"`
	// Configuration information for the hours of operation override: day, start time, and end time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html#cfn-connect-hoursofoperation-hoursofoperationoverride-overrideconfig
	//
	OverrideConfig interface{} `field:"optional" json:"overrideConfig" yaml:"overrideConfig"`
	// The description of the hours of operation override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html#cfn-connect-hoursofoperation-hoursofoperationoverride-overridedescription
	//
	OverrideDescription *string `field:"optional" json:"overrideDescription" yaml:"overrideDescription"`
	// The name of the hours of operation override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html#cfn-connect-hoursofoperation-hoursofoperationoverride-overridename
	//
	OverrideName *string `field:"optional" json:"overrideName" yaml:"overrideName"`
}

