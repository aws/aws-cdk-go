package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
//   ruleProperty := &RuleProperty{
//   	Arn: jsii.String("arn"),
//   	CreatedTime: jsii.String("createdTime"),
//   	Description: jsii.String("description"),
//   	DetectorId: jsii.String("detectorId"),
//   	Expression: jsii.String("expression"),
//   	Language: jsii.String("language"),
//   	LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   	Outcomes: []interface{}{
//   		&OutcomeProperty{
//   			Arn: jsii.String("arn"),
//   			CreatedTime: jsii.String("createdTime"),
//   			Description: jsii.String("description"),
//   			Inline: jsii.Boolean(false),
//   			LastUpdatedTime: jsii.String("lastUpdatedTime"),
//   			Name: jsii.String("name"),
//   			Tags: []cfnTag{
//   				&cfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	RuleId: jsii.String("ruleId"),
//   	RuleVersion: jsii.String("ruleVersion"),
//   	Tags: []*cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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

