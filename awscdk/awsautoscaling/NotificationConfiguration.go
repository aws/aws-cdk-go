package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// AutoScalingGroup fleet change notifications configurations.
//
// You can configure AutoScaling to send an SNS notification whenever your Auto Scaling group scales.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var scalingEvents scalingEvents
//   var topic topic
//
//   notificationConfiguration := &NotificationConfiguration{
//   	Topic: topic,
//
//   	// the properties below are optional
//   	ScalingEvents: scalingEvents,
//   }
//
type NotificationConfiguration struct {
	// SNS topic to send notifications about fleet scaling events.
	Topic awssns.ITopic `field:"required" json:"topic" yaml:"topic"`
	// Which fleet scaling events triggers a notification.
	// Default: ScalingEvents.ALL
	//
	ScalingEvents ScalingEvents `field:"optional" json:"scalingEvents" yaml:"scalingEvents"`
}

