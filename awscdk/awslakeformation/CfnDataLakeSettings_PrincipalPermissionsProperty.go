package awslakeformation


// Permissions granted to a principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   principalPermissionsProperty := &PrincipalPermissionsProperty{
//   	Permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	Principal: &DataLakePrincipalProperty{
//   		DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   }
//
type CfnDataLakeSettings_PrincipalPermissionsProperty struct {
	// The permissions that are granted to the principal.
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// The principal who is granted permissions.
	Principal interface{} `field:"required" json:"principal" yaml:"principal"`
}

