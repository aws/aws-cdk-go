package previewawsimagebuildermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsimagebuilder/previewawsimagebuildermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An image pipeline is the automation configuration for building secure OS images on AWS .
//
// The Image Builder image pipeline is associated with an image recipe that defines the build, validation, and test phases for an image build lifecycle. An image pipeline can be associated with an infrastructure configuration that defines where your image is built. You can define attributes, such as instance types, a subnet for your VPC, security groups, logging, and other infrastructure-related configurations. You can also associate your image pipeline with a distribution configuration to define how you would like to deploy your image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnImagePipelinePropsMixin := awscdkmixinspreview.Mixins.NewCfnImagePipelinePropsMixin(&CfnImagePipelineMixinProps{
//   	ContainerRecipeArn: jsii.String("containerRecipeArn"),
//   	Description: jsii.String("description"),
//   	DistributionConfigurationArn: jsii.String("distributionConfigurationArn"),
//   	EnhancedImageMetadataEnabled: jsii.Boolean(false),
//   	ExecutionRole: jsii.String("executionRole"),
//   	ImageRecipeArn: jsii.String("imageRecipeArn"),
//   	ImageScanningConfiguration: &ImageScanningConfigurationProperty{
//   		EcrConfiguration: &EcrConfigurationProperty{
//   			ContainerTags: []*string{
//   				jsii.String("containerTags"),
//   			},
//   			RepositoryName: jsii.String("repositoryName"),
//   		},
//   		ImageScanningEnabled: jsii.Boolean(false),
//   	},
//   	ImageTestsConfiguration: &ImageTestsConfigurationProperty{
//   		ImageTestsEnabled: jsii.Boolean(false),
//   		TimeoutMinutes: jsii.Number(123),
//   	},
//   	InfrastructureConfigurationArn: jsii.String("infrastructureConfigurationArn"),
//   	LoggingConfiguration: &PipelineLoggingConfigurationProperty{
//   		ImageLogGroupName: jsii.String("imageLogGroupName"),
//   		PipelineLogGroupName: jsii.String("pipelineLogGroupName"),
//   	},
//   	Name: jsii.String("name"),
//   	Schedule: &ScheduleProperty{
//   		AutoDisablePolicy: &AutoDisablePolicyProperty{
//   			FailureCount: jsii.Number(123),
//   		},
//   		PipelineExecutionStartCondition: jsii.String("pipelineExecutionStartCondition"),
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	Status: jsii.String("status"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Workflows: []interface{}{
//   		&WorkflowConfigurationProperty{
//   			OnFailure: jsii.String("onFailure"),
//   			ParallelGroup: jsii.String("parallelGroup"),
//   			Parameters: []interface{}{
//   				&WorkflowParameterProperty{
//   					Name: jsii.String("name"),
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   			WorkflowArn: jsii.String("workflowArn"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagepipeline.html
//
type CfnImagePipelinePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnImagePipelineMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnImagePipelinePropsMixin
type jsiiProxy_CfnImagePipelinePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnImagePipelinePropsMixin) Props() *CfnImagePipelineMixinProps {
	var returns *CfnImagePipelineMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePipelinePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ImageBuilder::ImagePipeline`.
func NewCfnImagePipelinePropsMixin(props *CfnImagePipelineMixinProps, options *mixins.CfnPropertyMixinOptions) CfnImagePipelinePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnImagePipelinePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnImagePipelinePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePipelinePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ImageBuilder::ImagePipeline`.
func NewCfnImagePipelinePropsMixin_Override(c CfnImagePipelinePropsMixin, props *CfnImagePipelineMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePipelinePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnImagePipelinePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnImagePipelinePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePipelinePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnImagePipelinePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePipelinePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnImagePipelinePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnImagePipelinePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

