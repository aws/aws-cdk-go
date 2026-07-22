package awsconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a rule for the specified Connect Customer instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var assignContactCategoryActions interface{}
//   var emptyValue interface{}
//   var endAssociatedTasksActions interface{}
//   var mergeStrategy IMergeStrategy
//
//   cfnRulePropsMixin := awscdkcfnpropertymixins.Aws_connect.NewCfnRulePropsMixin(&CfnRuleMixinProps{
//   	Actions: &ActionsProperty{
//   		AssignContactCategoryActions: []interface{}{
//   			assignContactCategoryActions,
//   		},
//   		AssignSlaActions: []interface{}{
//   			&AssignSlaActionProperty{
//   				CaseSlaConfiguration: &CaseSlaConfigurationProperty{
//   					FieldId: jsii.String("fieldId"),
//   					Name: jsii.String("name"),
//   					TargetFieldValues: []interface{}{
//   						&SlaTargetFieldValueProperty{
//   							StringValue: jsii.String("stringValue"),
//   						},
//   					},
//   					TargetSlaMinutes: jsii.Number(123),
//   					Type: jsii.String("type"),
//   				},
//   				SlaAssignmentType: jsii.String("slaAssignmentType"),
//   			},
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
//   				Exclusion: &NotificationRecipientTypeProperty{
//   					UserArns: []*string{
//   						jsii.String("userArns"),
//   					},
//   					UserTags: map[string]*string{
//   						"userTagsKey": jsii.String("userTags"),
//   					},
//   				},
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
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-rule.html
//
type CfnRulePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRuleMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRulePropsMixin
type jsiiProxy_CfnRulePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnRulePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::Rule`.
func NewCfnRulePropsMixin(props *CfnRuleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::Rule`.
func NewCfnRulePropsMixin_Override(c CfnRulePropsMixin, props *CfnRuleMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnRulePropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_connect.CfnRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRulePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

