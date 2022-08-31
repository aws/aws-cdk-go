package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnHoursOfOperation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHoursOfOperationProps := &cfnHoursOfOperationProps{
//   	config: []interface{}{
//   		&hoursOfOperationConfigProperty{
//   			day: jsii.String("day"),
//   			endTime: &hoursOfOperationTimeSliceProperty{
//   				hours: jsii.Number(123),
//   				minutes: jsii.Number(123),
//   			},
//   			startTime: &hoursOfOperationTimeSliceProperty{
//   				hours: jsii.Number(123),
//   				minutes: jsii.Number(123),
//   			},
//   		},
//   	},
//   	instanceArn: jsii.String("instanceArn"),
//   	name: jsii.String("name"),
//   	timeZone: jsii.String("timeZone"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnHoursOfOperationProps struct {
	// Configuration information for the hours of operation.
	Config interface{} `field:"required" json:"config" yaml:"config"`
	// The Amazon Resource Name (ARN) for the instance.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name for the hours of operation.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The time zone for the hours of operation.
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// The description for the hours of operation.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

