package interfacesawslambda


// A reference to a WebFunctionEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webFunctionEndpointReference := &WebFunctionEndpointReference{
//   	EndpointArn: jsii.String("endpointArn"),
//   }
//
type WebFunctionEndpointReference struct {
	// The EndpointArn of the WebFunctionEndpoint resource.
	EndpointArn *string `field:"required" json:"endpointArn" yaml:"endpointArn"`
}

