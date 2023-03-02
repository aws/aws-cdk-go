package awscognito


// Properties for defining a `CfnIdentityPoolRoleAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var roles interface{}
//
//   cfnIdentityPoolRoleAttachmentProps := &cfnIdentityPoolRoleAttachmentProps{
//   	identityPoolId: jsii.String("identityPoolId"),
//
//   	// the properties below are optional
//   	roleMappings: map[string]interface{}{
//   		"roleMappingsKey": &RoleMappingProperty{
//   			"type": jsii.String("type"),
//
//   			// the properties below are optional
//   			"ambiguousRoleResolution": jsii.String("ambiguousRoleResolution"),
//   			"identityProvider": jsii.String("identityProvider"),
//   			"rulesConfiguration": &RulesConfigurationTypeProperty{
//   				"rules": []interface{}{
//   					&MappingRuleProperty{
//   						"claim": jsii.String("claim"),
//   						"matchType": jsii.String("matchType"),
//   						"roleArn": jsii.String("roleArn"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	roles: roles,
//   }
//
type CfnIdentityPoolRoleAttachmentProps struct {
	// An identity pool ID in the format `REGION:GUID` .
	IdentityPoolId *string `field:"required" json:"identityPoolId" yaml:"identityPoolId"`
	// How users for a specific identity provider are mapped to roles.
	//
	// This is a string to the `RoleMapping` object map. The string identifies the identity provider. For example: `graph.facebook.com` or `cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id` .
	//
	// If the `IdentityProvider` field isn't provided in this object, the string is used as the identity provider name.
	//
	// For more information, see the [RoleMapping property](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html) .
	RoleMappings interface{} `field:"optional" json:"roleMappings" yaml:"roleMappings"`
	// The map of the roles associated with this pool.
	//
	// For a given role, the key is either "authenticated" or "unauthenticated". The value is the role ARN.
	Roles interface{} `field:"optional" json:"roles" yaml:"roles"`
}

