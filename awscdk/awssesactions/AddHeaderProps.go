package awssesactions


// Construction properties for a add header action.
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
type AddHeaderProps struct {
	// The name of the header to add.
	//
	// Must be between 1 and 50 characters,
	// inclusive, and consist of alphanumeric (a-z, A-Z, 0-9) characters
	// and dashes only.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the header to add.
	//
	// Must be less than 2048 characters,
	// and must not contain newline characters ("\r" or "\n").
	Value *string `field:"required" json:"value" yaml:"value"`
}

