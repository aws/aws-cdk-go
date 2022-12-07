package awsconfig


// Construction properties for a CustomPolicy.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   samplePolicyText := "\n# This rule checks if point in time recovery (PITR) is enabled on active Amazon DynamoDB tables\nlet status = ['ACTIVE']\n\nrule tableisactive when\n    resourceType == \"AWS::DynamoDB::Table\" {\n    configuration.tableStatus == %status\n}\n\nrule checkcompliance when\n    resourceType == \"AWS::DynamoDB::Table\"\n    tableisactive {\n        let pitr = supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus\n        %pitr == \"ENABLED\"\n}\n"
//
//   config.NewCustomPolicy(stack, jsii.String("Custom"), &customPolicyProps{
//   	policyText: samplePolicyText,
//   	enableDebugLog: jsii.Boolean(true),
//   	ruleScope: config.ruleScope.fromResources([]resourceType{
//   		config.*resourceType_DYNAMODB_TABLE(),
//   	}),
//   })
//
type CustomPolicyProps struct {
	// A name for the AWS Config rule.
	ConfigRuleName *string `field:"optional" json:"configRuleName" yaml:"configRuleName"`
	// A description about this AWS Config rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Input parameter values that are passed to the AWS Config rule.
	InputParameters *map[string]interface{} `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency at which the AWS Config rule runs evaluations.
	MaximumExecutionFrequency MaximumExecutionFrequency `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources trigger an evaluation for an AWS Config rule.
	RuleScope RuleScope `field:"optional" json:"ruleScope" yaml:"ruleScope"`
	// The policy definition containing the logic for your AWS Config Custom Policy rule.
	PolicyText *string `field:"required" json:"policyText" yaml:"policyText"`
	// The boolean expression for enabling debug logging for your AWS Config Custom Policy rule.
	EnableDebugLog *bool `field:"optional" json:"enableDebugLog" yaml:"enableDebugLog"`
}

