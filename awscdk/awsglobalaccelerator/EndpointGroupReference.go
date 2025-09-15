package awsglobalaccelerator


// A reference to a EndpointGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointGroupReference := &EndpointGroupReference{
//   	EndpointGroupArn: jsii.String("endpointGroupArn"),
//   }
//
type EndpointGroupReference struct {
	// The EndpointGroupArn of the EndpointGroup resource.
	EndpointGroupArn *string `field:"required" json:"endpointGroupArn" yaml:"endpointGroupArn"`
}

