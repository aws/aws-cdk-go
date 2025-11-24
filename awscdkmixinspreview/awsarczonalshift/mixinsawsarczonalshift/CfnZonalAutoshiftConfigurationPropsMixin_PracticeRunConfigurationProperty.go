package mixinsawsarczonalshift


// A practice run configuration for a resource includes the Amazon CloudWatch alarms that you've specified for a practice run, as well as any blocked dates or blocked windows for the practice run.
//
// When a resource has a practice run configuation, ARC starts weekly zonal shifts for the resource, to shift traffic away from an Availability Zone. Weekly practice runs help you to make sure that your application can continue to operate normally with the loss of one Availability Zone.
//
// You can update or delete a practice run configuration. When you delete a practice run configuration, zonal autoshift is disabled for the resource. A practice run configuration is required when zonal autoshift is enabled.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   practiceRunConfigurationProperty := &PracticeRunConfigurationProperty{
//   	BlockedDates: []*string{
//   		jsii.String("blockedDates"),
//   	},
//   	BlockedWindows: []*string{
//   		jsii.String("blockedWindows"),
//   	},
//   	BlockingAlarms: []interface{}{
//   		&ControlConditionProperty{
//   			AlarmIdentifier: jsii.String("alarmIdentifier"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	OutcomeAlarms: []interface{}{
//   		&ControlConditionProperty{
//   			AlarmIdentifier: jsii.String("alarmIdentifier"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arczonalshift-zonalautoshiftconfiguration-practicerunconfiguration.html
//
type CfnZonalAutoshiftConfigurationPropsMixin_PracticeRunConfigurationProperty struct {
	// An array of one or more dates that you can specify when AWS does not start practice runs for a resource.
	//
	// Dates are in UTC.
	//
	// Specify blocked dates in the format `YYYY-MM-DD` , separated by spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arczonalshift-zonalautoshiftconfiguration-practicerunconfiguration.html#cfn-arczonalshift-zonalautoshiftconfiguration-practicerunconfiguration-blockeddates
	//
	BlockedDates *[]*string `field:"optional" json:"blockedDates" yaml:"blockedDates"`
	// An array of one or more days and times that you can specify when ARC does not start practice runs for a resource.
	//
	// Days and times are in UTC.
	//
	// Specify blocked windows in the format `DAY:HH:MM-DAY:HH:MM` , separated by spaces. For example, `MON:18:30-MON:19:30 TUE:18:30-TUE:19:30` .
	//
	// > Blocked windows have to start and end on the same day. Windows that span multiple days aren't supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arczonalshift-zonalautoshiftconfiguration-practicerunconfiguration.html#cfn-arczonalshift-zonalautoshiftconfiguration-practicerunconfiguration-blockedwindows
	//
	BlockedWindows *[]*string `field:"optional" json:"blockedWindows" yaml:"blockedWindows"`
	// An optional alarm that you can specify that blocks practice runs when the alarm is in an `ALARM` state.
	//
	// When a blocking alarm goes into an `ALARM` state, it prevents practice runs from being started, and ends practice runs that are in progress.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arczonalshift-zonalautoshiftconfiguration-practicerunconfiguration.html#cfn-arczonalshift-zonalautoshiftconfiguration-practicerunconfiguration-blockingalarms
	//
	BlockingAlarms interface{} `field:"optional" json:"blockingAlarms" yaml:"blockingAlarms"`
	// The alarm that you specify to monitor the health of your application during practice runs.
	//
	// When the outcome alarm goes into an `ALARM` state, the practice run is ended and the outcome is set to `FAILED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arczonalshift-zonalautoshiftconfiguration-practicerunconfiguration.html#cfn-arczonalshift-zonalautoshiftconfiguration-practicerunconfiguration-outcomealarms
	//
	OutcomeAlarms interface{} `field:"optional" json:"outcomeAlarms" yaml:"outcomeAlarms"`
}

