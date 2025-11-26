package previewawsec2events

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Instance aws.ec2@EC2InstanceStateChangeNotification event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2InstanceStateChangeNotificationProps := &EC2InstanceStateChangeNotificationProps{
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	InstanceId: []*string{
//   		jsii.String("instanceId"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_EC2InstanceStateChangeNotification_EC2InstanceStateChangeNotificationProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// instance-id property.
	//
	// Specify an array of string values to match this event if the actual value of instance-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Instance reference.
	//
	// Experimental.
	InstanceId *[]*string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
}

