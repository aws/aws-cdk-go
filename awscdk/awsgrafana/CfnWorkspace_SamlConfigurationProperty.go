package awsgrafana


// A structure containing information about how this workspace works with SAML.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   samlConfigurationProperty := &SamlConfigurationProperty{
//   	IdpMetadata: &IdpMetadataProperty{
//   		Url: jsii.String("url"),
//   		Xml: jsii.String("xml"),
//   	},
//
//   	// the properties below are optional
//   	AllowedOrganizations: []*string{
//   		jsii.String("allowedOrganizations"),
//   	},
//   	AssertionAttributes: &AssertionAttributesProperty{
//   		Email: jsii.String("email"),
//   		Groups: jsii.String("groups"),
//   		Login: jsii.String("login"),
//   		Name: jsii.String("name"),
//   		Org: jsii.String("org"),
//   		Role: jsii.String("role"),
//   	},
//   	LoginValidityDuration: jsii.Number(123),
//   	RoleValues: &RoleValuesProperty{
//   		Admin: []*string{
//   			jsii.String("admin"),
//   		},
//   		Editor: []*string{
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

