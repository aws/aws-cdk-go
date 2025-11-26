package previewawsverifiedpermissionsmixins


// The token type that you want to process from your OIDC identity provider.
//
// Your policy store can process either identity (ID) or access tokens from a given OIDC identity source.
//
// This data type is part of a [OpenIdConnectConfiguration](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_OpenIdConnectConfiguration.html) structure, which is a parameter of [CreateIdentitySource](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreateIdentitySource.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   openIdConnectTokenSelectionProperty := &OpenIdConnectTokenSelectionProperty{
//   	AccessTokenOnly: &OpenIdConnectAccessTokenConfigurationProperty{
//   		Audiences: []*string{
//   			jsii.String("audiences"),
//   		},
//   		PrincipalIdClaim: jsii.String("principalIdClaim"),
//   	},
//   	IdentityTokenOnly: &OpenIdConnectIdentityTokenConfigurationProperty{
//   		ClientIds: []*string{
//   			jsii.String("clientIds"),
//   		},
//   		PrincipalIdClaim: jsii.String("principalIdClaim"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnecttokenselection.html
//
type CfnIdentitySourcePropsMixin_OpenIdConnectTokenSelectionProperty struct {
	// The OIDC configuration for processing access tokens.
	//
	// Contains allowed audience claims, for example `https://auth.example.com` , and the claim that you want to map to the principal, for example `sub` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnecttokenselection.html#cfn-verifiedpermissions-identitysource-openidconnecttokenselection-accesstokenonly
	//
	AccessTokenOnly interface{} `field:"optional" json:"accessTokenOnly" yaml:"accessTokenOnly"`
	// The OIDC configuration for processing identity (ID) tokens.
	//
	// Contains allowed client ID claims, for example `1example23456789` , and the claim that you want to map to the principal, for example `sub` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnecttokenselection.html#cfn-verifiedpermissions-identitysource-openidconnecttokenselection-identitytokenonly
	//
	IdentityTokenOnly interface{} `field:"optional" json:"identityTokenOnly" yaml:"identityTokenOnly"`
}

