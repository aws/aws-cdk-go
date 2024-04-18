package awscognito


// Properties to initialize UserPoolIdentityProviderSaml.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   // specify the metadata as a file content
//   // specify the metadata as a file content
//   cognito.NewUserPoolIdentityProviderSaml(this, jsii.String("userpoolIdpFile"), &UserPoolIdentityProviderSamlProps{
//   	UserPool: userpool,
//   	Metadata: cognito.UserPoolIdentityProviderSamlMetadata_File(jsii.String("my-file-contents")),
//   	// Whether to require encrypted SAML assertions from IdP
//   	EncryptedResponses: jsii.Boolean(true),
//   	// The signing algorithm for the SAML requests
//   	RequestSigningAlgorithm: cognito.SigningAlgorithm_RSA_SHA256,
//   	// Enable IdP initiated SAML auth flow
//   	IdpInitiated: jsii.Boolean(true),
//   })
//
//   // specify the metadata as a URL
//   // specify the metadata as a URL
//   cognito.NewUserPoolIdentityProviderSaml(this, jsii.String("userpoolidpUrl"), &UserPoolIdentityProviderSamlProps{
//   	UserPool: userpool,
//   	Metadata: cognito.UserPoolIdentityProviderSamlMetadata_Url(jsii.String("https://my-metadata-url.com")),
//   })
//
type UserPoolIdentityProviderSamlProps struct {
	// The user pool to which this construct provides identities.
	UserPool IUserPool `field:"required" json:"userPool" yaml:"userPool"`
	// Mapping attributes from the identity provider to standard and custom attributes of the user pool.
	// Default: - no attribute mapping.
	//
	AttributeMapping *AttributeMapping `field:"optional" json:"attributeMapping" yaml:"attributeMapping"`
	// The SAML metadata.
	Metadata UserPoolIdentityProviderSamlMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// Whether to require encrypted SAML assertions from IdP.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-SAML-signing-encryption.html#cognito-user-pools-SAML-encryption
	//
	// Default: false.
	//
	EncryptedResponses *bool `field:"optional" json:"encryptedResponses" yaml:"encryptedResponses"`
	// Identifiers.
	//
	// Identifiers can be used to redirect users to the correct IdP in multitenant apps.
	// Default: - no identifiers used.
	//
	Identifiers *[]*string `field:"optional" json:"identifiers" yaml:"identifiers"`
	// Whether to enable IdP-initiated SAML auth flows.
	// Default: false.
	//
	IdpInitiated *bool `field:"optional" json:"idpInitiated" yaml:"idpInitiated"`
	// Whether to enable the "Sign-out flow" feature.
	// Default: - false.
	//
	IdpSignout *bool `field:"optional" json:"idpSignout" yaml:"idpSignout"`
	// The name of the provider.
	//
	// Must be between 3 and 32 characters.
	// Default: - the unique ID of the construct.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The signing algorithm for SAML requests.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-SAML-signing-encryption.html#cognito-user-pools-SAML-signing
	//
	// Default: - don't sign requests.
	//
	RequestSigningAlgorithm SigningAlgorithm `field:"optional" json:"requestSigningAlgorithm" yaml:"requestSigningAlgorithm"`
}

