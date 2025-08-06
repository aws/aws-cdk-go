package awsverifiedpermissions


// Contains configuration details of an OpenID Connect (OIDC) identity provider, or identity source, that Verified Permissions can use to generate entities from authenticated identities.
//
// It specifies the issuer URL, token type that you want to use, and policy store entity details.
//
// This data type is part of a [Configuration](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_Configuration.html) structure, which is a parameter to [CreateIdentitySource](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreateIdentitySource.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openIdConnectConfigurationProperty := &OpenIdConnectConfigurationProperty{
//   	Issuer: jsii.String("issuer"),
//   	TokenSelection: &OpenIdConnectTokenSelectionProperty{
//   		AccessTokenOnly: &OpenIdConnectAccessTokenConfigurationProperty{
//   			Audiences: []*string{
//   				jsii.String("audiences"),
//   			},
//   			PrincipalIdClaim: jsii.String("principalIdClaim"),
//   		},
//   		IdentityTokenOnly: &OpenIdConnectIdentityTokenConfigurationProperty{
//   			ClientIds: []*string{
//   				jsii.String("clientIds"),
//   			},
//   			PrincipalIdClaim: jsii.String("principalIdClaim"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	EntityIdPrefix: jsii.String("entityIdPrefix"),
//   	GroupConfiguration: &OpenIdConnectGroupConfigurationProperty{
//   		GroupClaim: jsii.String("groupClaim"),
//   		GroupEntityType: jsii.String("groupEntityType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectconfiguration.html
//
type CfnIdentitySource_OpenIdConnectConfigurationProperty struct {
	// The issuer URL of an OIDC identity provider.
	//
	// This URL must have an OIDC discovery endpoint at the path `.well-known/openid-configuration` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectconfiguration.html#cfn-verifiedpermissions-identitysource-openidconnectconfiguration-issuer
	//
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// The token type that you want to process from your OIDC identity provider.
	//
	// Your policy store can process either identity (ID) or access tokens from a given OIDC identity source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectconfiguration.html#cfn-verifiedpermissions-identitysource-openidconnectconfiguration-tokenselection
	//
	TokenSelection interface{} `field:"required" json:"tokenSelection" yaml:"tokenSelection"`
	// A descriptive string that you want to prefix to user entities from your OIDC identity provider.
	//
	// For example, if you set an `entityIdPrefix` of `MyOIDCProvider` , you can reference principals in your policies in the format `MyCorp::User::MyOIDCProvider|Carlos` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectconfiguration.html#cfn-verifiedpermissions-identitysource-openidconnectconfiguration-entityidprefix
	//
	EntityIdPrefix *string `field:"optional" json:"entityIdPrefix" yaml:"entityIdPrefix"`
	// The claim in OIDC identity provider tokens that indicates a user's group membership, and the entity type that you want to map it to.
	//
	// For example, this object can map the contents of a `groups` claim to `MyCorp::UserGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectconfiguration.html#cfn-verifiedpermissions-identitysource-openidconnectconfiguration-groupconfiguration
	//
	GroupConfiguration interface{} `field:"optional" json:"groupConfiguration" yaml:"groupConfiguration"`
}

