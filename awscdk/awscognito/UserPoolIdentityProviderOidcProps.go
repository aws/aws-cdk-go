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
//   userPoolIdentityProviderOidcProps := &UserPoolIdentityProviderOidcProps{
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	IssuerUrl: jsii.String("issuerUrl"),
//   	UserPool: userPool,
//
//   	// the properties below are optional
//   	AttributeMapping: &AttributeMapping{
//   		Address: providerAttribute,
//   		Birthdate: providerAttribute,
//   		Custom: map[string]*providerAttribute{
//   			"customKey": providerAttribute,
//   		},
//   		Email: providerAttribute,
//   		FamilyName: providerAttribute,
//   		Fullname: providerAttribute,
//   		Gender: providerAttribute,
//   		GivenName: providerAttribute,
//   		LastUpdateTime: providerAttribute,
//   		Locale: providerAttribute,
//   		MiddleName: providerAttribute,
//   		Nickname: providerAttribute,
//   		PhoneNumber: providerAttribute,
//   		PreferredUsername: providerAttribute,
//   		ProfilePage: providerAttribute,
//   		ProfilePicture: providerAttribute,
//   		Timezone: providerAttribute,
//   		Website: providerAttribute,
//   	},
//   	AttributeRequestMethod: awscdk.Aws_cognito.OidcAttributeRequestMethod_GET,
//   	Endpoints: &OidcEndpoints{
//   		Authorization: jsii.String("authorization"),
//   		JwksUri: jsii.String("jwksUri"),
//   		Token: jsii.String("token"),
//   		UserInfo: jsii.String("userInfo"),
//   	},
//   	Identifiers: []*string{
//   		jsii.String("identifiers"),
//   	},
//   	Name: jsii.String("name"),
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   }
//
type UserPoolIdentityProviderOidcProps struct {
	// The user pool to which this construct provides identities.
	UserPool IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	// Default: - no attribute mapping.
	//
	AttributeMapping *AttributeMapping `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
	// The client id.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Issuer URL.
	IssuerUrl *string `field:"required" json:"issuerUrl" yaml:"issuerUrl"`
	// The method to use to request attributes.
	// Default: OidcAttributeRequestMethod.GET
	//
	AttributeRequestMethod OidcAttributeRequestMethod `field:"optional" json:"attributeRequestMethod" yaml:"attributeRequestMethod"`
	// OpenID connect endpoints.
	// Default: - auto discovered with issuer URL.
	//
	Endpoints *OidcEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
	// Identifiers.
	//
	// Identifiers can be used to redirect users to the correct IdP in multitenant apps.
	// Default: - no identifiers used.
	//
	Identifiers *[]*string `field:"optional" json:"identifiers" yaml:"identifiers"`
	// The name of the provider.
	// Default: - the unique ID of the construct.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The OAuth 2.0 scopes that you will request from OpenID Connect. Scopes are groups of OpenID Connect user attributes to exchange with your app.
	// Default: ['openid'].
	//
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

