package previewawsssmincidentsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsssmincidents/previewawsssmincidentsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SSMIncidents::ResponsePlan` resource specifies the details of the response plan that are used when creating an incident.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResponsePlanPropsMixin := awscdkmixinspreview.Mixins.NewCfnResponsePlanPropsMixin(&CfnResponsePlanMixinProps{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			SsmAutomation: &SsmAutomationProperty{
//   				DocumentName: jsii.String("documentName"),
//   				DocumentVersion: jsii.String("documentVersion"),
//   				DynamicParameters: []interface{}{
//   					&DynamicSsmParameterProperty{
//   						Key: jsii.String("key"),
//   						Value: &DynamicSsmParameterValueProperty{
//   							Variable: jsii.String("variable"),
//   						},
//   					},
//   				},
//   				Parameters: []interface{}{
//   					&SsmParameterProperty{
//   						Key: jsii.String("key"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   				TargetAccount: jsii.String("targetAccount"),
//   			},
//   		},
//   	},
//   	ChatChannel: &ChatChannelProperty{
//   		ChatbotSns: []*string{
//   			jsii.String("chatbotSns"),
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	Engagements: []*string{
//   		jsii.String("engagements"),
//   	},
//   	IncidentTemplate: &IncidentTemplateProperty{
//   		DedupeString: jsii.String("dedupeString"),
//   		Impact: jsii.Number(123),
//   		IncidentTags: []interface{}{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		NotificationTargets: []interface{}{
//   			&NotificationTargetItemProperty{
//   				SnsTopicArn: jsii.String("snsTopicArn"),
//   			},
//   		},
//   		Summary: jsii.String("summary"),
//   		Title: jsii.String("title"),
//   	},
//   	Integrations: []interface{}{
//   		&IntegrationProperty{
//   			PagerDutyConfiguration: &PagerDutyConfigurationProperty{
//   				Name: jsii.String("name"),
//   				PagerDutyIncidentConfiguration: &PagerDutyIncidentConfigurationProperty{
//   					ServiceId: jsii.String("serviceId"),
//   				},
//   				SecretId: jsii.String("secretId"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html
//
type CfnResponsePlanPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnResponsePlanMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResponsePlanPropsMixin
type jsiiProxy_CfnResponsePlanPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnResponsePlanPropsMixin) Props() *CfnResponsePlanMixinProps {
	var returns *CfnResponsePlanMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResponsePlanPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSMIncidents::ResponsePlan`.
func NewCfnResponsePlanPropsMixin(props *CfnResponsePlanMixinProps, options *mixins.CfnPropertyMixinOptions) CfnResponsePlanPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResponsePlanPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResponsePlanPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSMIncidents::ResponsePlan`.
func NewCfnResponsePlanPropsMixin_Override(c CfnResponsePlanPropsMixin, props *CfnResponsePlanMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnResponsePlanPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResponsePlanPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResponsePlanPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResponsePlanPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnResponsePlanPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

