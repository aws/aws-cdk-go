package previewawsverifiedpermissionsmixins


// The configuration of an OpenID Connect (OIDC) identity source for handling access token claims.
//
// Contains the claim that you want to identify as the principal in an authorization request, and the values of the `aud` claim, or audiences, that you want to accept.
//
// This data type is part of a [OpenIdConnectTokenSelection](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_OpenIdConnectTokenSelection.html) structure, which is a parameter of [CreateIdentitySource](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreateIdentitySource.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   openIdConnectAccessTokenConfigurationProperty := &OpenIdConnectAccessTokenConfigurationProperty{
//   	Audiences: []*string{
//   		jsii.String("audiences"),
//   	},
//   	PrincipalIdClaim: jsii.String("principalIdClaim"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectaccesstokenconfiguration.html
//
type CfnIdentitySourcePropsMixin_OpenIdConnectAccessTokenConfigurationProperty struct {
	// The access token `aud` claim values that you want to accept in your policy store.
	//
	// For example, `https://myapp.example.com, https://myapp2.example.com` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectaccesstokenconfiguration.html#cfn-verifiedpermissions-identitysource-openidconnectaccesstokenconfiguration-audiences
	//
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
	// The claim that determines the principal in OIDC access tokens.
	//
	// For example, `sub` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectaccesstokenconfiguration.html#cfn-verifiedpermissions-identitysource-openidconnectaccesstokenconfiguration-principalidclaim
	//
	// Default: - "sub".
	//
	PrincipalIdClaim *string `field:"optional" json:"principalIdClaim" yaml:"principalIdClaim"`
}

