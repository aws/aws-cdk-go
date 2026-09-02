package interfacesawssagemaker


// A reference to a EndpointConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointConfigReference := &EndpointConfigReference{
//   	EndpointConfigArn: jsii.String("endpointConfigArn"),
//   }
//
type EndpointConfigReference struct {
	// The EndpointConfigArn of the EndpointConfig resource.
	EndpointConfigArn *string `field:"required" json:"endpointConfigArn" yaml:"endpointConfigArn"`
}

