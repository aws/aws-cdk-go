package awsconfig


// organization managed rule metadata such as resource type and ID of AWS resource along with the rule identifier.
//
// It also provides the frequency with which you want AWS Config to run evaluations for the rule if the trigger type is periodic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationManagedRuleMetadataProperty := &OrganizationManagedRuleMetadataProperty{
//   	RuleIdentifier: jsii.String("ruleIdentifier"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	InputParameters: jsii.String("inputParameters"),
//   	MaximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   	ResourceIdScope: jsii.String("resourceIdScope"),
//   	ResourceTypesScope: []*string{
//   		jsii.String("resourceTypesScope"),
//   	},
//   	TagKeyScope: jsii.String("tagKeyScope"),
//   	TagValueScope: jsii.String("tagValueScope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html
//
type CfnOrganizationConfigRule_OrganizationManagedRuleMetadataProperty struct {
	// For organization config managed rules, a predefined identifier from a list.
	//
	// For example, `IAM_PASSWORD_POLICY` is a managed rule. To reference a managed rule, see [Using AWS Config managed rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-ruleidentifier
	//
	RuleIdentifier *string `field:"required" json:"ruleIdentifier" yaml:"ruleIdentifier"`
	// The description that you provide for your organization AWS Config rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A string, in JSON format, that is passed to your organization AWS Config rule Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-inputparameters
	//
	InputParameters *string `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency with which AWS Config runs evaluations for a rule.
	//
	// This is for an AWS Config managed rule that is triggered at a periodic frequency.
	//
	// > By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the `MaximumExecutionFrequency` parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-maximumexecutionfrequency
	//
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// The ID of the AWS resource that was evaluated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-resourceidscope
	//
	ResourceIdScope *string `field:"optional" json:"resourceIdScope" yaml:"resourceIdScope"`
	// The type of the AWS resource that was evaluated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-resourcetypesscope
	//
	ResourceTypesScope *[]*string `field:"optional" json:"resourceTypesScope" yaml:"resourceTypesScope"`
	// One part of a key-value pair that make up a tag.
	//
	// A key is a general label that acts like a category for more specific tag values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-tagkeyscope
	//
	TagKeyScope *string `field:"optional" json:"tagKeyScope" yaml:"tagKeyScope"`
	// The optional part of a key-value pair that make up a tag.
	//
	// A value acts as a descriptor within a tag category (key).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html#cfn-config-organizationconfigrule-organizationmanagedrulemetadata-tagvaluescope
	//
	TagValueScope *string `field:"optional" json:"tagValueScope" yaml:"tagValueScope"`
}

