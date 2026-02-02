package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SageMaker::Model` resource to create a model to host at an Amazon SageMaker endpoint.
//
// For more information, see [Deploying a Model on Amazon SageMaker Hosting Services](https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works-hosting.html) in the *Amazon SageMaker Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var environment interface{}
//
//   cfnModelPropsMixin := awscdkmixinspreview.Mixins.NewCfnModelPropsMixin(&CfnModelMixinProps{
//   	Containers: []interface{}{
//   		&ContainerDefinitionProperty{
//   			ContainerHostname: jsii.String("containerHostname"),
//   			Environment: environment,
//   			Image: jsii.String("image"),
//   			ImageConfig: &ImageConfigProperty{
//   				RepositoryAccessMode: jsii.String("repositoryAccessMode"),
//   				RepositoryAuthConfig: &RepositoryAuthConfigProperty{
//   					RepositoryCredentialsProviderArn: jsii.String("repositoryCredentialsProviderArn"),
//   				},
//   			},
//   			InferenceSpecificationName: jsii.String("inferenceSpecificationName"),
//   			Mode: jsii.String("mode"),
//   			ModelDataSource: &ModelDataSourceProperty{
//   				S3DataSource: &S3DataSourceProperty{
//   					CompressionType: jsii.String("compressionType"),
//   					HubAccessConfig: &HubAccessConfigProperty{
//   						HubContentArn: jsii.String("hubContentArn"),
//   					},
//   					ModelAccessConfig: &ModelAccessConfigProperty{
//   						AcceptEula: jsii.Boolean(false),
//   					},
//   					S3DataType: jsii.String("s3DataType"),
//   					S3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//   			ModelDataUrl: jsii.String("modelDataUrl"),
//   			ModelPackageName: jsii.String("modelPackageName"),
//   			MultiModelConfig: &MultiModelConfigProperty{
//   				ModelCacheSetting: jsii.String("modelCacheSetting"),
//   			},
//   		},
//   	},
//   	EnableNetworkIsolation: jsii.Boolean(false),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	InferenceExecutionConfig: &InferenceExecutionConfigProperty{
//   		Mode: jsii.String("mode"),
//   	},
//   	ModelName: jsii.String("modelName"),
//   	PrimaryContainer: &ContainerDefinitionProperty{
//   		ContainerHostname: jsii.String("containerHostname"),
//   		Environment: environment,
//   		Image: jsii.String("image"),
//   		ImageConfig: &ImageConfigProperty{
//   			RepositoryAccessMode: jsii.String("repositoryAccessMode"),
//   			RepositoryAuthConfig: &RepositoryAuthConfigProperty{
//   				RepositoryCredentialsProviderArn: jsii.String("repositoryCredentialsProviderArn"),
//   			},
//   		},
//   		InferenceSpecificationName: jsii.String("inferenceSpecificationName"),
//   		Mode: jsii.String("mode"),
//   		ModelDataSource: &ModelDataSourceProperty{
//   			S3DataSource: &S3DataSourceProperty{
//   				CompressionType: jsii.String("compressionType"),
//   				HubAccessConfig: &HubAccessConfigProperty{
//   					HubContentArn: jsii.String("hubContentArn"),
//   				},
//   				ModelAccessConfig: &ModelAccessConfigProperty{
//   					AcceptEula: jsii.Boolean(false),
//   				},
//   				S3DataType: jsii.String("s3DataType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   		ModelDataUrl: jsii.String("modelDataUrl"),
//   		ModelPackageName: jsii.String("modelPackageName"),
//   		MultiModelConfig: &MultiModelConfigProperty{
//   			ModelCacheSetting: jsii.String("modelCacheSetting"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-model.html
//
type CfnModelPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnModelMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnModelPropsMixin
type jsiiProxy_CfnModelPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnModelPropsMixin) Props() *CfnModelMixinProps {
	var returns *CfnModelMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::Model`.
func NewCfnModelPropsMixin(props *CfnModelMixinProps, options *mixins.CfnPropertyMixinOptions) CfnModelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnModelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnModelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::Model`.
func NewCfnModelPropsMixin_Override(c CfnModelPropsMixin, props *CfnModelMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnModelPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnModelPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModelPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModelPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnModelPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

