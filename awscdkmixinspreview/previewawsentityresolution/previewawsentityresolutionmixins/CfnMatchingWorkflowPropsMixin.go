package previewawsentityresolutionmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsentityresolution/previewawsentityresolutionmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a matching workflow that defines the configuration for a data processing job.
//
// The workflow name must be unique. To modify an existing workflow, use `UpdateMatchingWorkflow` .
//
// > For workflows where `resolutionType` is `ML_MATCHING` or `PROVIDER` , incremental processing is not supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMatchingWorkflowPropsMixin := awscdkmixinspreview.Mixins.NewCfnMatchingWorkflowPropsMixin(&CfnMatchingWorkflowMixinProps{
//   	Description: jsii.String("description"),
//   	IncrementalRunConfig: &IncrementalRunConfigProperty{
//   		IncrementalRunType: jsii.String("incrementalRunType"),
//   	},
//   	InputSourceConfig: []interface{}{
//   		&InputSourceProperty{
//   			ApplyNormalization: jsii.Boolean(false),
//   			InputSourceArn: jsii.String("inputSourceArn"),
//   			SchemaArn: jsii.String("schemaArn"),
//   		},
//   	},
//   	OutputSourceConfig: []interface{}{
//   		&OutputSourceProperty{
//   			ApplyNormalization: jsii.Boolean(false),
//   			CustomerProfilesIntegrationConfig: &CustomerProfilesIntegrationConfigProperty{
//   				DomainArn: jsii.String("domainArn"),
//   				ObjectTypeArn: jsii.String("objectTypeArn"),
//   			},
//   			KmsArn: jsii.String("kmsArn"),
//   			Output: []interface{}{
//   				&OutputAttributeProperty{
//   					Hashed: jsii.Boolean(false),
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			OutputS3Path: jsii.String("outputS3Path"),
//   		},
//   	},
//   	ResolutionTechniques: &ResolutionTechniquesProperty{
//   		ProviderProperties: &ProviderPropertiesProperty{
//   			IntermediateSourceConfiguration: &IntermediateSourceConfigurationProperty{
//   				IntermediateS3Path: jsii.String("intermediateS3Path"),
//   			},
//   			ProviderConfiguration: map[string]*string{
//   				"providerConfigurationKey": jsii.String("providerConfiguration"),
//   			},
//   			ProviderServiceArn: jsii.String("providerServiceArn"),
//   		},
//   		ResolutionType: jsii.String("resolutionType"),
//   		RuleBasedProperties: &RuleBasedPropertiesProperty{
//   			AttributeMatchingModel: jsii.String("attributeMatchingModel"),
//   			MatchPurpose: jsii.String("matchPurpose"),
//   			Rules: []interface{}{
//   				&RuleProperty{
//   					MatchingKeys: []*string{
//   						jsii.String("matchingKeys"),
//   					},
//   					RuleName: jsii.String("ruleName"),
//   				},
//   			},
//   		},
//   		RuleConditionProperties: &RuleConditionPropertiesProperty{
//   			Rules: []interface{}{
//   				&RuleConditionProperty{
//   					Condition: jsii.String("condition"),
//   					RuleName: jsii.String("ruleName"),
//   				},
//   			},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-entityresolution-matchingworkflow.html
//
type CfnMatchingWorkflowPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMatchingWorkflowMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMatchingWorkflowPropsMixin
type jsiiProxy_CfnMatchingWorkflowPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMatchingWorkflowPropsMixin) Props() *CfnMatchingWorkflowMixinProps {
	var returns *CfnMatchingWorkflowMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMatchingWorkflowPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EntityResolution::MatchingWorkflow`.
func NewCfnMatchingWorkflowPropsMixin(props *CfnMatchingWorkflowMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMatchingWorkflowPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMatchingWorkflowPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMatchingWorkflowPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EntityResolution::MatchingWorkflow`.
func NewCfnMatchingWorkflowPropsMixin_Override(c CfnMatchingWorkflowPropsMixin, props *CfnMatchingWorkflowMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMatchingWorkflowPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMatchingWorkflowPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMatchingWorkflowPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_entityresolution.mixins.CfnMatchingWorkflowPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMatchingWorkflowPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMatchingWorkflowPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

