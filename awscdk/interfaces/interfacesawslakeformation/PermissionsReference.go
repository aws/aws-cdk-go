package interfacesawslakeformation


// A reference to a Permissions resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   permissionsReference := &PermissionsReference{
//   	PermissionsId: jsii.String("permissionsId"),
//   }
//
type PermissionsReference struct {
	// The Id of the Permissions resource.
	PermissionsId *string `field:"required" json:"permissionsId" yaml:"permissionsId"`
}

