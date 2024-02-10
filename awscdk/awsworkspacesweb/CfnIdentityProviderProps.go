package awsworkspacesweb


// Properties for defining a `CfnIdentityProvider`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIdentityProviderProps := &CfnIdentityProviderProps{
//   	IdentityProviderDetails: map[string]*string{
//   		"identityProviderDetailsKey": jsii.String("identityProviderDetails"),
//   	},
//   	IdentityProviderName: jsii.String("identityProviderName"),
//   	IdentityProviderType: jsii.String("identityProviderType"),
//
//   	// the properties below are optional
//   	PortalArn: jsii.String("portalArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-identityprovider.html
//
type CfnIdentityProviderProps struct {
	// The identity provider details. The following list describes the provider detail keys for each identity provider type.
	//
	// - For Google and Login with Amazon:
	//
	// - `client_id`
	// - `client_secret`
	// - `authorize_scopes`
	// - For Facebook:
	//
	// - `client_id`
	// - `client_secret`
	// - `authorize_scopes`
	// - `api_version`
	// - For Sign in with Apple:
	//
	// - `client_id`
	// - `team_id`
	// - `key_id`
	// - `private_key`
	// - `authorize_scopes`
	// - For OIDC providers:
	//
	// - `client_id`
	// - `client_secret`
	// - `attributes_request_method`
	// - `oidc_issuer`
	// - `authorize_scopes`
	// - `authorize_url` *if not available from discovery URL specified by oidc_issuer key*
	// - `token_url` *if not available from discovery URL specified by oidc_issuer key*
	// - `attributes_url` *if not available from discovery URL specified by oidc_issuer key*
	// - `jwks_uri` *if not available from discovery URL specified by oidc_issuer key*
	// - For SAML providers:
	//
	// - `MetadataFile` OR `MetadataURL`
	// - `IDPSignout` (boolean) *optional*
	// - `IDPInit` (boolean) *optional*
	// - `RequestSigningAlgorithm` (string) *optional* - Only accepts `rsa-sha256`
	// - `EncryptedResponses` (boolean) *optional*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-identityprovider.html#cfn-workspacesweb-identityprovider-identityproviderdetails
	//
	IdentityProviderDetails interface{} `field:"required" json:"identityProviderDetails" yaml:"identityProviderDetails"`
	// The identity provider name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-identityprovider.html#cfn-workspacesweb-identityprovider-identityprovidername
	//
	IdentityProviderName *string `field:"required" json:"identityProviderName" yaml:"identityProviderName"`
	// The identity provider type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-identityprovider.html#cfn-workspacesweb-identityprovider-identityprovidertype
	//
	IdentityProviderType *string `field:"required" json:"identityProviderType" yaml:"identityProviderType"`
	// The ARN of the identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesweb-identityprovider.html#cfn-workspacesweb-identityprovider-portalarn
	//
	PortalArn *string `field:"optional" json:"portalArn" yaml:"portalArn"`
}

