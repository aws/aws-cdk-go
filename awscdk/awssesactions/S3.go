package awssesactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsses"
	"github.com/aws/aws-cdk-go/awscdk/awssesactions/internal"
)

// Saves the received message to an Amazon S3 bucket and, optionally, publishes a notification to Amazon SNS.
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
type S3 interface {
	awsses.IReceiptRuleAction
	// Returns the receipt rule action specification.
	// Experimental.
	Bind(rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for S3
type jsiiProxy_S3 struct {
	internal.Type__awssesIReceiptRuleAction
}

// Experimental.
func NewS3(props *S3Props) S3 {
	_init_.Initialize()

	if err := validateNewS3Parameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3{}

	_jsii_.Create(
		"monocdk.aws_ses_actions.S3",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3_Override(s S3, props *S3Props) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ses_actions.S3",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_S3) Bind(rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	if err := s.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

