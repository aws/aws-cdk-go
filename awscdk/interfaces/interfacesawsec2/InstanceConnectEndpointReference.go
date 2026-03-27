package interfacesawsec2


// A reference to a InstanceConnectEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceConnectEndpointReference := &InstanceConnectEndpointReference{
//   	InstanceConnectEndpointArn: jsii.String("instanceConnectEndpointArn"),
//   	InstanceConnectEndpointId: jsii.String("instanceConnectEndpointId"),
//   }
//
type InstanceConnectEndpointReference struct {
	// The ARN of the InstanceConnectEndpoint resource.
	InstanceConnectEndpointArn *string `field:"required" json:"instanceConnectEndpointArn" yaml:"instanceConnectEndpointArn"`
	// The Id of the InstanceConnectEndpoint resource.
	InstanceConnectEndpointId *string `field:"required" json:"instanceConnectEndpointId" yaml:"instanceConnectEndpointId"`
}

