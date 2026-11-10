package interfacesawsmediaconnect


// A reference to a DestinationEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationEndpointReference := &DestinationEndpointReference{
//   	DestinationEndpointArn: jsii.String("destinationEndpointArn"),
//   }
//
type DestinationEndpointReference struct {
	// The Arn of the DestinationEndpoint resource.
	DestinationEndpointArn *string `field:"required" json:"destinationEndpointArn" yaml:"destinationEndpointArn"`
}

