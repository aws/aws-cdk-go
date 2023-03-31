package awscognito


// Properties for defining a `CfnUserPoolIdentityProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributeMapping interface{}
//   var providerDetails interface{}
//
//   cfnUserPoolIdentityProviderProps := &cfnUserPoolIdentityProviderProps{
//   	providerName: jsii.String("providerName"),
//   	providerType: jsii.String("providerType"),
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	attributeMapping: attributeMapping,
//   	idpIdentifiers: []*string{
//   		jsii.String("idpIdentifiers"),
//   	},
//   	providerDetails: providerDetails,
//   }
//
type CfnUserPoolIdentityProviderProps struct {
	// The IdP name.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// The IdP type.
	ProviderType *string `field:"required" json:"providerType" yaml:"providerType"`
	// The user pool ID.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// A mapping of IdP attributes to standard and custom user pool attributes.
	AttributeMapping interface{} `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
	// A list of IdP identifiers.
	IdpIdentifiers *[]*string `field:"optional" json:"idpIdentifiers" yaml:"idpIdentifiers"`
	// The IdP details. The following list describes the provider detail keys for each IdP type.
	//
	// - For Google and Login with Amazon:
	//
	// - client_id
	// - client_secret
	// - authorize_scopes
	// - For Facebook:
	//
	// - client_id
	// - client_secret
	// - authorize_scopes
	// - api_version
	// - For Sign in with Apple:
	//
	// - client_id
	// - team_id
	// - key_id
	// - private_key
	// - authorize_scopes
	// - For OpenID Connect (OIDC) providers:
	//
	// - client_id
	// - client_secret
	// - attributes_request_method
	// - oidc_issuer
	// - authorize_scopes
	// - The following keys are only present if Amazon Cognito didn't discover them at the `oidc_issuer` URL.
	//
	// - authorize_url
	// - token_url
	// - attributes_url
	// - jwks_uri
	// - Amazon Cognito sets the value of the following keys automatically. They are read-only.
	//
	// - attributes_url_add_attributes
	// - For SAML providers:
	//
	// - MetadataFile or MetadataURL
	// - IDPSignout *optional*.
	ProviderDetails interface{} `field:"optional" json:"providerDetails" yaml:"providerDetails"`
}

