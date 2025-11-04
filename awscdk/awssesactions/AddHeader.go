package awssesactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssesactions/internal"
)

// Adds a header to the received email.
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
type AddHeader interface {
	awsses.IReceiptRuleAction
	// Returns the receipt rule action specification.
	Bind(receiptRule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for AddHeader
type jsiiProxy_AddHeader struct {
	internal.Type__awssesIReceiptRuleAction
}

func NewAddHeader(props *AddHeaderProps) AddHeader {
	_init_.Initialize()

	if err := validateNewAddHeaderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AddHeader{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.AddHeader",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAddHeader_Override(a AddHeader, props *AddHeaderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.AddHeader",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AddHeader) Bind(receiptRule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	if err := a.validateBindParameters(receiptRule); err != nil {
		panic(err)
	}
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{receiptRule},
		&returns,
	)

	return returns
}

