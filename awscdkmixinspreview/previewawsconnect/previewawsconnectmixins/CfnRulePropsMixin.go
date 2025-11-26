package previewawsconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsconnect/previewawsconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a rule for the specified Amazon Connect instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var assignContactCategoryActions interface{}
//   var emptyValue interface{}
//   var endAssociatedTasksActions interface{}
//
//   cfnRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnRulePropsMixin(&CfnRuleMixinProps{
//   	Actions: &ActionsProperty{
//   		AssignContactCategoryActions: []interface{}{
//   			assignContactCategoryActions,
//   		},
//   		CreateCaseActions: []interface{}{
//   			&CreateCaseActionProperty{
//   				Fields: []interface{}{
//   					&FieldProperty{
//   						Id: jsii.String("id"),
//   						Value: &FieldValueProperty{
//   							BooleanValue: jsii.Boolean(false),
//   							DoubleValue: jsii.Number(123),
//   							EmptyValue: emptyValue,
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   				},
//   				TemplateId: jsii.String("templateId"),
//   			},
//   		},
//   		EndAssociatedTasksActions: []interface{}{
//   			endAssociatedTasksActions,
//   		},
//   		EventBridgeActions: []interface{}{
//   			&EventBridgeActionProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		SendNotificationActions: []interface{}{
//   			&SendNotificationActionProperty{
//   				Content: jsii.String("content"),
//   				ContentType: jsii.String("contentType"),
//   				DeliveryMethod: jsii.String("deliveryMethod"),
//   				Recipient: &NotificationRecipientTypeProperty{
//   					UserArns: []*string{
//   						jsii.String("userArns"),
//   					},
//   					UserTags: map[string]*string{
//   						"userTagsKey": jsii.String("userTags"),
//   					},
//   				},
//   				Subject: jsii.String("subject"),
//   			},
//   		},
//   		SubmitAutoEvaluationActions: []interface{}{
//   			&SubmitAutoEvaluationActionProperty{
//   				EvaluationFormArn: jsii.String("evaluationFormArn"),
//   			},
//   		},
//   		TaskActions: []interface{}{
//   			&TaskActionProperty{
//   				ContactFlowArn: jsii.String("contactFlowArn"),
//   				Description: jsii.String("description"),
//   				Name: jsii.String("name"),
//   				References: map[string]interface{}{
//   					"referencesKey": &ReferenceProperty{
//   						"type": jsii.String("type"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		UpdateCaseActions: []interface{}{
//   			&UpdateCaseActionProperty{
//   				Fields: []interface{}{
//   					&FieldProperty{
//   						Id: jsii.String("id"),
//   						Value: &FieldValueProperty{
//   							BooleanValue: jsii.Boolean(false),
//   							DoubleValue: jsii.Number(123),
//   							EmptyValue: emptyValue,
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Function: jsii.String("function"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	PublishStatus: jsii.String("publishStatus"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TriggerEventSource: &RuleTriggerEventSourceProperty{
//   		EventSourceName: jsii.String("eventSourceName"),
//   		IntegrationAssociationArn: jsii.String("integrationAssociationArn"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html
//
type CfnRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRulePropsMixin
type jsiiProxy_CfnRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRulePropsMixin) Props() *CfnRuleMixinProps {
	var returns *CfnRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::Rule`.
func NewCfnRulePropsMixin(props *CfnRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::Rule`.
func NewCfnRulePropsMixin_Override(c CfnRulePropsMixin, props *CfnRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

