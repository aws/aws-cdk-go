package previewawsbudgetsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbudgets/previewawsbudgetsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Budgets::BudgetsAction` resource enables you to take predefined actions that are initiated when a budget threshold has been exceeded.
//
// For more information, see [Managing Your Costs with Budgets](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/budgets-managing-costs.html) in the *Billing and Cost Management User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBudgetsActionPropsMixin := awscdkmixinspreview.Mixins.NewCfnBudgetsActionPropsMixin(&CfnBudgetsActionMixinProps{
//   	ActionThreshold: &ActionThresholdProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	ActionType: jsii.String("actionType"),
//   	ApprovalModel: jsii.String("approvalModel"),
//   	BudgetName: jsii.String("budgetName"),
//   	Definition: &DefinitionProperty{
//   		IamActionDefinition: &IamActionDefinitionProperty{
//   			Groups: []*string{
//   				jsii.String("groups"),
//   			},
//   			PolicyArn: jsii.String("policyArn"),
//   			Roles: []*string{
//   				jsii.String("roles"),
//   			},
//   			Users: []*string{
//   				jsii.String("users"),
//   			},
//   		},
//   		ScpActionDefinition: &ScpActionDefinitionProperty{
//   			PolicyId: jsii.String("policyId"),
//   			TargetIds: []*string{
//   				jsii.String("targetIds"),
//   			},
//   		},
//   		SsmActionDefinition: &SsmActionDefinitionProperty{
//   			InstanceIds: []*string{
//   				jsii.String("instanceIds"),
//   			},
//   			Region: jsii.String("region"),
//   			Subtype: jsii.String("subtype"),
//   		},
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	NotificationType: jsii.String("notificationType"),
//   	ResourceTags: []ResourceTagProperty{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Subscribers: []interface{}{
//   		&SubscriberProperty{
//   			Address: jsii.String("address"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html
//
type CfnBudgetsActionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnBudgetsActionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBudgetsActionPropsMixin
type jsiiProxy_CfnBudgetsActionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnBudgetsActionPropsMixin) Props() *CfnBudgetsActionMixinProps {
	var returns *CfnBudgetsActionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBudgetsActionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Budgets::BudgetsAction`.
func NewCfnBudgetsActionPropsMixin(props *CfnBudgetsActionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnBudgetsActionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBudgetsActionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBudgetsActionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetsActionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Budgets::BudgetsAction`.
func NewCfnBudgetsActionPropsMixin_Override(c CfnBudgetsActionPropsMixin, props *CfnBudgetsActionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetsActionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnBudgetsActionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBudgetsActionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetsActionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBudgetsActionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetsActionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBudgetsActionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnBudgetsActionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

