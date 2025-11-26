package previewawssesmixins


// Properties for CfnReceiptRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnReceiptRuleMixinProps := &CfnReceiptRuleMixinProps{
//   	After: jsii.String("after"),
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
//   					StatusCode: jsii.String("statusCode"),
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   				ConnectAction: &ConnectActionProperty{
//   					IamRoleArn: jsii.String("iamRoleArn"),
//   					InstanceArn: jsii.String("instanceArn"),
//   				},
//   				LambdaAction: &LambdaActionProperty{
//   					FunctionArn: jsii.String("functionArn"),
//   					InvocationType: jsii.String("invocationType"),
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   				S3Action: &S3ActionProperty{
//   					BucketName: jsii.String("bucketName"),
//   					IamRoleArn: jsii.String("iamRoleArn"),
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
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   				WorkmailAction: &WorkmailActionProperty{
//   					OrganizationArn: jsii.String("organizationArn"),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptrule.html
//
type CfnReceiptRuleMixinProps struct {
	// The name of an existing rule after which the new rule is placed.
	//
	// If this parameter is null, the new rule is inserted at the beginning of the rule list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptrule.html#cfn-ses-receiptrule-after
	//
	After *string `field:"optional" json:"after" yaml:"after"`
	// A data structure that contains the specified rule's name, actions, recipients, domains, enabled status, scan status, and TLS policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptrule.html#cfn-ses-receiptrule-rule
	//
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// The name of the rule set where the receipt rule is added.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptrule.html#cfn-ses-receiptrule-rulesetname
	//
	RuleSetName *string `field:"optional" json:"ruleSetName" yaml:"ruleSetName"`
}

