package mixinsawsomics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsomics/mixinsawsomics/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new workflow version for the workflow that you specify with the `workflowId` parameter.
//
// When you create a new version of a workflow, you need to specify the configuration for the new version. It doesn't inherit any configuration values from the workflow.
//
// Provide a version name that is unique for this workflow. You cannot change the name after HealthOmics creates the version.
//
// > Don't include any personally identifiable information (PII) in the version name. Version names appear in the workflow version ARN.
//
// For more information, see [Workflow versioning in AWS HealthOmics](https://docs.aws.amazon.com/omics/latest/dev/workflow-versions.html) in the *AWS HealthOmics User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkflowVersionPropsMixin := awscdkmixinspreview.Mixins.NewCfnWorkflowVersionPropsMixin(&CfnWorkflowVersionMixinProps{
//   	Accelerators: jsii.String("accelerators"),
//   	ContainerRegistryMap: &ContainerRegistryMapProperty{
//   		ImageMappings: []interface{}{
//   			&ImageMappingProperty{
//   				DestinationImage: jsii.String("destinationImage"),
//   				SourceImage: jsii.String("sourceImage"),
//   			},
//   		},
//   		RegistryMappings: []interface{}{
//   			&RegistryMappingProperty{
//   				EcrAccountId: jsii.String("ecrAccountId"),
//   				EcrRepositoryPrefix: jsii.String("ecrRepositoryPrefix"),
//   				UpstreamRegistryUrl: jsii.String("upstreamRegistryUrl"),
//   				UpstreamRepositoryPrefix: jsii.String("upstreamRepositoryPrefix"),
//   			},
//   		},
//   	},
//   	ContainerRegistryMapUri: jsii.String("containerRegistryMapUri"),
//   	DefinitionRepository: &DefinitionRepositoryProperty{
//   		ConnectionArn: jsii.String("connectionArn"),
//   		ExcludeFilePatterns: []*string{
//   			jsii.String("excludeFilePatterns"),
//   		},
//   		FullRepositoryId: jsii.String("fullRepositoryId"),
//   		SourceReference: &SourceReferenceProperty{
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	DefinitionUri: jsii.String("definitionUri"),
//   	Description: jsii.String("description"),
//   	Engine: jsii.String("engine"),
//   	Main: jsii.String("main"),
//   	ParameterTemplate: map[string]interface{}{
//   		"parameterTemplateKey": &WorkflowParameterProperty{
//   			"description": jsii.String("description"),
//   			"optional": jsii.Boolean(false),
//   		},
//   	},
//   	ParameterTemplatePath: jsii.String("parameterTemplatePath"),
//   	ReadmeMarkdown: jsii.String("readmeMarkdown"),
//   	ReadmePath: jsii.String("readmePath"),
//   	ReadmeUri: jsii.String("readmeUri"),
//   	StorageCapacity: jsii.Number(123),
//   	StorageType: jsii.String("storageType"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	VersionName: jsii.String("versionName"),
//   	WorkflowBucketOwnerId: jsii.String("workflowBucketOwnerId"),
//   	WorkflowId: jsii.String("workflowId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflowversion.html
//
type CfnWorkflowVersionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnWorkflowVersionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkflowVersionPropsMixin
type jsiiProxy_CfnWorkflowVersionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnWorkflowVersionPropsMixin) Props() *CfnWorkflowVersionMixinProps {
	var returns *CfnWorkflowVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflowVersionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Omics::WorkflowVersion`.
func NewCfnWorkflowVersionPropsMixin(props *CfnWorkflowVersionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnWorkflowVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkflowVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkflowVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.mixins.CfnWorkflowVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Omics::WorkflowVersion`.
func NewCfnWorkflowVersionPropsMixin_Override(c CfnWorkflowVersionPropsMixin, props *CfnWorkflowVersionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.mixins.CfnWorkflowVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnWorkflowVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkflowVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.mixins.CfnWorkflowVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkflowVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_omics.mixins.CfnWorkflowVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkflowVersionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnWorkflowVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

