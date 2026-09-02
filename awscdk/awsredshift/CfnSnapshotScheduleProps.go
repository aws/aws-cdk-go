package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSnapshotSchedule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSnapshotScheduleProps := &CfnSnapshotScheduleProps{
//   	ScheduleDefinitions: []*string{
//   		jsii.String("scheduleDefinitions"),
//   	},
//   	ScheduleIdentifier: jsii.String("scheduleIdentifier"),
//
//   	// the properties below are optional
//   	ScheduleDescription: jsii.String("scheduleDescription"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-snapshotschedule.html
//
type CfnSnapshotScheduleProps struct {
	// The definition of the snapshot schedule.
	//
	// The definition is made up of schedule expressions, for example "cron(30 12 *)" or "rate(12 hours)".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-snapshotschedule.html#cfn-redshift-snapshotschedule-scheduledefinitions
	//
	ScheduleDefinitions *[]*string `field:"required" json:"scheduleDefinitions" yaml:"scheduleDefinitions"`
	// A unique identifier for the snapshot schedule.
	//
	// Only alphanumeric characters are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-snapshotschedule.html#cfn-redshift-snapshotschedule-scheduleidentifier
	//
	ScheduleIdentifier *string `field:"required" json:"scheduleIdentifier" yaml:"scheduleIdentifier"`
	// The description of the snapshot schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-snapshotschedule.html#cfn-redshift-snapshotschedule-scheduledescription
	//
	ScheduleDescription *string `field:"optional" json:"scheduleDescription" yaml:"scheduleDescription"`
	// An optional set of tags for the snapshot schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-snapshotschedule.html#cfn-redshift-snapshotschedule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

