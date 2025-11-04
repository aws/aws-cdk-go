package awssesactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssesactions/internal"
)

// Publishes the email content within a notification to Amazon SNS.
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
//   	Rules: []ReceiptRuleOptions{
//   		&ReceiptRuleOptions{
//   			Recipients: []*string{
//   				jsii.String("hello@aws.com"),
//   			},
//   			Actions: []IReceiptRuleAction{
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
//   		&ReceiptRuleOptions{
//   			Recipients: []*string{
//   				jsii.String("aws.com"),
//   			},
//   			Actions: []IReceiptRuleAction{
//   				actions.NewSns(&SnsProps{
//   					Topic: *Topic,
//   				}),
//   			},
//   		},
//   	},
//   })
//
type Sns interface {
	awsses.IReceiptRuleAction
	// Returns the receipt rule action specification.
	Bind(receiptRule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for Sns
type jsiiProxy_Sns struct {
	internal.Type__awssesIReceiptRuleAction
}

func NewSns(props *SnsProps) Sns {
	_init_.Initialize()

	if err := validateNewSnsParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Sns{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.Sns",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewSns_Override(s Sns, props *SnsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.Sns",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_Sns) Bind(receiptRule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	if err := s.validateBindParameters(receiptRule); err != nil {
		panic(err)
	}
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{receiptRule},
		&returns,
	)

	return returns
}

