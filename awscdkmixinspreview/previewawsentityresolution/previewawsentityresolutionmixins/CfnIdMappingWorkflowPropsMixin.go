package previewawsentityresolutionmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsentityresolution/previewawsentityresolutionmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an `IdMappingWorkflow` object which stores the configuration of the data processing job to be run.
//
// Each `IdMappingWorkflow` must have a unique workflow name. To modify an existing workflow, use the UpdateIdMappingWorkflow API.
//
// > Incremental processing is not supported for ID mapping workflows.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdMappingWorkflowPropsMixin := awscdkmixinspreview.Mixins.NewCfnIdMappingWorkflowPropsMixin(&CfnIdMappingWorkflowMixinProps{
//   	Description: jsii.String("description"),
//   	IdMappingIncrementalRunConfig: &IdMappingIncrementalRunConfigProperty{
//   		IncrementalRunType: jsii.String("incrementalRunType"),
//   	},
//   	IdMappingTechniques: &IdMappingTechniquesProperty{
//   		IdMappingType: jsii.String("idMappingType"),
//   		NormalizationVersion: jsii.String("normalizationVersion"),
//   		ProviderProperties: &ProviderPropertiesProperty{
//   			IntermediateSourceConfiguration: &IntermediateSourceConfigurationProperty{
//   				IntermediateS3Path: jsii.String("intermediateS3Path"),
//   			},
//   			ProviderConfiguration: map[string]*string{
//   				"providerConfigurationKey": jsii.String("providerConfiguration"),
//   			},
//   			ProviderServiceArn: jsii.String("providerServiceArn"),
//   		},
//   		RuleBasedProperties: &IdMappingRuleBasedPropertiesProperty{
//   			AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   			RecordMatchingModel: jsii.String("recordMatchingModel"),
//   			RuleDefinitionType: jsii.String("ruleDefinitionType"),
//   			Rules: []interface{}{
//   				&RuleProperty{
//   					MatchingKeys: []*string{
//   						jsii.String("matchingKeys"),
//   					},
//   					RuleName: jsii.String("ruleName"),
//   				},
//   			},
//   		},
//   	},
//   	InputSourceConfig: []interface{}{
//   		&IdMappingWorkflowInputSourceProperty{
//   			InputSourceArn: jsii.String("inputSourceArn"),
//   			SchemaArn: jsii.String("schemaArn"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	OutputSourceConfig: []interface{}{
//   		&IdMappingWorkflowOutputSourceProperty{
//   			KmsArn: jsii.String("kmsArn"),
//   			OutputS3Path: jsii.String("outputS3Path"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkflowName: jsii.String("workflowName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-idmappingworkflow.html
//
type CfnIdMappingWorkflowPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIdMappingWorkflowMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIdMappingWorkflowPropsMixin
type jsiiProxy_CfnIdMappingWorkflowPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIdMappingWorkflowPropsMixin) Props() *CfnIdMappingWorkflowMixinProps {
	var returns *CfnIdMappingWorkflowMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIdMappingWorkflowPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EntityResolution::IdMappingWorkflow`.
func NewCfnIdMappingWorkflowPropsMixin(props *CfnIdMappingWorkflowMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIdMappingWorkflowPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIdMappingWorkflowPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIdMappingWorkflowPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EntityResolution::IdMappingWorkflow`.
func NewCfnIdMappingWorkflowPropsMixin_Override(c CfnIdMappingWorkflowPropsMixin, props *CfnIdMappingWorkflowMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIdMappingWorkflowPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIdMappingWorkflowPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIdMappingWorkflowPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnIdMappingWorkflowPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIdMappingWorkflowPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnIdMappingWorkflowPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

