package awscloudwatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   wallClockWindowProperty := &WallClockWindowProperty{
//   	Timezone: jsii.String("timezone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-wallclockwindow.html
//
type CfnAlarmPropsMixin_WallClockWindowProperty struct {
	// The timezone for wall clock evaluation, in IANA time zone format (e.g., America/New_York, UTC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alarm-wallclockwindow.html#cfn-cloudwatch-alarm-wallclockwindow-timezone
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

