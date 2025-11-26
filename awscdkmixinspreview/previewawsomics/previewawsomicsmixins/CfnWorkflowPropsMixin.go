package previewawsomicsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsomics/previewawsomicsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a private workflow. Before you create a private workflow, you must create and configure these required resources:.
//
// - *Workflow definition file:* A workflow definition file written in WDL, Nextflow, or CWL. The workflow definition specifies the inputs and outputs for runs that use the workflow. It also includes specifications for the runs and run tasks for your workflow, including compute and memory requirements. The workflow definition file must be in `.zip` format. For more information, see [Workflow definition files](https://docs.aws.amazon.com/omics/latest/dev/workflow-definition-files.html) in AWS HealthOmics.
//
// - You can use Amazon Q CLI to build and validate your workflow definition files in WDL, Nextflow, and CWL. For more information, see [Example prompts for Amazon Q CLI](https://docs.aws.amazon.com/omics/latest/dev/getting-started.html#omics-q-prompts) and the [AWS HealthOmics Agentic generative AI tutorial](https://docs.aws.amazon.com/https://github.com/aws-samples/aws-healthomics-tutorials/tree/main/generative-ai) on GitHub.
// - *(Optional) Parameter template file:* A parameter template file written in JSON. Create the file to define the run parameters, or AWS HealthOmics generates the parameter template for you. For more information, see [Parameter template files for HealthOmics workflows](https://docs.aws.amazon.com/omics/latest/dev/parameter-templates.html) .
// - *ECR container images:* Create container images for the workflow in a private ECR repository, or synchronize images from a supported upstream registry with your Amazon ECR private repository.
// - *(Optional) Sentieon licenses:* Request a Sentieon license to use the Sentieon software in private workflows.
//
// For more information, see [Creating or updating a private workflow in AWS HealthOmics](https://docs.aws.amazon.com/omics/latest/dev/creating-private-workflows.html) in the *AWS HealthOmics User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkflowPropsMixin := awscdkmixinspreview.Mixins.NewCfnWorkflowPropsMixin(&CfnWorkflowMixinProps{
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
//   	Name: jsii.String("name"),
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
//   	WorkflowBucketOwnerId: jsii.String("workflowBucketOwnerId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-workflow.html
//
type CfnWorkflowPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnWorkflowMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkflowPropsMixin
type jsiiProxy_CfnWorkflowPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnWorkflowPropsMixin) Props() *CfnWorkflowMixinProps {
	var returns *CfnWorkflowMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflowPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Omics::Workflow`.
func NewCfnWorkflowPropsMixin(props *CfnWorkflowMixinProps, options *mixins.CfnPropertyMixinOptions) CfnWorkflowPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkflowPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkflowPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.mixins.CfnWorkflowPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Omics::Workflow`.
func NewCfnWorkflowPropsMixin_Override(c CfnWorkflowPropsMixin, props *CfnWorkflowMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.mixins.CfnWorkflowPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnWorkflowPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkflowPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.mixins.CfnWorkflowPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkflowPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_omics.mixins.CfnWorkflowPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkflowPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnWorkflowPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

