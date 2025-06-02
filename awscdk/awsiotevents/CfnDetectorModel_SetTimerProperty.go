package awsiotevents


// Information needed to set the timer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   setTimerProperty := &SetTimerProperty{
//   	TimerName: jsii.String("timerName"),
//
//   	// the properties below are optional
//   	DurationExpression: jsii.String("durationExpression"),
//   	Seconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-settimer.html
//
type CfnDetectorModel_SetTimerProperty struct {
	// The name of the timer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-settimer.html#cfn-iotevents-detectormodel-settimer-timername
	//
	TimerName *string `field:"required" json:"timerName" yaml:"timerName"`
	// The duration of the timer, in seconds.
	//
	// You can use a string expression that includes numbers, variables ( `$variable.<variable-name>` ), and input values ( `$input.<input-name>.<path-to-datum>` ) as the duration. The range of the duration is 1-31622400 seconds. To ensure accuracy, the minimum duration is 60 seconds. The evaluated result of the duration is rounded down to the nearest whole number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-settimer.html#cfn-iotevents-detectormodel-settimer-durationexpression
	//
	DurationExpression *string `field:"optional" json:"durationExpression" yaml:"durationExpression"`
	// The number of seconds until the timer expires.
	//
	// The minimum value is 60 seconds to ensure accuracy. The maximum value is 31622400 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-settimer.html#cfn-iotevents-detectormodel-settimer-seconds
	//
	Seconds *float64 `field:"optional" json:"seconds" yaml:"seconds"`
}

