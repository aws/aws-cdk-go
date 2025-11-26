package previewawsec2mixins


// Describes the OpenID Connect (OIDC) options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   nativeApplicationOidcOptionsProperty := &NativeApplicationOidcOptionsProperty{
//   	AuthorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//   	Issuer: jsii.String("issuer"),
//   	PublicSigningKeyEndpoint: jsii.String("publicSigningKeyEndpoint"),
//   	Scope: jsii.String("scope"),
//   	TokenEndpoint: jsii.String("tokenEndpoint"),
//   	UserInfoEndpoint: jsii.String("userInfoEndpoint"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions.html
//
type CfnVerifiedAccessTrustProviderPropsMixin_NativeApplicationOidcOptionsProperty struct {
	// The authorization endpoint of the IdP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions-authorizationendpoint
	//
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The OAuth 2.0 client identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions-clientid
	//
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The OAuth 2.0 client secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions-clientsecret
	//
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The OIDC issuer identifier of the IdP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions-issuer
	//
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// The public signing key endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions-publicsigningkeyendpoint
	//
	PublicSigningKeyEndpoint *string `field:"optional" json:"publicSigningKeyEndpoint" yaml:"publicSigningKeyEndpoint"`
	// The set of user claims to be requested from the IdP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The token endpoint of the IdP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions-tokenendpoint
	//
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The user info endpoint of the IdP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions.html#cfn-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions-userinfoendpoint
	//
	UserInfoEndpoint *string `field:"optional" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
}

