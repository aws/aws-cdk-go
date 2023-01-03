package awsgrafana


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
	// `CfnWorkspace.SamlConfigurationProperty.IdpMetadata`.
	IdpMetadata interface{} `field:"required" json:"idpMetadata" yaml:"idpMetadata"`
	// `CfnWorkspace.SamlConfigurationProperty.AllowedOrganizations`.
	AllowedOrganizations *[]*string `field:"optional" json:"allowedOrganizations" yaml:"allowedOrganizations"`
	// `CfnWorkspace.SamlConfigurationProperty.AssertionAttributes`.
	AssertionAttributes interface{} `field:"optional" json:"assertionAttributes" yaml:"assertionAttributes"`
	// `CfnWorkspace.SamlConfigurationProperty.LoginValidityDuration`.
	LoginValidityDuration *float64 `field:"optional" json:"loginValidityDuration" yaml:"loginValidityDuration"`
	// `CfnWorkspace.SamlConfigurationProperty.RoleValues`.
	RoleValues interface{} `field:"optional" json:"roleValues" yaml:"roleValues"`
}

