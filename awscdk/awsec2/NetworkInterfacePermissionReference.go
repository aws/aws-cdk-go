package awsec2


// A reference to a NetworkInterfacePermission resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfacePermissionReference := &NetworkInterfacePermissionReference{
//   	NetworkInterfacePermissionId: jsii.String("networkInterfacePermissionId"),
//   }
//
type NetworkInterfacePermissionReference struct {
	// The Id of the NetworkInterfacePermission resource.
	NetworkInterfacePermissionId *string `field:"required" json:"networkInterfacePermissionId" yaml:"networkInterfacePermissionId"`
}

