package awsconfig


// An object that specifies organization custom rule metadata such as resource type, resource ID of AWS resource, Lambda function ARN, and organization trigger types that trigger AWS Config to evaluate your AWS resources against a rule.
//
// It also provides the frequency with which you want AWS Config to run evaluations for the rule if the trigger type is periodic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationCustomRuleMetadataProperty := &organizationCustomRuleMetadataProperty{
//   	lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   	organizationConfigRuleTriggerTypes: []*string{
//   		jsii.String("organizationConfigRuleTriggerTypes"),
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	inputParameters: jsii.String("inputParameters"),
//   	maximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   	resourceIdScope: jsii.String("resourceIdScope"),
//   	resourceTypesScope: []*string{
//   		jsii.String("resourceTypesScope"),
//   	},
//   	tagKeyScope: jsii.String("tagKeyScope"),
//   	tagValueScope: jsii.String("tagValueScope"),
//   }
//
type CfnOrganizationConfigRule_OrganizationCustomRuleMetadataProperty struct {
	// The lambda function ARN.
	LambdaFunctionArn *string `field:"required" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
	// The type of notification that triggers AWS Config to run an evaluation for a rule.
	//
	// You can specify the following notification types:
	//
	// - `ConfigurationItemChangeNotification` - Triggers an evaluation when AWS Config delivers a configuration item as a result of a resource change.
	// - `OversizedConfigurationItemChangeNotification` - Triggers an evaluation when AWS Config delivers an oversized configuration item. AWS Config may generate this notification type when a resource changes and the notification exceeds the maximum size allowed by Amazon SNS.
	// - `ScheduledNotification` - Triggers a periodic evaluation at the frequency specified for `MaximumExecutionFrequency` .
	OrganizationConfigRuleTriggerTypes *[]*string `field:"required" json:"organizationConfigRuleTriggerTypes" yaml:"organizationConfigRuleTriggerTypes"`
	// The description that you provide for your organization AWS Config rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A string, in JSON format, that is passed to your organization AWS Config rule Lambda function.
	InputParameters *string `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency with which AWS Config runs evaluations for a rule.
	//
	// Your custom rule is triggered when AWS Config delivers the configuration snapshot. For more information, see `ConfigSnapshotDeliveryProperties` .
	//
	// > By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the `MaximumExecutionFrequency` parameter.
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
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

