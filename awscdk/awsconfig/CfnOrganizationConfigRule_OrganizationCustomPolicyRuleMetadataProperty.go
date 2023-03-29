package awsconfig


// An object that specifies metadata for your organization's AWS Config Custom Policy rule.
//
// The metadata includes the runtime system in use, which accounts have debug logging enabled, and other custom rule metadata, such as resource type, resource ID of AWS resource, and organization trigger types that initiate AWS Config to evaluate AWS resources against a rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationCustomPolicyRuleMetadataProperty := &OrganizationCustomPolicyRuleMetadataProperty{
//   	PolicyText: jsii.String("policyText"),
//   	Runtime: jsii.String("runtime"),
//
//   	// the properties below are optional
//   	DebugLogDeliveryAccounts: []*string{
//   		jsii.String("debugLogDeliveryAccounts"),
//   	},
//   	Description: jsii.String("description"),
//   	InputParameters: jsii.String("inputParameters"),
//   	MaximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   	OrganizationConfigRuleTriggerTypes: []*string{
//   		jsii.String("organizationConfigRuleTriggerTypes"),
//   	},
//   	ResourceIdScope: jsii.String("resourceIdScope"),
//   	ResourceTypesScope: []*string{
//   		jsii.String("resourceTypesScope"),
//   	},
//   	TagKeyScope: jsii.String("tagKeyScope"),
//   	TagValueScope: jsii.String("tagValueScope"),
//   }
//
type CfnOrganizationConfigRule_OrganizationCustomPolicyRuleMetadataProperty struct {
	// The policy definition containing the logic for your organization AWS Config Custom Policy rule.
	PolicyText *string `field:"required" json:"policyText" yaml:"policyText"`
	// `CfnOrganizationConfigRule.OrganizationCustomPolicyRuleMetadataProperty.Runtime`.
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// A list of accounts that you can enable debug logging for your organization AWS Config Custom Policy rule.
	//
	// List is null when debug logging is enabled for all accounts.
	DebugLogDeliveryAccounts *[]*string `field:"optional" json:"debugLogDeliveryAccounts" yaml:"debugLogDeliveryAccounts"`
	// The description that you provide for your organization AWS Config Custom Policy rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A string, in JSON format, that is passed to your organization AWS Config Custom Policy rule.
	InputParameters *string `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency with which AWS Config runs evaluations for a rule.
	//
	// Your AWS Config Custom Policy rule is triggered when AWS Config delivers the configuration snapshot. For more information, see `ConfigSnapshotDeliveryProperties` .
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// The type of notification that initiates AWS Config to run an evaluation for a rule.
	//
	// For AWS Config Custom Policy rules, AWS Config supports change-initiated notification types:
	//
	// - `ConfigurationItemChangeNotification` - Initiates an evaluation when AWS Config delivers a configuration item as a result of a resource change.
	// - `OversizedConfigurationItemChangeNotification` - Initiates an evaluation when AWS Config delivers an oversized configuration item. AWS Config may generate this notification type when a resource changes and the notification exceeds the maximum size allowed by Amazon SNS.
	OrganizationConfigRuleTriggerTypes *[]*string `field:"optional" json:"organizationConfigRuleTriggerTypes" yaml:"organizationConfigRuleTriggerTypes"`
	// The ID of the AWS resource that was evaluated.
	ResourceIdScope *string `field:"optional" json:"resourceIdScope" yaml:"resourceIdScope"`
	// The type of the AWS resource that was evaluated.
	ResourceTypesScope *[]*string `field:"optional" json:"resourceTypesScope" yaml:"resourceTypesScope"`
	// One part of a key-value pair that make up a tag.
	//
	// A key is a general label that acts like a category for more specific tag values.
	TagKeyScope *string `field:"optional" json:"tagKeyScope" yaml:"tagKeyScope"`
	// The optional part of a key-value pair that make up a tag.
	//
	// A value acts as a descriptor within a tag category (key).
	TagValueScope *string `field:"optional" json:"tagValueScope" yaml:"tagValueScope"`
}

