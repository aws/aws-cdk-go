package awsconfig


// Construction properties for a CustomPolicy.
//
// Example:
//   samplePolicyText := `
//   # This rule checks if point in time recovery (PITR) is enabled on active Amazon DynamoDB tables
//   let status = ['ACTIVE']
//
//   rule tableisactive when
//       resourceType == "AWS::DynamoDB::Table" {
//       configuration.tableStatus == %status
//   }
//
//   rule checkcompliance when
//       resourceType == "AWS::DynamoDB::Table"
//       tableisactive {
//           let pitr = supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus
//           %pitr == "ENABLED"
//   }
//   `
//
//   config.NewCustomPolicy(this, jsii.String("Custom"), &CustomPolicyProps{
//   	PolicyText: samplePolicyText,
//   	EnableDebugLog: jsii.Boolean(true),
//   	RuleScope: config.RuleScope_FromResources([]resourceType{
//   		config.*resourceType_DYNAMODB_TABLE(),
//   	}),
//   })
//
type CustomPolicyProps struct {
	// A name for the AWS Config rule.
	// Default: - CloudFormation generated name.
	//
	ConfigRuleName *string `field:"optional" json:"configRuleName" yaml:"configRuleName"`
	// A description about this AWS Config rule.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The modes the AWS Config rule can be evaluated in.
	//
	// The valid values are distinct objects.
	// Default: - Detective evaluation mode only.
	//
	EvaluationModes EvaluationMode `field:"optional" json:"evaluationModes" yaml:"evaluationModes"`
	// Input parameter values that are passed to the AWS Config rule.
	// Default: - No input parameters.
	//
	InputParameters *map[string]interface{} `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency at which the AWS Config rule runs evaluations.
	// Default: MaximumExecutionFrequency.TWENTY_FOUR_HOURS
	//
	MaximumExecutionFrequency MaximumExecutionFrequency `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources trigger an evaluation for an AWS Config rule.
	// Default: - evaluations for the rule are triggered when any resource in the recording group changes.
	//
	RuleScope RuleScope `field:"optional" json:"ruleScope" yaml:"ruleScope"`
	// The policy definition containing the logic for your AWS Config Custom Policy rule.
	PolicyText *string `field:"required" json:"policyText" yaml:"policyText"`
	// The boolean expression for enabling debug logging for your AWS Config Custom Policy rule.
	// Default: false.
	//
	EnableDebugLog *bool `field:"optional" json:"enableDebugLog" yaml:"enableDebugLog"`
}

