package awsses


// Construction properties for a ReceiptRuleSet.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   ses.NewReceiptRuleSet(this, jsii.String("RuleSet"), &ReceiptRuleSetProps{
//   	Rules: []receiptRuleOptions{
//   		&receiptRuleOptions{
//   			Recipients: []*string{
//   				jsii.String("hello@aws.com"),
//   			},
//   			Actions: []iReceiptRuleAction{
//   				actions.NewAddHeader(&AddHeaderProps{
//   					Name: jsii.String("X-Special-Header"),
//   					Value: jsii.String("aws"),
//   				}),
//   				actions.NewS3(&S3Props{
//   					Bucket: *Bucket,
//   					ObjectKeyPrefix: jsii.String("emails/"),
//   					Topic: *Topic,
//   				}),
//   			},
//   		},
//   		&receiptRuleOptions{
//   			Recipients: []*string{
//   				jsii.String("aws.com"),
//   			},
//   			Actions: []*iReceiptRuleAction{
//   				actions.NewSns(&SnsProps{
//   					Topic: *Topic,
//   				}),
//   			},
//   		},
//   	},
//   })
//
type ReceiptRuleSetProps struct {
	// Whether to add a first rule to stop processing messages that have at least one spam indicator.
	// Default: false.
	//
	DropSpam *bool `field:"optional" json:"dropSpam" yaml:"dropSpam"`
	// The name for the receipt rule set.
	// Default: - A CloudFormation generated name.
	//
	ReceiptRuleSetName *string `field:"optional" json:"receiptRuleSetName" yaml:"receiptRuleSetName"`
	// The list of rules to add to this rule set.
	//
	// Rules are added in the same
	// order as they appear in the list.
	// Default: - No rules are added to the rule set.
	//
	Rules *[]*ReceiptRuleOptions `field:"optional" json:"rules" yaml:"rules"`
}

