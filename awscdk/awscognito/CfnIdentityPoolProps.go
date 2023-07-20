package awscognito


// Properties for defining a `CfnIdentityPool`.
//
// Example:
//   import cognito "github.com/aws/aws-cdk-go/awscdk"
//
//   var myProvider openIdConnectProvider
//
//   cognito.NewCfnIdentityPool(this, jsii.String("IdentityPool"), &CfnIdentityPoolProps{
//   	OpenIdConnectProviderArns: []*string{
//   		myProvider.OpenIdConnectProviderArn,
//   	},
//   	// And the other properties for your identity pool
//   	AllowUnauthenticatedIdentities: jsii.Boolean(false),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html
//
type CfnIdentityPoolProps struct {
	// Specifies whether the identity pool supports unauthenticated logins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-allowunauthenticatedidentities
	//
	AllowUnauthenticatedIdentities interface{} `field:"required" json:"allowUnauthenticatedIdentities" yaml:"allowUnauthenticatedIdentities"`
	// Enables the Basic (Classic) authentication flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-allowclassicflow
	//
	AllowClassicFlow interface{} `field:"optional" json:"allowClassicFlow" yaml:"allowClassicFlow"`
	// The events to configure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoevents
	//
	CognitoEvents interface{} `field:"optional" json:"cognitoEvents" yaml:"cognitoEvents"`
	// The Amazon Cognito user pools and their client IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoidentityproviders
	//
	CognitoIdentityProviders interface{} `field:"optional" json:"cognitoIdentityProviders" yaml:"cognitoIdentityProviders"`
	// Configuration options for configuring Amazon Cognito streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitostreams
	//
	CognitoStreams interface{} `field:"optional" json:"cognitoStreams" yaml:"cognitoStreams"`
	// The "domain" Amazon Cognito uses when referencing your users.
	//
	// This name acts as a placeholder that allows your backend and the Amazon Cognito service to communicate about the developer provider. For the `DeveloperProviderName` , you can use letters and periods (.), underscores (_), and dashes (-).
	//
	// *Minimum length* : 1
	//
	// *Maximum length* : 100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-developerprovidername
	//
	DeveloperProviderName *string `field:"optional" json:"developerProviderName" yaml:"developerProviderName"`
	// The name of your Amazon Cognito identity pool.
	//
	// *Minimum length* : 1
	//
	// *Maximum length* : 128
	//
	// *Pattern* : `[\w\s+=,.@-]+`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-identitypoolname
	//
	IdentityPoolName *string `field:"optional" json:"identityPoolName" yaml:"identityPoolName"`
	// The Amazon Resource Names (ARNs) of the OpenID connect providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-openidconnectproviderarns
	//
	OpenIdConnectProviderArns *[]*string `field:"optional" json:"openIdConnectProviderArns" yaml:"openIdConnectProviderArns"`
	// The configuration options to be applied to the identity pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-pushsync
	//
	PushSync interface{} `field:"optional" json:"pushSync" yaml:"pushSync"`
	// The Amazon Resource Names (ARNs) of the Security Assertion Markup Language (SAML) providers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-samlproviderarns
	//
	SamlProviderArns *[]*string `field:"optional" json:"samlProviderArns" yaml:"samlProviderArns"`
	// Key-value pairs that map provider names to provider app IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-supportedloginproviders
	//
	SupportedLoginProviders interface{} `field:"optional" json:"supportedLoginProviders" yaml:"supportedLoginProviders"`
}

