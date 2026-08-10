package previewawsstoragegatewayevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.storagegateway@StorageGatewayFileUploadEvent event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   storageGatewayFileUploadEventProps := &StorageGatewayFileUploadEventProps{
//   	Completed: []*string{
//   		jsii.String("completed"),
//   	},
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
//   	EventType: []*string{
//   		jsii.String("eventType"),
//   	},
//   	NotificationId: []*string{
//   		jsii.String("notificationId"),
//   	},
//   	RequestReceived: []*string{
//   		jsii.String("requestReceived"),
//   	},
//   }
//
// Experimental.
type StorageGatewayFileUploadEvent_StorageGatewayFileUploadEventProps struct {
	// completed property.
	//
	// Specify an array of string values to match this event if the actual value of completed is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Completed *[]*string `field:"optional" json:"completed" yaml:"completed"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// event-type property.
	//
	// Specify an array of string values to match this event if the actual value of event-type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// notification-id property.
	//
	// Specify an array of string values to match this event if the actual value of notification-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NotificationId *[]*string `field:"optional" json:"notificationId" yaml:"notificationId"`
	// request-received property.
	//
	// Specify an array of string values to match this event if the actual value of request-received is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestReceived *[]*string `field:"optional" json:"requestReceived" yaml:"requestReceived"`
}

