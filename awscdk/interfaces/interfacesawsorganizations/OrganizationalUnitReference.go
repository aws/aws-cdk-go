package interfacesawsorganizations


// A reference to a OrganizationalUnit resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationalUnitReference := &OrganizationalUnitReference{
//   	OrganizationalUnitArn: jsii.String("organizationalUnitArn"),
//   	OrganizationalUnitId: jsii.String("organizationalUnitId"),
//   }
//
type OrganizationalUnitReference struct {
	// The ARN of the OrganizationalUnit resource.
	OrganizationalUnitArn *string `field:"required" json:"organizationalUnitArn" yaml:"organizationalUnitArn"`
	// The Id of the OrganizationalUnit resource.
	OrganizationalUnitId *string `field:"required" json:"organizationalUnitId" yaml:"organizationalUnitId"`
}

