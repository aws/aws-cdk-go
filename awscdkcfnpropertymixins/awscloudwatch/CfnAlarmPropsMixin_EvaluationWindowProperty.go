package awscloudwatch


// The evaluation window that an alarm uses to select the range of metric data that it evaluates each time it runs.
//
// This is a union type. Set exactly one of its members, ``SlidingWindow`` or ``WallClockWindow``. If you don't set ``EvaluationWindow``, the alarm uses a ``SlidingWindow`` by default.
//  For more information, see [Alarm evaluation windows](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/alarm-evaluation-window.html) in the *CloudWatch User Guide*.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var slidingWindow interface{}
//
//   evaluationWindowProperty := &EvaluationWindowProperty{
//   	SlidingWindow: slidingWindow,
//   	WallClockWindow: &WallClockWindowProperty{
//   		Timezone: jsii.String("timezone"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-evaluationwindow.html
//
type CfnAlarmPropsMixin_EvaluationWindowProperty struct {
	// A sliding window, which advances each time the alarm is evaluated, forming a rolling time window.
	//
	// This is the default evaluation window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-evaluationwindow.html#cfn-cloudwatch-alarm-evaluationwindow-slidingwindow
	//
	SlidingWindow interface{} `field:"optional" json:"slidingWindow" yaml:"slidingWindow"`
	// An evaluation window that aligns the evaluated range to fixed clock boundaries that match the alarm's period, such as the top of the hour, midnight, or the start of the calendar week, optionally in a specific time zone.
	//
	// When you use a wall clock window, the alarm's period must be 1 minute (60 seconds), 5 minutes (300 seconds), 1 hour (3,600 seconds), 1 day (86,400 seconds), or 1 week (604,800 seconds). Other period values aren't supported with a wall clock window.
	//  Choose a wall clock window when your monitoring is tied to a business or calendar period, such as daily reports, batch jobs, or backups, or when you want alarm evaluations to match the periods shown on a metric dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-evaluationwindow.html#cfn-cloudwatch-alarm-evaluationwindow-wallclockwindow
	//
	WallClockWindow interface{} `field:"optional" json:"wallClockWindow" yaml:"wallClockWindow"`
}

