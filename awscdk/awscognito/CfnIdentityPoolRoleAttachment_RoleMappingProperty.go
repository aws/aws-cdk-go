package awscognito


// One of a set of `RoleMappings` , a property of the [AWS::Cognito::IdentityPoolRoleAttachment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypoolroleattachment.html) resource that defines the role-mapping attributes of an Amazon Cognito identity pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   roleMappingProperty := &RoleMappingProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AmbiguousRoleResolution: jsii.String("ambiguousRoleResolution"),
//   	IdentityProvider: jsii.String("identityProvider"),
//   	RulesConfiguration: &RulesConfigurationTypeProperty{
//   		Rules: []interface{}{
//   			&MappingRuleProperty{
//   				Claim: jsii.String("claim"),
//   				MatchType: jsii.String("matchType"),
//   				RoleArn: jsii.String("roleArn"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html
//
type CfnIdentityPoolRoleAttachment_RoleMappingProperty struct {
	// The role mapping type.
	//
	// Token will use `cognito:roles` and `cognito:preferred_role` claims from the Cognito identity provider token to map groups to roles. Rules will attempt to match claims from the token to map to a role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html#cfn-cognito-identitypoolroleattachment-rolemapping-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// If you specify Token or Rules as the `Type` , `AmbiguousRoleResolution` is required.
	//
	// Specifies the action to be taken if either no rules match the claim value for the `Rules` type, or there is no `cognito:preferred_role` claim and there are multiple `cognito:roles` matches for the `Token` type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html#cfn-cognito-identitypoolroleattachment-rolemapping-ambiguousroleresolution
	//
	AmbiguousRoleResolution *string `field:"optional" json:"ambiguousRoleResolution" yaml:"ambiguousRoleResolution"`
	// Identifier for the identity provider for which the role is mapped.
	//
	// For example: `graph.facebook.com` or `cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id (http://cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id)` . This is the identity provider that is used by the user for authentication.
	//
	// If the identity provider property isn't provided, the key of the entry in the `RoleMappings` map is used as the identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html#cfn-cognito-identitypoolroleattachment-rolemapping-identityprovider
	//
	IdentityProvider *string `field:"optional" json:"identityProvider" yaml:"identityProvider"`
	// The rules to be used for mapping users to roles.
	//
	// If you specify "Rules" as the role-mapping type, RulesConfiguration is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html#cfn-cognito-identitypoolroleattachment-rolemapping-rulesconfiguration
	//
	RulesConfiguration interface{} `field:"optional" json:"rulesConfiguration" yaml:"rulesConfiguration"`
}

