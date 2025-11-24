package mixinsawsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsses/mixinsawsses/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a receipt rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnReceiptRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnReceiptRulePropsMixin(&CfnReceiptRuleMixinProps{
//   	After: jsii.String("after"),
//   	Rule: &RuleProperty{
//   		Actions: []interface{}{
//   			&ActionProperty{
//   				AddHeaderAction: &AddHeaderActionProperty{
//   					HeaderName: jsii.String("headerName"),
//   					HeaderValue: jsii.String("headerValue"),
//   				},
//   				BounceAction: &BounceActionProperty{
//   					Message: jsii.String("message"),
//   					Sender: jsii.String("sender"),
//   					SmtpReplyCode: jsii.String("smtpReplyCode"),
//   					StatusCode: jsii.String("statusCode"),
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   				ConnectAction: &ConnectActionProperty{
//   					IamRoleArn: jsii.String("iamRoleArn"),
//   					InstanceArn: jsii.String("instanceArn"),
//   				},
//   				LambdaAction: &LambdaActionProperty{
//   					FunctionArn: jsii.String("functionArn"),
//   					InvocationType: jsii.String("invocationType"),
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   				S3Action: &S3ActionProperty{
//   					BucketName: jsii.String("bucketName"),
//   					IamRoleArn: jsii.String("iamRoleArn"),
//   					KmsKeyArn: jsii.String("kmsKeyArn"),
//   					ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   				SnsAction: &SNSActionProperty{
//   					Encoding: jsii.String("encoding"),
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   				StopAction: &StopActionProperty{
//   					Scope: jsii.String("scope"),
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   				WorkmailAction: &WorkmailActionProperty{
//   					OrganizationArn: jsii.String("organizationArn"),
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   			},
//   		},
//   		Enabled: jsii.Boolean(false),
//   		Name: jsii.String("name"),
//   		Recipients: []*string{
//   			jsii.String("recipients"),
//   		},
//   		ScanEnabled: jsii.Boolean(false),
//   		TlsPolicy: jsii.String("tlsPolicy"),
//   	},
//   	RuleSetName: jsii.String("ruleSetName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptrule.html
//
type CfnReceiptRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnReceiptRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnReceiptRulePropsMixin
type jsiiProxy_CfnReceiptRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnReceiptRulePropsMixin) Props() *CfnReceiptRuleMixinProps {
	var returns *CfnReceiptRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReceiptRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SES::ReceiptRule`.
func NewCfnReceiptRulePropsMixin(props *CfnReceiptRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnReceiptRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnReceiptRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnReceiptRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SES::ReceiptRule`.
func NewCfnReceiptRulePropsMixin_Override(c CfnReceiptRulePropsMixin, props *CfnReceiptRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnReceiptRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnReceiptRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReceiptRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnReceiptRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnReceiptRulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReceiptRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

