package awssesactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

// Construction properties for a SNS action.
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
// Experimental.
type SnsProps struct {
	// The SNS topic to notify.
	// Experimental.
	Topic awssns.ITopic `field:"required" json:"topic" yaml:"topic"`
	// The encoding to use for the email within the Amazon SNS notification.
	// Experimental.
	Encoding EmailEncoding `field:"optional" json:"encoding" yaml:"encoding"`
}

