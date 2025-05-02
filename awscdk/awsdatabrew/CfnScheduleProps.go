package awsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSchedule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScheduleProps := &CfnScheduleProps{
//   	CronExpression: jsii.String("cronExpression"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	JobNames: []*string{
//   		jsii.String("jobNames"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-schedule.html
//
type CfnScheduleProps struct {
	// The dates and times when the job is to run.
	//
	// For more information, see [Working with cron expressions for recipe jobs](https://docs.aws.amazon.com/databrew/latest/dg/jobs.recipe.html#jobs.cron) in the *AWS Glue DataBrew Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-schedule.html#cfn-databrew-schedule-cronexpression
	//
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// The name of the schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-schedule.html#cfn-databrew-schedule-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of jobs to be run, according to the schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-schedule.html#cfn-databrew-schedule-jobnames
	//
	JobNames *[]*string `field:"optional" json:"jobNames" yaml:"jobNames"`
	// Metadata tags that have been applied to the schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-schedule.html#cfn-databrew-schedule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

