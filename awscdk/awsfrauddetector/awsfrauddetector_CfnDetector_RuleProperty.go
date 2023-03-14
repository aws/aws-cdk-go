package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// A rule.
//
// Rule is a condition that tells Amazon Fraud Detector how to interpret variables values during a fraud prediction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &ruleProperty{
//   	arn: jsii.String("arn"),
//   	createdTime: jsii.String("createdTime"),
//   	description: jsii.String("description"),
//   	detectorId: jsii.String("detectorId"),
//   	expression: jsii.String("expression"),
//   	language: jsii.String("language"),
//   	lastUpdatedTime: jsii.String("lastUpdatedTime"),
//   	outcomes: []interface{}{
//   		&outcomeProperty{
//   			arn: jsii.String("arn"),
//   			createdTime: jsii.String("createdTime"),
//   			description: jsii.String("description"),
//   			inline: jsii.Boolean(false),
//   			lastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			name: jsii.String("name"),
//   			tags: []cfnTag{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	ruleId: jsii.String("ruleId"),
//   	ruleVersion: jsii.String("ruleVersion"),
//   	tags: []*cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDetector_RuleProperty struct {
	// The rule ARN.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Timestamp for when the rule was created.
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// The rule description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The detector for which the rule is associated.
	DetectorId *string `field:"optional" json:"detectorId" yaml:"detectorId"`
	// The rule expression.
	//
	// A rule expression captures the business logic. For more information, see [Rule language reference](https://docs.aws.amazon.com/frauddetector/latest/ug/rule-language-reference.html) .
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The rule language.
	Language *string `field:"optional" json:"language" yaml:"language"`
	// Timestamp for when the rule was last updated.
	LastUpdatedTime *string `field:"optional" json:"lastUpdatedTime" yaml:"lastUpdatedTime"`
	// The rule outcome.
	Outcomes interface{} `field:"optional" json:"outcomes" yaml:"outcomes"`
	// The rule ID.
	RuleId *string `field:"optional" json:"ruleId" yaml:"ruleId"`
	// The rule version.
	RuleVersion *string `field:"optional" json:"ruleVersion" yaml:"ruleVersion"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

