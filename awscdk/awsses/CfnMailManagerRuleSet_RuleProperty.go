package awsses


// A rule contains conditions, "unless conditions" and actions.
//
// For each envelope recipient of an email, if all conditions match and none of the "unless conditions" match, then all of the actions are executed sequentially. If no conditions are provided, the rule always applies and the actions are implicitly executed. If only "unless conditions" are provided, the rule applies if the email does not match the evaluation of the "unless conditions".
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var drop interface{}
//
//   ruleProperty := &RuleProperty{
//   	Actions: []interface{}{
//   		&RuleActionProperty{
//   			AddHeader: &AddHeaderActionProperty{
//   				HeaderName: jsii.String("headerName"),
//   				HeaderValue: jsii.String("headerValue"),
//   			},
//   			Archive: &ArchiveActionProperty{
//   				TargetArchive: jsii.String("targetArchive"),
//
//   				// the properties below are optional
//   				ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   			},
//   			DeliverToMailbox: &DeliverToMailboxActionProperty{
//   				MailboxArn: jsii.String("mailboxArn"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   			},
//   			DeliverToQBusiness: &DeliverToQBusinessActionProperty{
//   				ApplicationId: jsii.String("applicationId"),
//   				IndexId: jsii.String("indexId"),
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   			},
//   			Drop: drop,
//   			PublishToSns: &SnsActionProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				TopicArn: jsii.String("topicArn"),
//
//   				// the properties below are optional
//   				ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   				Encoding: jsii.String("encoding"),
//   				PayloadType: jsii.String("payloadType"),
//   			},
//   			Relay: &RelayActionProperty{
//   				Relay: jsii.String("relay"),
//
//   				// the properties below are optional
//   				ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   				MailFrom: jsii.String("mailFrom"),
//   			},
//   			ReplaceRecipient: &ReplaceRecipientActionProperty{
//   				ReplaceWith: []*string{
//   					jsii.String("replaceWith"),
//   				},
//   			},
//   			Send: &SendActionProperty{
//   				RoleArn: jsii.String("roleArn"),
//
//   				// the properties below are optional
//   				ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   			},
//   			WriteToS3: &S3ActionProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				S3Bucket: jsii.String("s3Bucket"),
//
//   				// the properties below are optional
//   				ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   				S3Prefix: jsii.String("s3Prefix"),
//   				S3SseKmsKeyId: jsii.String("s3SseKmsKeyId"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Conditions: []interface{}{
//   		&RuleConditionProperty{
//   			BooleanExpression: &RuleBooleanExpressionProperty{
//   				Evaluate: &RuleBooleanToEvaluateProperty{
//   					Analysis: &AnalysisProperty{
//   						Analyzer: jsii.String("analyzer"),
//   						ResultField: jsii.String("resultField"),
//   					},
//   					Attribute: jsii.String("attribute"),
//   					IsInAddressList: &RuleIsInAddressListProperty{
//   						AddressLists: []*string{
//   							jsii.String("addressLists"),
//   						},
//   						Attribute: jsii.String("attribute"),
//   					},
//   				},
//   				Operator: jsii.String("operator"),
//   			},
//   			DmarcExpression: &RuleDmarcExpressionProperty{
//   				Operator: jsii.String("operator"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			IpExpression: &RuleIpExpressionProperty{
//   				Evaluate: &RuleIpToEvaluateProperty{
//   					Attribute: jsii.String("attribute"),
//   				},
//   				Operator: jsii.String("operator"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			NumberExpression: &RuleNumberExpressionProperty{
//   				Evaluate: &RuleNumberToEvaluateProperty{
//   					Attribute: jsii.String("attribute"),
//   				},
//   				Operator: jsii.String("operator"),
//   				Value: jsii.Number(123),
//   			},
//   			StringExpression: &RuleStringExpressionProperty{
//   				Evaluate: &RuleStringToEvaluateProperty{
//   					Analysis: &AnalysisProperty{
//   						Analyzer: jsii.String("analyzer"),
//   						ResultField: jsii.String("resultField"),
//   					},
//   					Attribute: jsii.String("attribute"),
//   					MimeHeaderAttribute: jsii.String("mimeHeaderAttribute"),
//   				},
//   				Operator: jsii.String("operator"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			VerdictExpression: &RuleVerdictExpressionProperty{
//   				Evaluate: &RuleVerdictToEvaluateProperty{
//   					Analysis: &AnalysisProperty{
//   						Analyzer: jsii.String("analyzer"),
//   						ResultField: jsii.String("resultField"),
//   					},
//   					Attribute: jsii.String("attribute"),
//   				},
//   				Operator: jsii.String("operator"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Unless: []interface{}{
//   		&RuleConditionProperty{
//   			BooleanExpression: &RuleBooleanExpressionProperty{
//   				Evaluate: &RuleBooleanToEvaluateProperty{
//   					Analysis: &AnalysisProperty{
//   						Analyzer: jsii.String("analyzer"),
//   						ResultField: jsii.String("resultField"),
//   					},
//   					Attribute: jsii.String("attribute"),
//   					IsInAddressList: &RuleIsInAddressListProperty{
//   						AddressLists: []*string{
//   							jsii.String("addressLists"),
//   						},
//   						Attribute: jsii.String("attribute"),
//   					},
//   				},
//   				Operator: jsii.String("operator"),
//   			},
//   			DmarcExpression: &RuleDmarcExpressionProperty{
//   				Operator: jsii.String("operator"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			IpExpression: &RuleIpExpressionProperty{
//   				Evaluate: &RuleIpToEvaluateProperty{
//   					Attribute: jsii.String("attribute"),
//   				},
//   				Operator: jsii.String("operator"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			NumberExpression: &RuleNumberExpressionProperty{
//   				Evaluate: &RuleNumberToEvaluateProperty{
//   					Attribute: jsii.String("attribute"),
//   				},
//   				Operator: jsii.String("operator"),
//   				Value: jsii.Number(123),
//   			},
//   			StringExpression: &RuleStringExpressionProperty{
//   				Evaluate: &RuleStringToEvaluateProperty{
//   					Analysis: &AnalysisProperty{
//   						Analyzer: jsii.String("analyzer"),
//   						ResultField: jsii.String("resultField"),
//   					},
//   					Attribute: jsii.String("attribute"),
//   					MimeHeaderAttribute: jsii.String("mimeHeaderAttribute"),
//   				},
//   				Operator: jsii.String("operator"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			VerdictExpression: &RuleVerdictExpressionProperty{
//   				Evaluate: &RuleVerdictToEvaluateProperty{
//   					Analysis: &AnalysisProperty{
//   						Analyzer: jsii.String("analyzer"),
//   						ResultField: jsii.String("resultField"),
//   					},
//   					Attribute: jsii.String("attribute"),
//   				},
//   				Operator: jsii.String("operator"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rule.html
//
type CfnMailManagerRuleSet_RuleProperty struct {
	// The list of actions to execute when the conditions match the incoming email, and none of the "unless conditions" match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rule.html#cfn-ses-mailmanagerruleset-rule-actions
	//
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The conditions of this rule.
	//
	// All conditions must match the email for the actions to be executed. An empty list of conditions means that all emails match, but are still subject to any "unless conditions"
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rule.html#cfn-ses-mailmanagerruleset-rule-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The user-friendly name of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rule.html#cfn-ses-mailmanagerruleset-rule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The "unless conditions" of this rule.
	//
	// None of the conditions can match the email for the actions to be executed. If any of these conditions do match the email, then the actions are not executed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-rule.html#cfn-ses-mailmanagerruleset-rule-unless
	//
	Unless interface{} `field:"optional" json:"unless" yaml:"unless"`
}

