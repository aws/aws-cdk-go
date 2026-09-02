package awssagemaker


// The configuration of an OIDC Identity Provider (IdP) private workforce.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   oidcConfigProperty := &OidcConfigProperty{
//   	AuthenticationRequestExtraParams: map[string]*string{
//   		"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   	},
//   	AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	Issuer: jsii.String("issuer"),
//   	JwksUri: jsii.String("jwksUri"),
//   	LogoutEndpoint: jsii.String("logoutEndpoint"),
//   	Scope: jsii.String("scope"),
//   	TokenEndpoint: jsii.String("tokenEndpoint"),
//   	UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-oidcconfig.html
//
type CfnWorkforcePropsMixin_OidcConfigProperty struct {
	// A string to string map of identifiers specific to the custom identity provider (IdP) being used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-oidcconfig.html#cfn-sagemaker-workforce-oidcconfig-authenticationrequestextraparams
	//
	AuthenticationRequestExtraParams interface{} `field:"optional" json:"authenticationRequestExtraParams" yaml:"authenticationRequestExtraParams"`
	// The OIDC IdP authorization endpoint used to configure your private workforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-oidcconfig.html#cfn-sagemaker-workforce-oidcconfig-authorizationendpoint
	//
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The OIDC IdP client ID used to configure your private workforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-oidcconfig.html#cfn-sagemaker-workforce-oidcconfig-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The OIDC IdP client secret used to configure your private workforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-oidcconfig.html#cfn-sagemaker-workforce-oidcconfig-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The OIDC IdP issuer used to configure your private workforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-oidcconfig.html#cfn-sagemaker-workforce-oidcconfig-issuer
	//
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// The OIDC IdP JSON Web Key Set (Jwks) URI used to configure your private workforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-oidcconfig.html#cfn-sagemaker-workforce-oidcconfig-jwksuri
	//
	JwksUri *string `field:"optional" json:"jwksUri" yaml:"jwksUri"`
	// The OIDC IdP logout endpoint used to configure your private workforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-oidcconfig.html#cfn-sagemaker-workforce-oidcconfig-logoutendpoint
	//
	LogoutEndpoint *string `field:"optional" json:"logoutEndpoint" yaml:"logoutEndpoint"`
	// An array of string identifiers used to refer to the specific pieces of user data or claims that the client application wants to access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-oidcconfig.html#cfn-sagemaker-workforce-oidcconfig-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The OIDC IdP token endpoint used to configure your private workforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-oidcconfig.html#cfn-sagemaker-workforce-oidcconfig-tokenendpoint
	//
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The OIDC IdP user info endpoint used to configure your private workforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-workforce-oidcconfig.html#cfn-sagemaker-workforce-oidcconfig-userinfoendpoint
	//
	UserInfoEndpoint *string `field:"optional" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
}

