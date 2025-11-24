package mixinsawsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsconfig/mixinsawsconfig/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds or updates an AWS Config rule for your entire organization to evaluate if your AWS resources comply with your desired configurations.
//
// For information on how many organization AWS Config rules you can have per account, see [*Service Limits*](https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html) in the *AWS Config Developer Guide* .
//
// Only a management account and a delegated administrator can create or update an organization AWS Config rule. When calling the `OrganizationConfigRule` resource with a delegated administrator, you must ensure AWS Organizations `ListDelegatedAdministrator` permissions are added. An organization can have up to 3 delegated administrators.
//
// The `OrganizationConfigRule` resource enables organization service access through the `EnableAWSServiceAccess` action and creates a service-linked role `AWSServiceRoleForConfigMultiAccountSetup` in the management or delegated administrator account of your organization. The service-linked role is created only when the role does not exist in the caller account. AWS Config verifies the existence of role with `GetRole` action.
//
// To use the `OrganizationConfigRule` resource with delegated administrator, register a delegated administrator by calling AWS Organization `register-delegated-administrator` for `config-multiaccountsetup.amazonaws.com` .
//
// There are two types of rules: *AWS Config Managed Rules* and *AWS Config Custom Rules* . You can use `PutOrganizationConfigRule` to create both AWS Config Managed Rules and AWS Config Custom Rules.
//
// AWS Config Managed Rules are predefined, customizable rules created by AWS Config . For a list of managed rules, see [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) . If you are adding an AWS Config managed rule, you must specify the rule's identifier for the `RuleIdentifier` key.
//
// AWS Config Custom Rules are rules that you create from scratch. There are two ways to create AWS Config custom rules: with Lambda functions ( [AWS Lambda Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/gettingstarted-concepts.html#gettingstarted-concepts-function) ) and with Guard ( [Guard GitHub Repository](https://docs.aws.amazon.com/https://github.com/aws-cloudformation/cloudformation-guard) ), a policy-as-code language. AWS Config custom rules created with AWS Lambda are called *AWS Config Custom Lambda Rules* and AWS Config custom rules created with Guard are called *AWS Config Custom Policy Rules* .
//
// If you are adding a new AWS Config Custom Lambda rule, you first need to create an AWS Lambda function in the management account or a delegated administrator that the rule invokes to evaluate your resources. You also need to create an IAM role in the managed account that can be assumed by the Lambda function. When you use `PutOrganizationConfigRule` to add a Custom Lambda rule to AWS Config , you must specify the Amazon Resource Name (ARN) that AWS Lambda assigns to the function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOrganizationConfigRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnOrganizationConfigRulePropsMixin(&CfnOrganizationConfigRuleMixinProps{
//   	ExcludedAccounts: []*string{
//   		jsii.String("excludedAccounts"),
//   	},
//   	OrganizationConfigRuleName: jsii.String("organizationConfigRuleName"),
//   	OrganizationCustomPolicyRuleMetadata: &OrganizationCustomPolicyRuleMetadataProperty{
//   		DebugLogDeliveryAccounts: []*string{
//   			jsii.String("debugLogDeliveryAccounts"),
//   		},
//   		Description: jsii.String("description"),
//   		InputParameters: jsii.String("inputParameters"),
//   		MaximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		OrganizationConfigRuleTriggerTypes: []*string{
//   			jsii.String("organizationConfigRuleTriggerTypes"),
//   		},
//   		PolicyText: jsii.String("policyText"),
//   		ResourceIdScope: jsii.String("resourceIdScope"),
//   		ResourceTypesScope: []*string{
//   			jsii.String("resourceTypesScope"),
//   		},
//   		Runtime: jsii.String("runtime"),
//   		TagKeyScope: jsii.String("tagKeyScope"),
//   		TagValueScope: jsii.String("tagValueScope"),
//   	},
//   	OrganizationCustomRuleMetadata: &OrganizationCustomRuleMetadataProperty{
//   		Description: jsii.String("description"),
//   		InputParameters: jsii.String("inputParameters"),
//   		LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   		MaximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		OrganizationConfigRuleTriggerTypes: []*string{
//   			jsii.String("organizationConfigRuleTriggerTypes"),
//   		},
//   		ResourceIdScope: jsii.String("resourceIdScope"),
//   		ResourceTypesScope: []*string{
//   			jsii.String("resourceTypesScope"),
//   		},
//   		TagKeyScope: jsii.String("tagKeyScope"),
//   		TagValueScope: jsii.String("tagValueScope"),
//   	},
//   	OrganizationManagedRuleMetadata: &OrganizationManagedRuleMetadataProperty{
//   		Description: jsii.String("description"),
//   		InputParameters: jsii.String("inputParameters"),
//   		MaximumExecutionFrequency: jsii.String("maximumExecutionFrequency"),
//   		ResourceIdScope: jsii.String("resourceIdScope"),
//   		ResourceTypesScope: []*string{
//   			jsii.String("resourceTypesScope"),
//   		},
//   		RuleIdentifier: jsii.String("ruleIdentifier"),
//   		TagKeyScope: jsii.String("tagKeyScope"),
//   		TagValueScope: jsii.String("tagValueScope"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconfigrule.html
//
type CfnOrganizationConfigRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnOrganizationConfigRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOrganizationConfigRulePropsMixin
type jsiiProxy_CfnOrganizationConfigRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnOrganizationConfigRulePropsMixin) Props() *CfnOrganizationConfigRuleMixinProps {
	var returns *CfnOrganizationConfigRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationConfigRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Config::OrganizationConfigRule`.
func NewCfnOrganizationConfigRulePropsMixin(props *CfnOrganizationConfigRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnOrganizationConfigRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOrganizationConfigRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOrganizationConfigRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnOrganizationConfigRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Config::OrganizationConfigRule`.
func NewCfnOrganizationConfigRulePropsMixin_Override(c CfnOrganizationConfigRulePropsMixin, props *CfnOrganizationConfigRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnOrganizationConfigRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnOrganizationConfigRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOrganizationConfigRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnOrganizationConfigRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOrganizationConfigRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_config.mixins.CfnOrganizationConfigRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOrganizationConfigRulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnOrganizationConfigRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

