package awsevents


// The event buses the endpoint is associated with.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointEventBusProperty := &EndpointEventBusProperty{
//   	EventBusArn: jsii.String("eventBusArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-endpoint-endpointeventbus.html
//
type CfnEndpoint_EndpointEventBusProperty struct {
	// The ARN of the event bus the endpoint is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-endpoint-endpointeventbus.html#cfn-events-endpoint-endpointeventbus-eventbusarn
	//
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
}

