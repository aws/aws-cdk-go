package awsgrafana


// A structure containing information about how this workspace works with SAML.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   samlConfigurationProperty := &samlConfigurationProperty{
//   	idpMetadata: &idpMetadataProperty{
//   		url: jsii.String("url"),
//   		xml: jsii.String("xml"),
//   	},
//
//   	// the properties below are optional
//   	allowedOrganizations: []*string{
//   		jsii.String("allowedOrganizations"),
//   	},
//   	assertionAttributes: &assertionAttributesProperty{
//   		email: jsii.String("email"),
//   		groups: jsii.String("groups"),
//   		login: jsii.String("login"),
//   		name: jsii.String("name"),
//   		org: jsii.String("org"),
//   		role: jsii.String("role"),
//   	},
//   	loginValidityDuration: jsii.Number(123),
//   	roleValues: &roleValuesProperty{
//   		admin: []*string{
//   			jsii.String("admin"),
//   		},
//   		editor: []*string{
//   			jsii.String("editor"),
//   		},
//   	},
//   }
//
type CfnWorkspace_SamlConfigurationProperty struct {
	// A structure containing the identity provider (IdP) metadata used to integrate the identity provider with this workspace.
	IdpMetadata interface{} `field:"required" json:"idpMetadata" yaml:"idpMetadata"`
	// Lists which organizations defined in the SAML assertion are allowed to use the Amazon Managed Grafana workspace.
	//
	// If this is empty, all organizations in the assertion attribute have access.
	AllowedOrganizations *[]*string `field:"optional" json:"allowedOrganizations" yaml:"allowedOrganizations"`
	// A structure that defines which attributes in the SAML assertion are to be used to define information about the users authenticated by that IdP to use the workspace.
	AssertionAttributes interface{} `field:"optional" json:"assertionAttributes" yaml:"assertionAttributes"`
	// How long a sign-on session by a SAML user is valid, before the user has to sign on again.
	LoginValidityDuration *float64 `field:"optional" json:"loginValidityDuration" yaml:"loginValidityDuration"`
	// A structure containing arrays that map group names in the SAML assertion to the Grafana `Admin` and `Editor` roles in the workspace.
	RoleValues interface{} `field:"optional" json:"roleValues" yaml:"roleValues"`
}

