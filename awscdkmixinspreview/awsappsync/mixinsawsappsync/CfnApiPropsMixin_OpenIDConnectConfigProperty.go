package mixinsawsappsync


// Describes an OpenID Connect (OIDC) configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   openIDConnectConfigProperty := &OpenIDConnectConfigProperty{
//   	AuthTtl: jsii.Number(123),
//   	ClientId: jsii.String("clientId"),
//   	IatTtl: jsii.Number(123),
//   	Issuer: jsii.String("issuer"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-openidconnectconfig.html
//
type CfnApiPropsMixin_OpenIDConnectConfigProperty struct {
	// The number of milliseconds that a token is valid after being authenticated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-openidconnectconfig.html#cfn-appsync-api-openidconnectconfig-authttl
	//
	AuthTtl *float64 `field:"optional" json:"authTtl" yaml:"authTtl"`
	// The client identifier of the relying party at the OpenID identity provider.
	//
	// This identifier is typically obtained when the relying party is registered with the OpenID identity provider. You can specify a regular expression so that AWS AppSync can validate against multiple client identifiers at a time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-openidconnectconfig.html#cfn-appsync-api-openidconnectconfig-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The number of milliseconds that a token is valid after it's issued to a user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-openidconnectconfig.html#cfn-appsync-api-openidconnectconfig-iatttl
	//
	IatTtl *float64 `field:"optional" json:"iatTtl" yaml:"iatTtl"`
	// The issuer for the OIDC configuration.
	//
	// The issuer returned by discovery must exactly match the value of `iss` in the ID token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-openidconnectconfig.html#cfn-appsync-api-openidconnectconfig-issuer
	//
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
}

