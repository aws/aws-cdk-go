package previewawsec2events

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.ec2@EC2FastLaunchStateChangeNotification event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2FastLaunchStateChangeNotificationProps := &EC2FastLaunchStateChangeNotificationProps{
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
//   	ImageId: []*string{
//   		jsii.String("imageId"),
//   	},
//   	ResourceType: []*string{
//   		jsii.String("resourceType"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   	StateTransitionReason: []*string{
//   		jsii.String("stateTransitionReason"),
//   	},
//   }
//
// Experimental.
type EC2FastLaunchStateChangeNotification_EC2FastLaunchStateChangeNotificationProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// imageId property.
	//
	// Specify an array of string values to match this event if the actual value of imageId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageId *[]*string `field:"optional" json:"imageId" yaml:"imageId"`
	// resourceType property.
	//
	// Specify an array of string values to match this event if the actual value of resourceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceType *[]*string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
	// stateTransitionReason property.
	//
	// Specify an array of string values to match this event if the actual value of stateTransitionReason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StateTransitionReason *[]*string `field:"optional" json:"stateTransitionReason" yaml:"stateTransitionReason"`
}

