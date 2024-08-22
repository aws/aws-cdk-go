package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMailManagerRuleSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var drop interface{}
//
//   cfnMailManagerRuleSetProps := &CfnMailManagerRuleSetProps{
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			Actions: []interface{}{
//   				&RuleActionProperty{
//   					AddHeader: &AddHeaderActionProperty{
//   						HeaderName: jsii.String("headerName"),
//   						HeaderValue: jsii.String("headerValue"),
//   					},
//   					Archive: &ArchiveActionProperty{
//   						TargetArchive: jsii.String("targetArchive"),
//
//   						// the properties below are optional
//   						ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   					},
//   					DeliverToMailbox: &DeliverToMailboxActionProperty{
//   						MailboxArn: jsii.String("mailboxArn"),
//   						RoleArn: jsii.String("roleArn"),
//
//   						// the properties below are optional
//   						ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   					},
//   					Drop: drop,
//   					Relay: &RelayActionProperty{
//   						Relay: jsii.String("relay"),
//
//   						// the properties below are optional
//   						ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   						MailFrom: jsii.String("mailFrom"),
//   					},
//   					ReplaceRecipient: &ReplaceRecipientActionProperty{
//   						ReplaceWith: []*string{
//   							jsii.String("replaceWith"),
//   						},
//   					},
//   					Send: &SendActionProperty{
//   						RoleArn: jsii.String("roleArn"),
//
//   						// the properties below are optional
//   						ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   					},
//   					WriteToS3: &S3ActionProperty{
//   						RoleArn: jsii.String("roleArn"),
//   						S3Bucket: jsii.String("s3Bucket"),
//
//   						// the properties below are optional
//   						ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   						S3Prefix: jsii.String("s3Prefix"),
//   						S3SseKmsKeyId: jsii.String("s3SseKmsKeyId"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			Conditions: []interface{}{
//   				&RuleConditionProperty{
//   					BooleanExpression: &RuleBooleanExpressionProperty{
//   						Evaluate: &RuleBooleanToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   					},
//   					DmarcExpression: &RuleDmarcExpressionProperty{
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					IpExpression: &RuleIpExpressionProperty{
//   						Evaluate: &RuleIpToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					NumberExpression: &RuleNumberExpressionProperty{
//   						Evaluate: &RuleNumberToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Value: jsii.Number(123),
//   					},
//   					StringExpression: &RuleStringExpressionProperty{
//   						Evaluate: &RuleStringToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					VerdictExpression: &RuleVerdictExpressionProperty{
//   						Evaluate: &RuleVerdictToEvaluateProperty{
//   							Analysis: &AnalysisProperty{
//   								Analyzer: jsii.String("analyzer"),
//   								ResultField: jsii.String("resultField"),
//   							},
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			Unless: []interface{}{
//   				&RuleConditionProperty{
//   					BooleanExpression: &RuleBooleanExpressionProperty{
//   						Evaluate: &RuleBooleanToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   					},
//   					DmarcExpression: &RuleDmarcExpressionProperty{
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					IpExpression: &RuleIpExpressionProperty{
//   						Evaluate: &RuleIpToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					NumberExpression: &RuleNumberExpressionProperty{
//   						Evaluate: &RuleNumberToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Value: jsii.Number(123),
//   					},
//   					StringExpression: &RuleStringExpressionProperty{
//   						Evaluate: &RuleStringToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					VerdictExpression: &RuleVerdictExpressionProperty{
//   						Evaluate: &RuleVerdictToEvaluateProperty{
//   							Analysis: &AnalysisProperty{
//   								Analyzer: jsii.String("analyzer"),
//   								ResultField: jsii.String("resultField"),
//   							},
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	RuleSetName: jsii.String("ruleSetName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerruleset.html
//
type CfnMailManagerRuleSetProps struct {
	// Conditional rules that are evaluated for determining actions on email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerruleset.html#cfn-ses-mailmanagerruleset-rules
	//
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// A user-friendly name for the rule set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerruleset.html#cfn-ses-mailmanagerruleset-rulesetname
	//
	RuleSetName *string `field:"optional" json:"ruleSetName" yaml:"ruleSetName"`
	// The tags used to organize, track, or control access for the resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerruleset.html#cfn-ses-mailmanagerruleset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

