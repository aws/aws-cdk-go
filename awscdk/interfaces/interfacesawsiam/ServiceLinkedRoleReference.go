package interfacesawsiam


// A reference to a ServiceLinkedRole resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceLinkedRoleReference := &ServiceLinkedRoleReference{
//   	RoleName: jsii.String("roleName"),
//   }
//
type ServiceLinkedRoleReference struct {
	// The RoleName of the ServiceLinkedRole resource.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
}

