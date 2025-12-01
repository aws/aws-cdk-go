package interfacesawssagemaker


// A reference to a EndpointConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointConfigReference := &EndpointConfigReference{
//   	EndpointConfigId: jsii.String("endpointConfigId"),
//   	EndpointConfigName: jsii.String("endpointConfigName"),
//   }
//
type EndpointConfigReference struct {
	// The Id of the EndpointConfig resource.
	EndpointConfigId *string `field:"required" json:"endpointConfigId" yaml:"endpointConfigId"`
	// The EndpointConfigName of the EndpointConfig resource.
	EndpointConfigName *string `field:"required" json:"endpointConfigName" yaml:"endpointConfigName"`
}

