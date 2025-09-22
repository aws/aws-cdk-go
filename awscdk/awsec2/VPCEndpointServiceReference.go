package awsec2


// A reference to a VPCEndpointService resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCEndpointServiceReference := &VPCEndpointServiceReference{
//   	ServiceId: jsii.String("serviceId"),
//   }
//
type VPCEndpointServiceReference struct {
	// The ServiceId of the VPCEndpointService resource.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
}

