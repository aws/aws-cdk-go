package awscloudwatch


// An evaluation window that aligns the evaluated range to fixed clock boundaries that match the alarm's period, such as the top of the hour, midnight, or the start of the calendar week, optionally in a specific time zone.
//
// When you use a wall clock window, the alarm's period must be 1 minute (60 seconds), 5 minutes (300 seconds), 1 hour (3,600 seconds), 1 day (86,400 seconds), or 1 week (604,800 seconds). Other period values aren't supported with a wall clock window.
//  Choose a wall clock window when your monitoring is tied to a business or calendar period, such as daily reports, batch jobs, or backups, or when you want alarm evaluations to match the periods shown on a metric dashboard.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wallClockWindowProperty := &WallClockWindowProperty{
//   	Timezone: jsii.String("timezone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-wallclockwindow.html
//
type CfnAlarm_WallClockWindowProperty struct {
	// The time zone to use when the alarm aligns the evaluation window to clock boundaries.
	//
	// You can specify an IANA time zone name (for example, ``America/New_York``), a fixed UTC offset (for example, ``+05:30``), or an offset-prefixed identifier (for example, ``UTC+05:30``). The offset must be aligned to a multiple of 5 minutes. If you don't specify a time zone, CloudWatch uses ``UTC``.
	//  The time zone affects window alignment for all periods, including periods of one hour or shorter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-wallclockwindow.html#cfn-cloudwatch-alarm-wallclockwindow-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

