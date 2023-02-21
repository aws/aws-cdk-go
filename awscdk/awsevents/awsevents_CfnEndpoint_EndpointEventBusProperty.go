package awsevents


// The event buses the endpoint is associated with.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointEventBusProperty := &endpointEventBusProperty{
//   	eventBusArn: jsii.String("eventBusArn"),
//   }
//
type CfnEndpoint_EndpointEventBusProperty struct {
	// The ARN of the event bus the endpoint is associated with.
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
}

