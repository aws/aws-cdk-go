package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnInstanceEventWindow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceEventWindowProps := &CfnInstanceEventWindowProps{
//   	CronExpression: jsii.String("cronExpression"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeRanges: []interface{}{
//   		&TimeRangeProperty{
//   			EndHour: jsii.Number(123),
//   			EndWeekDay: jsii.String("endWeekDay"),
//   			StartHour: jsii.Number(123),
//   			StartWeekDay: jsii.String("startWeekDay"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-instanceeventwindow.html
//
type CfnInstanceEventWindowProps struct {
	// The cron expression defined for the event window.
	//
	// Exactly one of TimeRanges or CronExpression must be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-instanceeventwindow.html#cfn-ec2-instanceeventwindow-cronexpression
	//
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// The name of the event window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-instanceeventwindow.html#cfn-ec2-instanceeventwindow-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags applied to the event window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-instanceeventwindow.html#cfn-ec2-instanceeventwindow-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The time ranges of the event window.
	//
	// Exactly one of TimeRanges or CronExpression must be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-instanceeventwindow.html#cfn-ec2-instanceeventwindow-timeranges
	//
	TimeRanges interface{} `field:"optional" json:"timeRanges" yaml:"timeRanges"`
}

