package mixinsawsentityresolution

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsentityresolution/mixinsawsentityresolution/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an ID namespace object which will help customers provide metadata explaining their dataset and how to use it.
//
// Each ID namespace must have a unique name. To modify an existing ID namespace, use the UpdateIdNamespace API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdNamespacePropsMixin := awscdkmixinspreview.Mixins.NewCfnIdNamespacePropsMixin(&CfnIdNamespaceMixinProps{
//   	Description: jsii.String("description"),
//   	IdMappingWorkflowProperties: []interface{}{
//   		&IdNamespaceIdMappingWorkflowPropertiesProperty{
//   			IdMappingType: jsii.String("idMappingType"),
//   			ProviderProperties: &NamespaceProviderPropertiesProperty{
//   				ProviderConfiguration: map[string]*string{
//   					"providerConfigurationKey": jsii.String("providerConfiguration"),
//   				},
//   				ProviderServiceArn: jsii.String("providerServiceArn"),
//   			},
//   			RuleBasedProperties: &NamespaceRuleBasedPropertiesProperty{
//   				AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   				RecordMatchingModels: []*string{
//   					jsii.String("recordMatchingModels"),
//   				},
//   				RuleDefinitionTypes: []*string{
//   					jsii.String("ruleDefinitionTypes"),
//   				},
//   				Rules: []interface{}{
//   					&RuleProperty{
//   						MatchingKeys: []*string{
//   							jsii.String("matchingKeys"),
//   						},
//   						RuleName: jsii.String("ruleName"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	IdNamespaceName: jsii.String("idNamespaceName"),
//   	InputSourceConfig: []interface{}{
//   		&IdNamespaceInputSourceProperty{
//   			InputSourceArn: jsii.String("inputSourceArn"),
//   			SchemaName: jsii.String("schemaName"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idnamespace.html
//
type CfnIdNamespacePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIdNamespaceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIdNamespacePropsMixin
type jsiiProxy_CfnIdNamespacePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIdNamespacePropsMixin) Props() *CfnIdNamespaceMixinProps {
	var returns *CfnIdNamespaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdNamespacePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EntityResolution::IdNamespace`.
func NewCfnIdNamespacePropsMixin(props *CfnIdNamespaceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIdNamespacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIdNamespacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIdNamespacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdNamespacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EntityResolution::IdNamespace`.
func NewCfnIdNamespacePropsMixin_Override(c CfnIdNamespacePropsMixin, props *CfnIdNamespaceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdNamespacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIdNamespacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIdNamespacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdNamespacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIdNamespacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdNamespacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIdNamespacePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnIdNamespacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

