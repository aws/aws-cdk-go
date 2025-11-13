package interfacesawslakeformation


// A reference to a PrincipalPermissions resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   principalPermissionsReference := &PrincipalPermissionsReference{
//   	PrincipalIdentifier: jsii.String("principalIdentifier"),
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   }
//
type PrincipalPermissionsReference struct {
	// The PrincipalIdentifier of the PrincipalPermissions resource.
	PrincipalIdentifier *string `field:"required" json:"principalIdentifier" yaml:"principalIdentifier"`
	// The ResourceIdentifier of the PrincipalPermissions resource.
	ResourceIdentifier *string `field:"required" json:"resourceIdentifier" yaml:"resourceIdentifier"`
}

