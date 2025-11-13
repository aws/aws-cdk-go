package interfacesawsiot


// A reference to a RoleAlias resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   roleAliasReference := &RoleAliasReference{
//   	RoleAlias: jsii.String("roleAlias"),
//   	RoleAliasArn: jsii.String("roleAliasArn"),
//   }
//
type RoleAliasReference struct {
	// The RoleAlias of the RoleAlias resource.
	RoleAlias *string `field:"required" json:"roleAlias" yaml:"roleAlias"`
	// The ARN of the RoleAlias resource.
	RoleAliasArn *string `field:"required" json:"roleAliasArn" yaml:"roleAliasArn"`
}

