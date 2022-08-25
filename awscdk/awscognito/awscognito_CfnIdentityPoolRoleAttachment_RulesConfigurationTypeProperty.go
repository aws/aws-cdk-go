package awscognito


// `RulesConfigurationType` is a subproperty of the [RoleMapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html) property that defines the rules to be used for mapping users to roles.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rulesConfigurationTypeProperty := &rulesConfigurationTypeProperty{
//   	rules: []interface{}{
//   		&mappingRuleProperty{
//   			claim: jsii.String("claim"),
//   			matchType: jsii.String("matchType"),
//   			roleArn: jsii.String("roleArn"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnIdentityPoolRoleAttachment_RulesConfigurationTypeProperty struct {
	// The rules.
	//
	// You can specify up to 25 rules per identity provider.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

