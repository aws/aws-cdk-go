package awsdms


// A reference to a Endpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointReference := &EndpointReference{
//   	EndpointId: jsii.String("endpointId"),
//   }
//
type EndpointReference struct {
	// The Id of the Endpoint resource.
	EndpointId *string `field:"required" json:"endpointId" yaml:"endpointId"`
}

