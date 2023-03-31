package awsconfig


// Properties for defining a `CfnOrganizationConfigRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOrganizationConfigRuleProps := &cfnOrganizationConfigRuleProps{
//   	organizationConfigRuleName: jsii.String("organizationConfigRuleName"),
//
//   	// the properties below are optional
//   	excludedAccounts: []*string{
//   		jsii.String("excludedAccounts"),
//   	},
//   	organizationCustomPolicyRuleMetadata: &organizationCustomPolicyRuleMetadataProperty{
//   		policyText: jsii.String("policyText"),
//   		runtime: jsii.String("runtime"),
//
//   		// the properties below are optional
//   		debugLogDeliveryAccounts: []*string{
//   			jsii.String("debugLogDeliveryAccounts"),
//   		},
//   		description: jsii.String("description"),
//   		inputParameters: jsii.String("inputParameters"),
//   		maximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		organizationConfigRuleTriggerTypes: []*string{
//   			jsii.String("organizationConfigRuleTriggerTypes"),
//   		},
//   		resourceIdScope: jsii.String("resourceIdScope"),
//   		resourceTypesScope: []*string{
//   			jsii.String("resourceTypesScope"),
//   		},
//   		tagKeyScope: jsii.String("tagKeyScope"),
//   		tagValueScope: jsii.String("tagValueScope"),
//   	},
//   	organizationCustomRuleMetadata: &organizationCustomRuleMetadataProperty{
//   		lambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   		organizationConfigRuleTriggerTypes: []*string{
//   			jsii.String("organizationConfigRuleTriggerTypes"),
//   		},
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		inputParameters: jsii.String("inputParameters"),
//   		maximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		resourceIdScope: jsii.String("resourceIdScope"),
//   		resourceTypesScope: []*string{
//   			jsii.String("resourceTypesScope"),
//   		},
//   		tagKeyScope: jsii.String("tagKeyScope"),
//   		tagValueScope: jsii.String("tagValueScope"),
//   	},
//   	organizationManagedRuleMetadata: &organizationManagedRuleMetadataProperty{
//   		ruleIdentifier: jsii.String("ruleIdentifier"),
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		inputParameters: jsii.String("inputParameters"),
//   		maximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		resourceIdScope: jsii.String("resourceIdScope"),
//   		resourceTypesScope: []*string{
//   			jsii.String("resourceTypesScope"),
//   		},
//   		tagKeyScope: jsii.String("tagKeyScope"),
//   		tagValueScope: jsii.String("tagValueScope"),
//   	},
//   }
//
type CfnOrganizationConfigRuleProps struct {
	// The name that you assign to organization AWS Config rule.
	OrganizationConfigRuleName *string `field:"required" json:"organizationConfigRuleName" yaml:"organizationConfigRuleName"`
	// A comma-separated list of accounts excluded from organization AWS Config rule.
	ExcludedAccounts *[]*string `field:"optional" json:"excludedAccounts" yaml:"excludedAccounts"`
	// An object that specifies metadata for your organization's AWS Config Custom Policy rule.
	//
	// The metadata includes the runtime system in use, which accounts have debug logging enabled, and other custom rule metadata, such as resource type, resource ID of AWS resource, and organization trigger types that initiate AWS Config to evaluate AWS resources against a rule.
	OrganizationCustomPolicyRuleMetadata interface{} `field:"optional" json:"organizationCustomPolicyRuleMetadata" yaml:"organizationCustomPolicyRuleMetadata"`
	// An `OrganizationCustomRuleMetadata` object.
	OrganizationCustomRuleMetadata interface{} `field:"optional" json:"organizationCustomRuleMetadata" yaml:"organizationCustomRuleMetadata"`
	// An `OrganizationManagedRuleMetadata` object.
	OrganizationManagedRuleMetadata interface{} `field:"optional" json:"organizationManagedRuleMetadata" yaml:"organizationManagedRuleMetadata"`
}

