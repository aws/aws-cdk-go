package awsses


// Properties for defining a `CfnReceiptRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReceiptRuleProps := &CfnReceiptRuleProps{
//   	Rule: &RuleProperty{
//   		Actions: []interface{}{
//   			&ActionProperty{
//   				AddHeaderAction: &AddHeaderActionProperty{
//   					HeaderName: jsii.String("headerName"),
//   					HeaderValue: jsii.String("headerValue"),
//   				},
//   				BounceAction: &BounceActionProperty{
//   					Message: jsii.String("message"),
//   					Sender: jsii.String("sender"),
//   					SmtpReplyCode: jsii.String("smtpReplyCode"),
//
//   					// the properties below are optional
//   					StatusCode: jsii.String("statusCode"),
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   				LambdaAction: &LambdaActionProperty{
//   					FunctionArn: jsii.String("functionArn"),
//
//   					// the properties below are optional
//   					InvocationType: jsii.String("invocationType"),
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   				S3Action: &S3ActionProperty{
//   					BucketName: jsii.String("bucketName"),
//
//   					// the properties below are optional
//   					KmsKeyArn: jsii.String("kmsKeyArn"),
//   					ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   				SnsAction: &SNSActionProperty{
//   					Encoding: jsii.String("encoding"),
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   				StopAction: &StopActionProperty{
//   					Scope: jsii.String("scope"),
//
//   					// the properties below are optional
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   				WorkmailAction: &WorkmailActionProperty{
//   					OrganizationArn: jsii.String("organizationArn"),
//
//   					// the properties below are optional
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   			},
//   		},
//   		Enabled: jsii.Boolean(false),
//   		Name: jsii.String("name"),
//   		Recipients: []*string{
//   			jsii.String("recipients"),
//   		},
//   		ScanEnabled: jsii.Boolean(false),
//   		TlsPolicy: jsii.String("tlsPolicy"),
//   	},
//   	RuleSetName: jsii.String("ruleSetName"),
//
//   	// the properties below are optional
//   	After: jsii.String("after"),
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

