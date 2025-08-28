package awscognito


// `RulesConfigurationType` is a subproperty of the [RoleMapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html) property that defines the rules to be used for mapping users to roles.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rulesConfigurationTypeProperty := &RulesConfigurationTypeProperty{
//   	Rules: []interface{}{
//   		&MappingRuleProperty{
//   			Claim: jsii.String("claim"),
//   			MatchType: jsii.String("matchType"),
//   			RoleArn: jsii.String("roleArn"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rulesconfigurationtype.html
//
type CfnIdentityPoolRoleAttachment_RulesConfigurationTypeProperty struct {
	// The rules.
	//
	// You can specify up to 25 rules per identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rulesconfigurationtype.html#cfn-cognito-identitypoolroleattachment-rulesconfigurationtype-rules
	//
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

