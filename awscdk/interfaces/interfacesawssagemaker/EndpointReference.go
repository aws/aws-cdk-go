package interfacesawssagemaker


// A reference to a Endpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointReference := &EndpointReference{
//   	EndpointArn: jsii.String("endpointArn"),
//   }
//
type EndpointReference struct {
	// The EndpointArn of the Endpoint resource.
	EndpointArn *string `field:"required" json:"endpointArn" yaml:"endpointArn"`
}

