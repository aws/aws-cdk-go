package awssesactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Construction properties for a S3 action.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import actions "github.com/aws/aws-cdk-go/awscdk"
//
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   ses.NewReceiptRuleSet(this, jsii.String("RuleSet"), &receiptRuleSetProps{
//   	rules: []receiptRuleOptions{
//   		&receiptRuleOptions{
//   			recipients: []*string{
//   				jsii.String("hello@aws.com"),
//   			},
//   			actions: []iReceiptRuleAction{
//   				actions.NewAddHeader(&addHeaderProps{
//   					name: jsii.String("X-Special-Header"),
//   					value: jsii.String("aws"),
//   				}),
//   				actions.NewS3(&s3Props{
//   					bucket: bucket,
//   					objectKeyPrefix: jsii.String("emails/"),
//   					topic: topic,
//   				}),
//   			},
//   		},
//   		&receiptRuleOptions{
//   			recipients: []*string{
//   				jsii.String("aws.com"),
//   			},
//   			actions: []*iReceiptRuleAction{
//   				actions.NewSns(&snsProps{
//   					topic: topic,
//   				}),
//   			},
//   		},
//   	},
//   })
//
type S3Props struct {
	// The S3 bucket that incoming email will be saved to.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The master key that SES should use to encrypt your emails before saving them to the S3 bucket.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The key prefix of the S3 bucket.
	ObjectKeyPrefix *string `field:"optional" json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
	// The SNS topic to notify when the S3 action is taken.
	Topic awssns.ITopic `field:"optional" json:"topic" yaml:"topic"`
}

