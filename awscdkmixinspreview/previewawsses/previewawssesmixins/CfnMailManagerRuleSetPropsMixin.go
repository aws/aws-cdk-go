package previewawssesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsses/previewawssesmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource to create a rule set for a Mail Manager ingress endpoint which contains a list of rules that are evaluated sequentially for each email.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var drop interface{}
//
//   cfnMailManagerRuleSetPropsMixin := awscdkmixinspreview.Mixins.NewCfnMailManagerRuleSetPropsMixin(&CfnMailManagerRuleSetMixinProps{
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			Actions: []interface{}{
//   				&RuleActionProperty{
//   					AddHeader: &AddHeaderActionProperty{
//   						HeaderName: jsii.String("headerName"),
//   						HeaderValue: jsii.String("headerValue"),
//   					},
//   					Archive: &ArchiveActionProperty{
//   						ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   						TargetArchive: jsii.String("targetArchive"),
//   					},
//   					DeliverToMailbox: &DeliverToMailboxActionProperty{
//   						ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   						MailboxArn: jsii.String("mailboxArn"),
//   						RoleArn: jsii.String("roleArn"),
//   					},
//   					DeliverToQBusiness: &DeliverToQBusinessActionProperty{
//   						ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   						ApplicationId: jsii.String("applicationId"),
//   						IndexId: jsii.String("indexId"),
//   						RoleArn: jsii.String("roleArn"),
//   					},
//   					Drop: drop,
//   					PublishToSns: &SnsActionProperty{
//   						ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   						Encoding: jsii.String("encoding"),
//   						PayloadType: jsii.String("payloadType"),
//   						RoleArn: jsii.String("roleArn"),
//   						TopicArn: jsii.String("topicArn"),
//   					},
//   					Relay: &RelayActionProperty{
//   						ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   						MailFrom: jsii.String("mailFrom"),
//   						Relay: jsii.String("relay"),
//   					},
//   					ReplaceRecipient: &ReplaceRecipientActionProperty{
//   						ReplaceWith: []*string{
//   							jsii.String("replaceWith"),
//   						},
//   					},
//   					Send: &SendActionProperty{
//   						ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   						RoleArn: jsii.String("roleArn"),
//   					},
//   					WriteToS3: &S3ActionProperty{
//   						ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   						RoleArn: jsii.String("roleArn"),
//   						S3Bucket: jsii.String("s3Bucket"),
//   						S3Prefix: jsii.String("s3Prefix"),
//   						S3SseKmsKeyId: jsii.String("s3SseKmsKeyId"),
//   					},
//   				},
//   			},
//   			Conditions: []interface{}{
//   				&RuleConditionProperty{
//   					BooleanExpression: &RuleBooleanExpressionProperty{
//   						Evaluate: &RuleBooleanToEvaluateProperty{
//   							Analysis: &AnalysisProperty{
//   								Analyzer: jsii.String("analyzer"),
//   								ResultField: jsii.String("resultField"),
//   							},
//   							Attribute: jsii.String("attribute"),
//   							IsInAddressList: &RuleIsInAddressListProperty{
//   								AddressLists: []*string{
//   									jsii.String("addressLists"),
//   								},
//   								Attribute: jsii.String("attribute"),
//   							},
//   						},
//   						Operator: jsii.String("operator"),
//   					},
//   					DmarcExpression: &RuleDmarcExpressionProperty{
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					IpExpression: &RuleIpExpressionProperty{
//   						Evaluate: &RuleIpToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					NumberExpression: &RuleNumberExpressionProperty{
//   						Evaluate: &RuleNumberToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Value: jsii.Number(123),
//   					},
//   					StringExpression: &RuleStringExpressionProperty{
//   						Evaluate: &RuleStringToEvaluateProperty{
//   							Analysis: &AnalysisProperty{
//   								Analyzer: jsii.String("analyzer"),
//   								ResultField: jsii.String("resultField"),
//   							},
//   							Attribute: jsii.String("attribute"),
//   							MimeHeaderAttribute: jsii.String("mimeHeaderAttribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					VerdictExpression: &RuleVerdictExpressionProperty{
//   						Evaluate: &RuleVerdictToEvaluateProperty{
//   							Analysis: &AnalysisProperty{
//   								Analyzer: jsii.String("analyzer"),
//   								ResultField: jsii.String("resultField"),
//   							},
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			Unless: []interface{}{
//   				&RuleConditionProperty{
//   					BooleanExpression: &RuleBooleanExpressionProperty{
//   						Evaluate: &RuleBooleanToEvaluateProperty{
//   							Analysis: &AnalysisProperty{
//   								Analyzer: jsii.String("analyzer"),
//   								ResultField: jsii.String("resultField"),
//   							},
//   							Attribute: jsii.String("attribute"),
//   							IsInAddressList: &RuleIsInAddressListProperty{
//   								AddressLists: []*string{
//   									jsii.String("addressLists"),
//   								},
//   								Attribute: jsii.String("attribute"),
//   							},
//   						},
//   						Operator: jsii.String("operator"),
//   					},
//   					DmarcExpression: &RuleDmarcExpressionProperty{
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					IpExpression: &RuleIpExpressionProperty{
//   						Evaluate: &RuleIpToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					NumberExpression: &RuleNumberExpressionProperty{
//   						Evaluate: &RuleNumberToEvaluateProperty{
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Value: jsii.Number(123),
//   					},
//   					StringExpression: &RuleStringExpressionProperty{
//   						Evaluate: &RuleStringToEvaluateProperty{
//   							Analysis: &AnalysisProperty{
//   								Analyzer: jsii.String("analyzer"),
//   								ResultField: jsii.String("resultField"),
//   							},
//   							Attribute: jsii.String("attribute"),
//   							MimeHeaderAttribute: jsii.String("mimeHeaderAttribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   					VerdictExpression: &RuleVerdictExpressionProperty{
//   						Evaluate: &RuleVerdictToEvaluateProperty{
//   							Analysis: &AnalysisProperty{
//   								Analyzer: jsii.String("analyzer"),
//   								ResultField: jsii.String("resultField"),
//   							},
//   							Attribute: jsii.String("attribute"),
//   						},
//   						Operator: jsii.String("operator"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	RuleSetName: jsii.String("ruleSetName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanagerruleset.html
//
type CfnMailManagerRuleSetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMailManagerRuleSetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMailManagerRuleSetPropsMixin
type jsiiProxy_CfnMailManagerRuleSetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMailManagerRuleSetPropsMixin) Props() *CfnMailManagerRuleSetMixinProps {
	var returns *CfnMailManagerRuleSetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMailManagerRuleSetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SES::MailManagerRuleSet`.
func NewCfnMailManagerRuleSetPropsMixin(props *CfnMailManagerRuleSetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMailManagerRuleSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMailManagerRuleSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMailManagerRuleSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SES::MailManagerRuleSet`.
func NewCfnMailManagerRuleSetPropsMixin_Override(c CfnMailManagerRuleSetPropsMixin, props *CfnMailManagerRuleSetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMailManagerRuleSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMailManagerRuleSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMailManagerRuleSetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMailManagerRuleSetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMailManagerRuleSetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

