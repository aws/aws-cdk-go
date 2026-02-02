package previewawsimagebuildermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsimagebuilder/previewawsimagebuildermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new image.
//
// This request will create a new image along with all of the configured output resources defined in the distribution configuration. You must specify exactly one recipe for your image, using either a ContainerRecipeArn or an ImageRecipeArn.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnImagePropsMixin := awscdkmixinspreview.Mixins.NewCfnImagePropsMixin(&CfnImageMixinProps{
//   	ContainerRecipeArn: jsii.String("containerRecipeArn"),
//   	DeletionSettings: &DeletionSettingsProperty{
//   		ExecutionRole: jsii.String("executionRole"),
//   	},
//   	DistributionConfigurationArn: jsii.String("distributionConfigurationArn"),
//   	EnhancedImageMetadataEnabled: jsii.Boolean(false),
//   	ExecutionRole: jsii.String("executionRole"),
//   	ImagePipelineExecutionSettings: &ImagePipelineExecutionSettingsProperty{
//   		DeploymentId: jsii.String("deploymentId"),
//   		OnUpdate: jsii.Boolean(false),
//   	},
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
//   	LoggingConfiguration: &ImageLoggingConfigurationProperty{
//   		LogGroupName: jsii.String("logGroupName"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-image.html
//
type CfnImagePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnImageMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnImagePropsMixin
type jsiiProxy_CfnImagePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnImagePropsMixin) Props() *CfnImageMixinProps {
	var returns *CfnImageMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnImagePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ImageBuilder::Image`.
func NewCfnImagePropsMixin(props *CfnImageMixinProps, options *mixins.CfnPropertyMixinOptions) CfnImagePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnImagePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnImagePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ImageBuilder::Image`.
func NewCfnImagePropsMixin_Override(c CfnImagePropsMixin, props *CfnImageMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnImagePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnImagePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnImagePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnImagePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnImagePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

