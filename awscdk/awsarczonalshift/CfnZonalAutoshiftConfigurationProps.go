package awsarczonalshift


// Properties for defining a `CfnZonalAutoshiftConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnZonalAutoshiftConfigurationProps := &CfnZonalAutoshiftConfigurationProps{
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//
//   	// the properties below are optional
//   	PracticeRunConfiguration: &PracticeRunConfigurationProperty{
//   		OutcomeAlarms: []interface{}{
//   			&ControlConditionProperty{
//   				AlarmIdentifier: jsii.String("alarmIdentifier"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		BlockedDates: []*string{
//   			jsii.String("blockedDates"),
//   		},
//   		BlockedWindows: []*string{
//   			jsii.String("blockedWindows"),
//   		},
//   		BlockingAlarms: []interface{}{
//   			&ControlConditionProperty{
//   				AlarmIdentifier: jsii.String("alarmIdentifier"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	ZonalAutoshiftStatus: jsii.String("zonalAutoshiftStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arczonalshift-zonalautoshiftconfiguration.html
//
type CfnZonalAutoshiftConfigurationProps struct {
	// The identifier for the resource that AWS shifts traffic for.
	//
	// The identifier is the Amazon Resource Name (ARN) for the resource.
	//
	// At this time, supported resources are Network Load Balancers and Application Load Balancers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arczonalshift-zonalautoshiftconfiguration.html#cfn-arczonalshift-zonalautoshiftconfiguration-resourceidentifier
	//
	ResourceIdentifier *string `field:"required" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// A practice run configuration for a resource includes the Amazon CloudWatch alarms that you've specified for a practice run, as well as any blocked dates or blocked windows for the practice run.
	//
	// When a resource has a practice run configuration, ARC shifts traffic for the resource weekly for practice runs.
	//
	// Practice runs are required for zonal autoshift. The zonal shifts that ARC starts for practice runs help you to ensure that shifting away traffic from an Availability Zone during an autoshift is safe for your application.
	//
	// You can update or delete a practice run configuration. Before you delete a practice run configuration, you must disable zonal autoshift for the resource. A practice run configuration is required when zonal autoshift is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arczonalshift-zonalautoshiftconfiguration.html#cfn-arczonalshift-zonalautoshiftconfiguration-practicerunconfiguration
	//
	PracticeRunConfiguration interface{} `field:"optional" json:"practiceRunConfiguration" yaml:"practiceRunConfiguration"`
	// When zonal autoshift is `ENABLED` , you authorize AWS to shift away resource traffic for an application from an Availability Zone during events, on your behalf, to help reduce time to recovery.
	//
	// Traffic is also shifted away for the required weekly practice runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arczonalshift-zonalautoshiftconfiguration.html#cfn-arczonalshift-zonalautoshiftconfiguration-zonalautoshiftstatus
	//
	ZonalAutoshiftStatus *string `field:"optional" json:"zonalAutoshiftStatus" yaml:"zonalAutoshiftStatus"`
}

