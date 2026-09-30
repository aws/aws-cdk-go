package awsbcm


// The window during which the schedule is active.
//
// When omitted the service defaults it. EndTime must be within three years of the time of the request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   schedulePeriodProperty := &SchedulePeriodProperty{
//   	EndTime: jsii.String("endTime"),
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-scheduleperiod.html
//
type CfnScheduledReportPropsMixin_SchedulePeriodProperty struct {
	// The time at which the schedule stops being active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-scheduleperiod.html#cfn-bcm-scheduledreport-scheduleperiod-endtime
	//
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The time at which the schedule becomes active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-scheduledreport-scheduleperiod.html#cfn-bcm-scheduledreport-scheduleperiod-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

