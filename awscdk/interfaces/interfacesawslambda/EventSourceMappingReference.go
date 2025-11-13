package interfacesawslambda


// A reference to a EventSourceMapping resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventSourceMappingReference := &EventSourceMappingReference{
//   	EventSourceMappingArn: jsii.String("eventSourceMappingArn"),
//   	EventSourceMappingId: jsii.String("eventSourceMappingId"),
//   }
//
type EventSourceMappingReference struct {
	// The ARN of the EventSourceMapping resource.
	EventSourceMappingArn *string `field:"required" json:"eventSourceMappingArn" yaml:"eventSourceMappingArn"`
	// The Id of the EventSourceMapping resource.
	EventSourceMappingId *string `field:"required" json:"eventSourceMappingId" yaml:"eventSourceMappingId"`
}

