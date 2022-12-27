package awscognito


// Properties to initialize UserPoolIdentityProviderSaml.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var providerAttribute providerAttribute
//   var userPool userPool
//   var userPoolIdentityProviderSamlMetadata userPoolIdentityProviderSamlMetadata
//
//   userPoolIdentityProviderSamlProps := &userPoolIdentityProviderSamlProps{
//   	metadata: userPoolIdentityProviderSamlMetadata,
//   	userPool: userPool,
//
//   	// the properties below are optional
//   	attributeMapping: &attributeMapping{
//   		address: providerAttribute,
//   		birthdate: providerAttribute,
//   		custom: map[string]*providerAttribute{
//   			"customKey": providerAttribute,
//   		},
//   		email: providerAttribute,
//   		familyName: providerAttribute,
//   		fullname: providerAttribute,
//   		gender: providerAttribute,
//   		givenName: providerAttribute,
//   		lastUpdateTime: providerAttribute,
//   		locale: providerAttribute,
//   		middleName: providerAttribute,
//   		nickname: providerAttribute,
//   		phoneNumber: providerAttribute,
//   		preferredUsername: providerAttribute,
//   		profilePage: providerAttribute,
//   		profilePicture: providerAttribute,
//   		timezone: providerAttribute,
//   		website: providerAttribute,
//   	},
//   	identifiers: []*string{
//   		jsii.String("identifiers"),
//   	},
//   	idpSignout: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   }
//
type UserPoolIdentityProviderSamlProps struct {
	// The user pool to which this construct provides identities.
	UserPool IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	AttributeMapping *AttributeMapping `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
	// The SAML metadata.
	Metadata UserPoolIdentityProviderSamlMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// Identifiers.
	//
	// Identifiers can be used to redirect users to the correct IdP in multitenant apps.
	Identifiers *[]*string `field:"optional" json:"identifiers" yaml:"identifiers"`
	// Whether to enable the "Sign-out flow" feature.
	IdpSignout *bool `field:"optional" json:"idpSignout" yaml:"idpSignout"`
	// The name of the provider.
	//
	// Must be between 3 and 32 characters.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

