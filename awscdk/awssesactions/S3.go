package awssesactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssesactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses"
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
type S3 interface {
	awsses.IReceiptRuleAction
	// Returns the receipt rule action specification.
	Bind(receiptRule interfacesawsses.IReceiptRuleRef) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for S3
type jsiiProxy_S3 struct {
	internal.Type__awssesIReceiptRuleAction
}

func NewS3(props *S3Props) S3 {
	_init_.Initialize()

	if err := validateNewS3Parameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.S3",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewS3_Override(s S3, props *S3Props) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.S3",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_S3) Bind(receiptRule interfacesawsses.IReceiptRuleRef) *awsses.ReceiptRuleActionConfig {
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

