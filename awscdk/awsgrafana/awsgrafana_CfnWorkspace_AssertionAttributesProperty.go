package awsgrafana


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
	// `CfnWorkspace.AssertionAttributesProperty.Email`.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// `CfnWorkspace.AssertionAttributesProperty.Groups`.
	Groups *string `field:"optional" json:"groups" yaml:"groups"`
	// `CfnWorkspace.AssertionAttributesProperty.Login`.
	Login *string `field:"optional" json:"login" yaml:"login"`
	// `CfnWorkspace.AssertionAttributesProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnWorkspace.AssertionAttributesProperty.Org`.
	Org *string `field:"optional" json:"org" yaml:"org"`
	// `CfnWorkspace.AssertionAttributesProperty.Role`.
	Role *string `field:"optional" json:"role" yaml:"role"`
}

