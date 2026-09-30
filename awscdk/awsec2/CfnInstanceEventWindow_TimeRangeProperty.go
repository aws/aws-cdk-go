package awsec2


// A time range during which an AWS-initiated maintenance event may occur.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeRangeProperty := &TimeRangeProperty{
//   	EndHour: jsii.Number(123),
//   	EndWeekDay: jsii.String("endWeekDay"),
//   	StartHour: jsii.Number(123),
//   	StartWeekDay: jsii.String("startWeekDay"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instanceeventwindow-timerange.html
//
type CfnInstanceEventWindow_TimeRangeProperty struct {
	// The hour when the time range ends.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instanceeventwindow-timerange.html#cfn-ec2-instanceeventwindow-timerange-endhour
	//
	EndHour *float64 `field:"optional" json:"endHour" yaml:"endHour"`
	// The day on which the time range ends.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instanceeventwindow-timerange.html#cfn-ec2-instanceeventwindow-timerange-endweekday
	//
	EndWeekDay *string `field:"optional" json:"endWeekDay" yaml:"endWeekDay"`
	// The hour when the time range begins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instanceeventwindow-timerange.html#cfn-ec2-instanceeventwindow-timerange-starthour
	//
	StartHour *float64 `field:"optional" json:"startHour" yaml:"startHour"`
	// The day on which the time range begins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instanceeventwindow-timerange.html#cfn-ec2-instanceeventwindow-timerange-startweekday
	//
	StartWeekDay *string `field:"optional" json:"startWeekDay" yaml:"startWeekDay"`
}

