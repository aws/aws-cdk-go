package awsscheduler

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnScheduleGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScheduleGroupProps := &cfnScheduleGroupProps{
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnScheduleGroupProps struct {
	// `AWS::Scheduler::ScheduleGroup.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::Scheduler::ScheduleGroup.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

