package awscognito


// Properties for defining a `CfnIdentityPool`.
//
// Example:
//   import cognito "github.com/aws/aws-cdk-go/awscdk"
//
//   var myProvider openIdConnectProvider
//
//   cognito.NewCfnIdentityPool(this, jsii.String("IdentityPool"), &cfnIdentityPoolProps{
//   	openIdConnectProviderArns: []*string{
//   		myProvider.openIdConnectProviderArn,
//   	},
//   	// And the other properties for your identity pool
//   	allowUnauthenticatedIdentities: jsii.Boolean(false),
//   })
//
type CfnIdentityPoolProps struct {
	// Specifies whether the identity pool supports unauthenticated logins.
	AllowUnauthenticatedIdentities interface{} `field:"required" json:"allowUnauthenticatedIdentities" yaml:"allowUnauthenticatedIdentities"`
	// Enables the Basic (Classic) authentication flow.
	AllowClassicFlow interface{} `field:"optional" json:"allowClassicFlow" yaml:"allowClassicFlow"`
	// The events to configure.
	CognitoEvents interface{} `field:"optional" json:"cognitoEvents" yaml:"cognitoEvents"`
	// The Amazon Cognito user pools and their client IDs.
	CognitoIdentityProviders interface{} `field:"optional" json:"cognitoIdentityProviders" yaml:"cognitoIdentityProviders"`
	// Configuration options for configuring Amazon Cognito streams.
	CognitoStreams interface{} `field:"optional" json:"cognitoStreams" yaml:"cognitoStreams"`
	// The "domain" Amazon Cognito uses when referencing your users.
	//
	// This name acts as a placeholder that allows your backend and the Amazon Cognito service to communicate about the developer provider. For the `DeveloperProviderName` , you can use letters and periods (.), underscores (_), and dashes (-).
	//
	// *Minimum length* : 1
	//
	// *Maximum length* : 100.
	DeveloperProviderName *string `field:"optional" json:"developerProviderName" yaml:"developerProviderName"`
	// The name of your Amazon Cognito identity pool.
	//
	// *Minimum length* : 1
	//
	// *Maximum length* : 128
	//
	// *Pattern* : `[\w\s+=,.@-]+`
	IdentityPoolName *string `field:"optional" json:"identityPoolName" yaml:"identityPoolName"`
	// The Amazon Resource Names (ARNs) of the OpenID connect providers.
	OpenIdConnectProviderArns *[]*string `field:"optional" json:"openIdConnectProviderArns" yaml:"openIdConnectProviderArns"`
	// The configuration options to be applied to the identity pool.
	PushSync interface{} `field:"optional" json:"pushSync" yaml:"pushSync"`
	// The Amazon Resource Names (ARNs) of the Security Assertion Markup Language (SAML) providers.
	SamlProviderArns *[]*string `field:"optional" json:"samlProviderArns" yaml:"samlProviderArns"`
	// Key-value pairs that map provider names to provider app IDs.
	SupportedLoginProviders interface{} `field:"optional" json:"supportedLoginProviders" yaml:"supportedLoginProviders"`
}

