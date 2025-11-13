package interfacesawsec2


// A reference to a VPCEndpointServicePermissions resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCEndpointServicePermissionsReference := &VPCEndpointServicePermissionsReference{
//   	ServiceId: jsii.String("serviceId"),
//   }
//
type VPCEndpointServicePermissionsReference struct {
	// The ServiceId of the VPCEndpointServicePermissions resource.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
}

