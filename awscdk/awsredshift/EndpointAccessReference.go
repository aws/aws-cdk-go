package awsredshift


// A reference to a EndpointAccess resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointAccessReference := &EndpointAccessReference{
//   	EndpointName: jsii.String("endpointName"),
//   }
//
type EndpointAccessReference struct {
	// The EndpointName of the EndpointAccess resource.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
}

