package previewawsconnectmixins


// Information about the hours of operation override config: day, start time, and end time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hoursOfOperationOverrideConfigProperty := &HoursOfOperationOverrideConfigProperty{
//   	Day: jsii.String("day"),
//   	EndTime: &OverrideTimeSliceProperty{
//   		Hours: jsii.Number(123),
//   		Minutes: jsii.Number(123),
//   	},
//   	StartTime: &OverrideTimeSliceProperty{
//   		Hours: jsii.Number(123),
//   		Minutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverrideconfig.html
//
type CfnHoursOfOperationPropsMixin_HoursOfOperationOverrideConfigProperty struct {
	// The day that the hours of operation override applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverrideconfig.html#cfn-connect-hoursofoperation-hoursofoperationoverrideconfig-day
	//
	Day *string `field:"optional" json:"day" yaml:"day"`
	// The end time that your contact center closes if overrides are applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverrideconfig.html#cfn-connect-hoursofoperation-hoursofoperationoverrideconfig-endtime
	//
	EndTime interface{} `field:"optional" json:"endTime" yaml:"endTime"`
	// The start time when your contact center opens if overrides are applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-hoursofoperation-hoursofoperationoverrideconfig.html#cfn-connect-hoursofoperation-hoursofoperationoverrideconfig-starttime
	//
	StartTime interface{} `field:"optional" json:"startTime" yaml:"startTime"`
}

