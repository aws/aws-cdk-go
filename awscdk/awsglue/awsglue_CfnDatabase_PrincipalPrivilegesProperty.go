package awsglue


// the permissions granted to a principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   principalPrivilegesProperty := &principalPrivilegesProperty{
//   	permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	principal: &dataLakePrincipalProperty{
//   		dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   }
//
type CfnDatabase_PrincipalPrivilegesProperty struct {
	// The permissions that are granted to the principal.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// The principal who is granted permissions.
	Principal interface{} `field:"optional" json:"principal" yaml:"principal"`
}

