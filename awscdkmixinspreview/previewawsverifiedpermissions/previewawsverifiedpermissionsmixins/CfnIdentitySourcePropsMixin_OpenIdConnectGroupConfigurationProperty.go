package previewawsverifiedpermissionsmixins


// The claim in OIDC identity provider tokens that indicates a user's group membership, and the entity type that you want to map it to.
//
// For example, this object can map the contents of a `groups` claim to `MyCorp::UserGroup` .
//
// This data type is part of a [OpenIdConnectConfiguration](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_OpenIdConnectConfiguration.html) structure, which is a parameter of [CreateIdentitySource](https://docs.aws.amazon.com/verifiedpermissions/latest/apireference/API_CreateIdentitySource.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   openIdConnectGroupConfigurationProperty := &OpenIdConnectGroupConfigurationProperty{
//   	GroupClaim: jsii.String("groupClaim"),
//   	GroupEntityType: jsii.String("groupEntityType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectgroupconfiguration.html
//
type CfnIdentitySourcePropsMixin_OpenIdConnectGroupConfigurationProperty struct {
	// The token claim that you want Verified Permissions to interpret as group membership.
	//
	// For example, `groups` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectgroupconfiguration.html#cfn-verifiedpermissions-identitysource-openidconnectgroupconfiguration-groupclaim
	//
	GroupClaim *string `field:"optional" json:"groupClaim" yaml:"groupClaim"`
	// The policy store entity type that you want to map your users' group claim to.
	//
	// For example, `MyCorp::UserGroup` . A group entity type is an entity that can have a user entity type as a member.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-verifiedpermissions-identitysource-openidconnectgroupconfiguration.html#cfn-verifiedpermissions-identitysource-openidconnectgroupconfiguration-groupentitytype
	//
	GroupEntityType *string `field:"optional" json:"groupEntityType" yaml:"groupEntityType"`
}

