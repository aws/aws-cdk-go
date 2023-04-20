package awssesactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsses"
	"github.com/aws/aws-cdk-go/awscdk/awssesactions/internal"
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
type Sns interface {
	awsses.IReceiptRuleAction
	// Returns the receipt rule action specification.
	// Experimental.
	Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for Sns
type jsiiProxy_Sns struct {
	internal.Type__awssesIReceiptRuleAction
}

// Experimental.
func NewSns(props *SnsProps) Sns {
	_init_.Initialize()

	if err := validateNewSnsParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Sns{}

	_jsii_.Create(
		"monocdk.aws_ses_actions.Sns",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSns_Override(s Sns, props *SnsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ses_actions.Sns",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_Sns) Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	if err := s.validateBindParameters(_rule); err != nil {
		panic(err)
	}
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_rule},
		&returns,
	)

	return returns
}

