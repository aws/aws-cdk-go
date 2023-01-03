package awsgrafana


// A structure that defines which attributes in the IdP assertion are to be used to define information about the users authenticated by the IdP to use the workspace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assertionAttributesProperty := &assertionAttributesProperty{
//   	email: jsii.String("email"),
//   	groups: jsii.String("groups"),
//   	login: jsii.String("login"),
//   	name: jsii.String("name"),
//   	org: jsii.String("org"),
//   	role: jsii.String("role"),
//   }
//
type CfnWorkspace_AssertionAttributesProperty struct {
	// The name of the attribute within the SAML assertion to use as the email names for SAML users.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// The name of the attribute within the SAML assertion to use as the user full "friendly" names for user groups.
	Groups *string `field:"optional" json:"groups" yaml:"groups"`
	// The name of the attribute within the SAML assertion to use as the login names for SAML users.
	Login *string `field:"optional" json:"login" yaml:"login"`
	// The name of the attribute within the SAML assertion to use as the user full "friendly" names for SAML users.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of the attribute within the SAML assertion to use as the user full "friendly" names for the users' organizations.
	Org *string `field:"optional" json:"org" yaml:"org"`
	// The name of the attribute within the SAML assertion to use as the user roles.
	Role *string `field:"optional" json:"role" yaml:"role"`
}

