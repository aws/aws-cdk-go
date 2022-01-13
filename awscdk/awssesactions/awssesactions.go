package awssesactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssesactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Adds a header to the received email.
//
// TODO: EXAMPLE
//
type AddHeader interface {
	awsses.IReceiptRuleAction
	Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for AddHeader
type jsiiProxy_AddHeader struct {
	internal.Type__awssesIReceiptRuleAction
}

func NewAddHeader(props *AddHeaderProps) AddHeader {
	_init_.Initialize()

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

// Returns the receipt rule action specification.
func (a *jsiiProxy_AddHeader) Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_rule},
		&returns,
	)

	return returns
}

// Construction properties for a add header action.
//
// TODO: EXAMPLE
//
type AddHeaderProps struct {
	// The name of the header to add.
	//
	// Must be between 1 and 50 characters,
	// inclusive, and consist of alphanumeric (a-z, A-Z, 0-9) characters
	// and dashes only.
	Name *string `json:"name" yaml:"name"`
	// The value of the header to add.
	//
	// Must be less than 2048 characters,
	// and must not contain newline characters ("\r" or "\n").
	Value *string `json:"value" yaml:"value"`
}

// Rejects the received email by returning a bounce response to the sender and, optionally, publishes a notification to Amazon SNS.
//
// TODO: EXAMPLE
//
type Bounce interface {
	awsses.IReceiptRuleAction
	Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for Bounce
type jsiiProxy_Bounce struct {
	internal.Type__awssesIReceiptRuleAction
}

func NewBounce(props *BounceProps) Bounce {
	_init_.Initialize()

	j := jsiiProxy_Bounce{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.Bounce",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewBounce_Override(b Bounce, props *BounceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.Bounce",
		[]interface{}{props},
		b,
	)
}

// Returns the receipt rule action specification.
func (b *jsiiProxy_Bounce) Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{_rule},
		&returns,
	)

	return returns
}

// Construction properties for a bounce action.
//
// TODO: EXAMPLE
//
type BounceProps struct {
	// The email address of the sender of the bounced email.
	//
	// This is the address
	// from which the bounce message will be sent.
	Sender *string `json:"sender" yaml:"sender"`
	// The template containing the message, reply code and status code.
	Template BounceTemplate `json:"template" yaml:"template"`
	// The SNS topic to notify when the bounce action is taken.
	Topic awssns.ITopic `json:"topic" yaml:"topic"`
}

// A bounce template.
//
// TODO: EXAMPLE
//
type BounceTemplate interface {
	Props() *BounceTemplateProps
}

// The jsii proxy struct for BounceTemplate
type jsiiProxy_BounceTemplate struct {
	_ byte // padding
}

func (j *jsiiProxy_BounceTemplate) Props() *BounceTemplateProps {
	var returns *BounceTemplateProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewBounceTemplate(props *BounceTemplateProps) BounceTemplate {
	_init_.Initialize()

	j := jsiiProxy_BounceTemplate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewBounceTemplate_Override(b BounceTemplate, props *BounceTemplateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		[]interface{}{props},
		b,
	)
}

func BounceTemplate_MAILBOX_DOES_NOT_EXIST() BounceTemplate {
	_init_.Initialize()
	var returns BounceTemplate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		"MAILBOX_DOES_NOT_EXIST",
		&returns,
	)
	return returns
}

func BounceTemplate_MAILBOX_FULL() BounceTemplate {
	_init_.Initialize()
	var returns BounceTemplate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		"MAILBOX_FULL",
		&returns,
	)
	return returns
}

func BounceTemplate_MESSAGE_CONTENT_REJECTED() BounceTemplate {
	_init_.Initialize()
	var returns BounceTemplate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		"MESSAGE_CONTENT_REJECTED",
		&returns,
	)
	return returns
}

func BounceTemplate_MESSAGE_TOO_LARGE() BounceTemplate {
	_init_.Initialize()
	var returns BounceTemplate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		"MESSAGE_TOO_LARGE",
		&returns,
	)
	return returns
}

func BounceTemplate_TEMPORARY_FAILURE() BounceTemplate {
	_init_.Initialize()
	var returns BounceTemplate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses_actions.BounceTemplate",
		"TEMPORARY_FAILURE",
		&returns,
	)
	return returns
}

// Construction properties for a BounceTemplate.
//
// TODO: EXAMPLE
//
type BounceTemplateProps struct {
	// Human-readable text to include in the bounce message.
	Message *string `json:"message" yaml:"message"`
	// The SMTP reply code, as defined by RFC 5321.
	// See: https://tools.ietf.org/html/rfc5321
	//
	SmtpReplyCode *string `json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// The SMTP enhanced status code, as defined by RFC 3463.
	// See: https://tools.ietf.org/html/rfc3463
	//
	StatusCode *string `json:"statusCode" yaml:"statusCode"`
}

// The type of email encoding to use for a SNS action.
type EmailEncoding string

const (
	EmailEncoding_BASE64 EmailEncoding = "BASE64"
	EmailEncoding_UTF8 EmailEncoding = "UTF8"
)

// Calls an AWS Lambda function, and optionally, publishes a notification to Amazon SNS.
//
// TODO: EXAMPLE
//
type Lambda interface {
	awsses.IReceiptRuleAction
	Bind(rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for Lambda
type jsiiProxy_Lambda struct {
	internal.Type__awssesIReceiptRuleAction
}

func NewLambda(props *LambdaProps) Lambda {
	_init_.Initialize()

	j := jsiiProxy_Lambda{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.Lambda",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewLambda_Override(l Lambda, props *LambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.Lambda",
		[]interface{}{props},
		l,
	)
}

// Returns the receipt rule action specification.
func (l *jsiiProxy_Lambda) Bind(rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// The type of invocation to use for a Lambda Action.
type LambdaInvocationType string

const (
	LambdaInvocationType_EVENT LambdaInvocationType = "EVENT"
	LambdaInvocationType_REQUEST_RESPONSE LambdaInvocationType = "REQUEST_RESPONSE"
)

// Construction properties for a Lambda action.
//
// TODO: EXAMPLE
//
type LambdaProps struct {
	// The Lambda function to invoke.
	Function awslambda.IFunction `json:"function" yaml:"function"`
	// The invocation type of the Lambda function.
	InvocationType LambdaInvocationType `json:"invocationType" yaml:"invocationType"`
	// The SNS topic to notify when the Lambda action is taken.
	Topic awssns.ITopic `json:"topic" yaml:"topic"`
}

// Saves the received message to an Amazon S3 bucket and, optionally, publishes a notification to Amazon SNS.
//
// TODO: EXAMPLE
//
type S3 interface {
	awsses.IReceiptRuleAction
	Bind(rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for S3
type jsiiProxy_S3 struct {
	internal.Type__awssesIReceiptRuleAction
}

func NewS3(props *S3Props) S3 {
	_init_.Initialize()

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

// Returns the receipt rule action specification.
func (s *jsiiProxy_S3) Bind(rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// Construction properties for a S3 action.
//
// TODO: EXAMPLE
//
type S3Props struct {
	// The S3 bucket that incoming email will be saved to.
	Bucket awss3.IBucket `json:"bucket" yaml:"bucket"`
	// The master key that SES should use to encrypt your emails before saving them to the S3 bucket.
	KmsKey awskms.IKey `json:"kmsKey" yaml:"kmsKey"`
	// The key prefix of the S3 bucket.
	ObjectKeyPrefix *string `json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
	// The SNS topic to notify when the S3 action is taken.
	Topic awssns.ITopic `json:"topic" yaml:"topic"`
}

// Publishes the email content within a notification to Amazon SNS.
//
// TODO: EXAMPLE
//
type Sns interface {
	awsses.IReceiptRuleAction
	Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for Sns
type jsiiProxy_Sns struct {
	internal.Type__awssesIReceiptRuleAction
}

func NewSns(props *SnsProps) Sns {
	_init_.Initialize()

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

// Returns the receipt rule action specification.
func (s *jsiiProxy_Sns) Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_rule},
		&returns,
	)

	return returns
}

// Construction properties for a SNS action.
//
// TODO: EXAMPLE
//
type SnsProps struct {
	// The SNS topic to notify.
	Topic awssns.ITopic `json:"topic" yaml:"topic"`
	// The encoding to use for the email within the Amazon SNS notification.
	Encoding EmailEncoding `json:"encoding" yaml:"encoding"`
}

// Terminates the evaluation of the receipt rule set and optionally publishes a notification to Amazon SNS.
//
// TODO: EXAMPLE
//
type Stop interface {
	awsses.IReceiptRuleAction
	Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for Stop
type jsiiProxy_Stop struct {
	internal.Type__awssesIReceiptRuleAction
}

func NewStop(props *StopProps) Stop {
	_init_.Initialize()

	j := jsiiProxy_Stop{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.Stop",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewStop_Override(s Stop, props *StopProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses_actions.Stop",
		[]interface{}{props},
		s,
	)
}

// Returns the receipt rule action specification.
func (s *jsiiProxy_Stop) Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_rule},
		&returns,
	)

	return returns
}

// Construction properties for a stop action.
//
// TODO: EXAMPLE
//
type StopProps struct {
	// The SNS topic to notify when the stop action is taken.
	Topic awssns.ITopic `json:"topic" yaml:"topic"`
}

