package awscognito


// Properties to initialize UserPoolIdentityProviderOidc.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var providerAttribute providerAttribute
//   var userPool userPool
//
//   userPoolIdentityProviderOidcProps := &userPoolIdentityProviderOidcProps{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//   	issuerUrl: jsii.String("issuerUrl"),
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
//   	attributeRequestMethod: awscdk.Aws_cognito.oidcAttributeRequestMethod_GET,
//   	endpoints: &oidcEndpoints{
//   		authorization: jsii.String("authorization"),
//   		jwksUri: jsii.String("jwksUri"),
//   		token: jsii.String("token"),
//   		userInfo: jsii.String("userInfo"),
//   	},
//   	identifiers: []*string{
//   		jsii.String("identifiers"),
//   	},
//   	name: jsii.String("name"),
//   	scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   }
//
type UserPoolIdentityProviderOidcProps struct {
	// The user pool to which this construct provides identities.
	UserPool IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	AttributeMapping *AttributeMapping `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
	// The client id.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Issuer URL.
	IssuerUrl *string `field:"required" json:"issuerUrl" yaml:"issuerUrl"`
	// The method to use to request attributes.
	AttributeRequestMethod OidcAttributeRequestMethod `field:"optional" json:"attributeRequestMethod" yaml:"attributeRequestMethod"`
	// OpenID connect endpoints.
	Endpoints *OidcEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
	// Identifiers.
	//
	// Identifiers can be used to redirect users to the correct IdP in multitenant apps.
	Identifiers *[]*string `field:"optional" json:"identifiers" yaml:"identifiers"`
	// The name of the provider.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The OAuth 2.0 scopes that you will request from OpenID Connect. Scopes are groups of OpenID Connect user attributes to exchange with your app.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

