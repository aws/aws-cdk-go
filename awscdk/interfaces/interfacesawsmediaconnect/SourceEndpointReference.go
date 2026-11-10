package interfacesawsmediaconnect


// A reference to a SourceEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceEndpointReference := &SourceEndpointReference{
//   	SourceEndpointArn: jsii.String("sourceEndpointArn"),
//   }
//
type SourceEndpointReference struct {
	// The Arn of the SourceEndpoint resource.
	SourceEndpointArn *string `field:"required" json:"sourceEndpointArn" yaml:"sourceEndpointArn"`
}

