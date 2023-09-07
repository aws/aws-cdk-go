package awsconfig


// Provides the CustomPolicyDetails, the rule owner ( `AWS` for managed rules, `CUSTOM_POLICY` for Custom Policy rules, and `CUSTOM_LAMBDA` for Custom Lambda rules), the rule identifier, and the events that cause the evaluation of your AWS resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &SourceProperty{
//   	Owner: jsii.String("owner"),
//
//   	// the properties below are optional
//   	CustomPolicyDetails: &CustomPolicyDetailsProperty{
//   		EnableDebugLogDelivery: jsii.Boolean(false),
//   		PolicyRuntime: jsii.String("policyRuntime"),
//   		PolicyText: jsii.String("policyText"),
//   	},
//   	SourceDetails: []interface{}{
//   		&SourceDetailProperty{
//   			EventSource: jsii.String("eventSource"),
//   			MessageType: jsii.String("messageType"),
//
//   			// the properties below are optional
//   			MaximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		},
//   	},
//   	SourceIdentifier: jsii.String("sourceIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html
//
type CfnConfigRule_SourceProperty struct {
	// Indicates whether AWS or the customer owns and manages the AWS Config rule.
	//
	// AWS Config Managed Rules are predefined rules owned by AWS . For more information, see [AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) in the *AWS Config developer guide* .
	//
	// AWS Config Custom Rules are rules that you can develop either with Guard ( `CUSTOM_POLICY` ) or AWS Lambda ( `CUSTOM_LAMBDA` ). For more information, see [AWS Config Custom Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules.html) in the *AWS Config developer guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-owner
	//
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// Custom policy details when rule is custom owned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-custompolicydetails
	//
	CustomPolicyDetails interface{} `field:"optional" json:"customPolicyDetails" yaml:"customPolicyDetails"`
	// Provides the source and the message types that cause AWS Config to evaluate your AWS resources against a rule.
	//
	// It also provides the frequency with which you want AWS Config to run evaluations for the rule if the trigger type is periodic.
	//
	// If the owner is set to `CUSTOM_POLICY` , the only acceptable values for the AWS Config rule trigger message type are `ConfigurationItemChangeNotification` and `OversizedConfigurationItemChangeNotification` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-sourcedetails
	//
	SourceDetails interface{} `field:"optional" json:"sourceDetails" yaml:"sourceDetails"`
	// For AWS Config Managed rules, a predefined identifier from a list.
	//
	// For example, `IAM_PASSWORD_POLICY` is a managed rule. To reference a managed rule, see [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) .
	//
	// For AWS Config Custom Lambda rules, the identifier is the Amazon Resource Name (ARN) of the rule's AWS Lambda function, such as `arn:aws:lambda:us-east-2:123456789012:function:custom_rule_name` .
	//
	// For AWS Config Custom Policy rules, this field will be ignored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-sourceidentifier
	//
	SourceIdentifier *string `field:"optional" json:"sourceIdentifier" yaml:"sourceIdentifier"`
}

