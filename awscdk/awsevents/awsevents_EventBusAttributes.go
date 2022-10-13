package awsevents


// Interface with properties necessary to import a reusable EventBus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBusAttributes := &eventBusAttributes{
//   	eventBusArn: jsii.String("eventBusArn"),
//   	eventBusName: jsii.String("eventBusName"),
//   	eventBusPolicy: jsii.String("eventBusPolicy"),
//
//   	// the properties below are optional
//   	eventSourceName: jsii.String("eventSourceName"),
//   }
//
type EventBusAttributes struct {
	// The ARN of this event bus resource.
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
	// The physical ID of this event bus resource.
	EventBusName *string `field:"required" json:"eventBusName" yaml:"eventBusName"`
	// The JSON policy of this event bus resource.
	EventBusPolicy *string `field:"required" json:"eventBusPolicy" yaml:"eventBusPolicy"`
	// The partner event source to associate with this event bus resource.
	EventSourceName *string `field:"optional" json:"eventSourceName" yaml:"eventSourceName"`
}

