package awsses


// Properties for defining a `CfnReceiptRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReceiptRuleProps := &cfnReceiptRuleProps{
//   	rule: &ruleProperty{
//   		actions: []interface{}{
//   			&actionProperty{
//   				addHeaderAction: &addHeaderActionProperty{
//   					headerName: jsii.String("headerName"),
//   					headerValue: jsii.String("headerValue"),
//   				},
//   				bounceAction: &bounceActionProperty{
//   					message: jsii.String("message"),
//   					sender: jsii.String("sender"),
//   					smtpReplyCode: jsii.String("smtpReplyCode"),
//
//   					// the properties below are optional
//   					statusCode: jsii.String("statusCode"),
//   					topicArn: jsii.String("topicArn"),
//   				},
//   				lambdaAction: &lambdaActionProperty{
//   					functionArn: jsii.String("functionArn"),
//
//   					// the properties below are optional
//   					invocationType: jsii.String("invocationType"),
//   					topicArn: jsii.String("topicArn"),
//   				},
//   				s3Action: &s3ActionProperty{
//   					bucketName: jsii.String("bucketName"),
//
//   					// the properties below are optional
//   					kmsKeyArn: jsii.String("kmsKeyArn"),
//   					objectKeyPrefix: jsii.String("objectKeyPrefix"),
//   					topicArn: jsii.String("topicArn"),
//   				},
//   				snsAction: &sNSActionProperty{
//   					encoding: jsii.String("encoding"),
//   					topicArn: jsii.String("topicArn"),
//   				},
//   				stopAction: &stopActionProperty{
//   					scope: jsii.String("scope"),
//
//   					// the properties below are optional
//   					topicArn: jsii.String("topicArn"),
//   				},
//   				workmailAction: &workmailActionProperty{
//   					organizationArn: jsii.String("organizationArn"),
//
//   					// the properties below are optional
//   					topicArn: jsii.String("topicArn"),
//   				},
//   			},
//   		},
//   		enabled: jsii.Boolean(false),
//   		name: jsii.String("name"),
//   		recipients: []*string{
//   			jsii.String("recipients"),
//   		},
//   		scanEnabled: jsii.Boolean(false),
//   		tlsPolicy: jsii.String("tlsPolicy"),
//   	},
//   	ruleSetName: jsii.String("ruleSetName"),
//
//   	// the properties below are optional
//   	after: jsii.String("after"),
//   }
//
type CfnReceiptRuleProps struct {
	// A data structure that contains the specified rule's name, actions, recipients, domains, enabled status, scan status, and TLS policy.
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
	// The name of the rule set where the receipt rule is added.
	RuleSetName *string `field:"required" json:"ruleSetName" yaml:"ruleSetName"`
	// The name of an existing rule after which the new rule is placed.
	//
	// If this parameter is null, the new rule is inserted at the beginning of the rule list.
	After *string `field:"optional" json:"after" yaml:"after"`
}

