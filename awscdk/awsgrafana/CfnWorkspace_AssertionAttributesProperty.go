package awsgrafana


// A structure that defines which attributes in the IdP assertion are to be used to define information about the users authenticated by the IdP to use the workspace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assertionAttributesProperty := &AssertionAttributesProperty{
//   	Email: jsii.String("email"),
//   	Groups: jsii.String("groups"),
//   	Login: jsii.String("login"),
//   	Name: jsii.String("name"),
//   	Org: jsii.String("org"),
//   	Role: jsii.String("role"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-assertionattributes.html
//
type CfnWorkspace_AssertionAttributesProperty struct {
	// The name of the attribute within the SAML assertion to use as the email names for SAML users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-assertionattributes.html#cfn-grafana-workspace-assertionattributes-email
	//
	Email *string `field:"optional" json:"email" yaml:"email"`
	// The name of the attribute within the SAML assertion to use as the user full "friendly" names for user groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-assertionattributes.html#cfn-grafana-workspace-assertionattributes-groups
	//
	Groups *string `field:"optional" json:"groups" yaml:"groups"`
	// The name of the attribute within the SAML assertion to use as the login names for SAML users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-assertionattributes.html#cfn-grafana-workspace-assertionattributes-login
	//
	Login *string `field:"optional" json:"login" yaml:"login"`
	// The name of the attribute within the SAML assertion to use as the user full "friendly" names for SAML users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-assertionattributes.html#cfn-grafana-workspace-assertionattributes-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of the attribute within the SAML assertion to use as the user full "friendly" names for the users' organizations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-assertionattributes.html#cfn-grafana-workspace-assertionattributes-org
	//
	Org *string `field:"optional" json:"org" yaml:"org"`
	// The name of the attribute within the SAML assertion to use as the user roles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-assertionattributes.html#cfn-grafana-workspace-assertionattributes-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
}

