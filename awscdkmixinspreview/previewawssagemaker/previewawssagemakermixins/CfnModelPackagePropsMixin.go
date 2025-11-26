package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A container for your trained model that can be deployed for SageMaker inference.
//
// This can include inference code, artifacts, and metadata. The model package type can be one of the following.
//
// - Versioned model: A part of a model package group in Model Registry.
// - Unversioned model: Not part of a model package group and used in AWS Marketplace.
//
// For more information, see [`CreateModelPackage`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateModelPackage.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var modelInput interface{}
//
//   cfnModelPackagePropsMixin := awscdkmixinspreview.Mixins.NewCfnModelPackagePropsMixin(&CfnModelPackageMixinProps{
//   	AdditionalInferenceSpecifications: []interface{}{
//   		&AdditionalInferenceSpecificationDefinitionProperty{
//   			Containers: []interface{}{
//   				&ModelPackageContainerDefinitionProperty{
//   					ContainerHostname: jsii.String("containerHostname"),
//   					Environment: map[string]*string{
//   						"environmentKey": jsii.String("environment"),
//   					},
//   					Framework: jsii.String("framework"),
//   					FrameworkVersion: jsii.String("frameworkVersion"),
//   					Image: jsii.String("image"),
//   					ImageDigest: jsii.String("imageDigest"),
//   					ModelDataSource: &ModelDataSourceProperty{
//   						S3DataSource: &S3ModelDataSourceProperty{
//   							CompressionType: jsii.String("compressionType"),
//   							ModelAccessConfig: &ModelAccessConfigProperty{
//   								AcceptEula: jsii.Boolean(false),
//   							},
//   							S3DataType: jsii.String("s3DataType"),
//   							S3Uri: jsii.String("s3Uri"),
//   						},
//   					},
//   					ModelDataUrl: jsii.String("modelDataUrl"),
//   					ModelInput: modelInput,
//   					NearestModelName: jsii.String("nearestModelName"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			Name: jsii.String("name"),
//   			SupportedContentTypes: []*string{
//   				jsii.String("supportedContentTypes"),
//   			},
//   			SupportedRealtimeInferenceInstanceTypes: []*string{
//   				jsii.String("supportedRealtimeInferenceInstanceTypes"),
//   			},
//   			SupportedResponseMimeTypes: []*string{
//   				jsii.String("supportedResponseMimeTypes"),
//   			},
//   			SupportedTransformInstanceTypes: []*string{
//   				jsii.String("supportedTransformInstanceTypes"),
//   			},
//   		},
//   	},
//   	AdditionalInferenceSpecificationsToAdd: []interface{}{
//   		&AdditionalInferenceSpecificationDefinitionProperty{
//   			Containers: []interface{}{
//   				&ModelPackageContainerDefinitionProperty{
//   					ContainerHostname: jsii.String("containerHostname"),
//   					Environment: map[string]*string{
//   						"environmentKey": jsii.String("environment"),
//   					},
//   					Framework: jsii.String("framework"),
//   					FrameworkVersion: jsii.String("frameworkVersion"),
//   					Image: jsii.String("image"),
//   					ImageDigest: jsii.String("imageDigest"),
//   					ModelDataSource: &ModelDataSourceProperty{
//   						S3DataSource: &S3ModelDataSourceProperty{
//   							CompressionType: jsii.String("compressionType"),
//   							ModelAccessConfig: &ModelAccessConfigProperty{
//   								AcceptEula: jsii.Boolean(false),
//   							},
//   							S3DataType: jsii.String("s3DataType"),
//   							S3Uri: jsii.String("s3Uri"),
//   						},
//   					},
//   					ModelDataUrl: jsii.String("modelDataUrl"),
//   					ModelInput: modelInput,
//   					NearestModelName: jsii.String("nearestModelName"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			Name: jsii.String("name"),
//   			SupportedContentTypes: []*string{
//   				jsii.String("supportedContentTypes"),
//   			},
//   			SupportedRealtimeInferenceInstanceTypes: []*string{
//   				jsii.String("supportedRealtimeInferenceInstanceTypes"),
//   			},
//   			SupportedResponseMimeTypes: []*string{
//   				jsii.String("supportedResponseMimeTypes"),
//   			},
//   			SupportedTransformInstanceTypes: []*string{
//   				jsii.String("supportedTransformInstanceTypes"),
//   			},
//   		},
//   	},
//   	ApprovalDescription: jsii.String("approvalDescription"),
//   	CertifyForMarketplace: jsii.Boolean(false),
//   	ClientToken: jsii.String("clientToken"),
//   	CustomerMetadataProperties: map[string]*string{
//   		"customerMetadataPropertiesKey": jsii.String("customerMetadataProperties"),
//   	},
//   	Domain: jsii.String("domain"),
//   	DriftCheckBaselines: &DriftCheckBaselinesProperty{
//   		Bias: &DriftCheckBiasProperty{
//   			ConfigFile: &FileSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   			PostTrainingConstraints: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   			PreTrainingConstraints: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   		Explainability: &DriftCheckExplainabilityProperty{
//   			ConfigFile: &FileSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   			Constraints: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   		ModelDataQuality: &DriftCheckModelDataQualityProperty{
//   			Constraints: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   			Statistics: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   		ModelQuality: &DriftCheckModelQualityProperty{
//   			Constraints: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   			Statistics: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   	},
//   	InferenceSpecification: &InferenceSpecificationProperty{
//   		Containers: []interface{}{
//   			&ModelPackageContainerDefinitionProperty{
//   				ContainerHostname: jsii.String("containerHostname"),
//   				Environment: map[string]*string{
//   					"environmentKey": jsii.String("environment"),
//   				},
//   				Framework: jsii.String("framework"),
//   				FrameworkVersion: jsii.String("frameworkVersion"),
//   				Image: jsii.String("image"),
//   				ImageDigest: jsii.String("imageDigest"),
//   				ModelDataSource: &ModelDataSourceProperty{
//   					S3DataSource: &S3ModelDataSourceProperty{
//   						CompressionType: jsii.String("compressionType"),
//   						ModelAccessConfig: &ModelAccessConfigProperty{
//   							AcceptEula: jsii.Boolean(false),
//   						},
//   						S3DataType: jsii.String("s3DataType"),
//   						S3Uri: jsii.String("s3Uri"),
//   					},
//   				},
//   				ModelDataUrl: jsii.String("modelDataUrl"),
//   				ModelInput: modelInput,
//   				NearestModelName: jsii.String("nearestModelName"),
//   			},
//   		},
//   		SupportedContentTypes: []*string{
//   			jsii.String("supportedContentTypes"),
//   		},
//   		SupportedRealtimeInferenceInstanceTypes: []*string{
//   			jsii.String("supportedRealtimeInferenceInstanceTypes"),
//   		},
//   		SupportedResponseMimeTypes: []*string{
//   			jsii.String("supportedResponseMimeTypes"),
//   		},
//   		SupportedTransformInstanceTypes: []*string{
//   			jsii.String("supportedTransformInstanceTypes"),
//   		},
//   	},
//   	LastModifiedTime: jsii.String("lastModifiedTime"),
//   	MetadataProperties: &MetadataPropertiesProperty{
//   		CommitId: jsii.String("commitId"),
//   		GeneratedBy: jsii.String("generatedBy"),
//   		ProjectId: jsii.String("projectId"),
//   		Repository: jsii.String("repository"),
//   	},
//   	ModelApprovalStatus: jsii.String("modelApprovalStatus"),
//   	ModelCard: &ModelCardProperty{
//   		ModelCardContent: jsii.String("modelCardContent"),
//   		ModelCardStatus: jsii.String("modelCardStatus"),
//   	},
//   	ModelMetrics: &ModelMetricsProperty{
//   		Bias: &BiasProperty{
//   			PostTrainingReport: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   			PreTrainingReport: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   			Report: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   		Explainability: &ExplainabilityProperty{
//   			Report: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   		ModelDataQuality: &ModelDataQualityProperty{
//   			Constraints: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   			Statistics: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   		ModelQuality: &ModelQualityProperty{
//   			Constraints: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   			Statistics: &MetricsSourceProperty{
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//   	},
//   	ModelPackageDescription: jsii.String("modelPackageDescription"),
//   	ModelPackageGroupName: jsii.String("modelPackageGroupName"),
//   	ModelPackageName: jsii.String("modelPackageName"),
//   	ModelPackageStatusDetails: &ModelPackageStatusDetailsProperty{
//   		ValidationStatuses: []interface{}{
//   			&ModelPackageStatusItemProperty{
//   				FailureReason: jsii.String("failureReason"),
//   				Name: jsii.String("name"),
//   				Status: jsii.String("status"),
//   			},
//   		},
//   	},
//   	ModelPackageVersion: jsii.Number(123),
//   	SamplePayloadUrl: jsii.String("samplePayloadUrl"),
//   	SecurityConfig: &SecurityConfigProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	SkipModelValidation: jsii.String("skipModelValidation"),
//   	SourceAlgorithmSpecification: &SourceAlgorithmSpecificationProperty{
//   		SourceAlgorithms: []interface{}{
//   			&SourceAlgorithmProperty{
//   				AlgorithmName: jsii.String("algorithmName"),
//   				ModelDataUrl: jsii.String("modelDataUrl"),
//   			},
//   		},
//   	},
//   	SourceUri: jsii.String("sourceUri"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Task: jsii.String("task"),
//   	ValidationSpecification: &ValidationSpecificationProperty{
//   		ValidationProfiles: []interface{}{
//   			&ValidationProfileProperty{
//   				ProfileName: jsii.String("profileName"),
//   				TransformJobDefinition: &TransformJobDefinitionProperty{
//   					BatchStrategy: jsii.String("batchStrategy"),
//   					Environment: map[string]*string{
//   						"environmentKey": jsii.String("environment"),
//   					},
//   					MaxConcurrentTransforms: jsii.Number(123),
//   					MaxPayloadInMb: jsii.Number(123),
//   					TransformInput: &TransformInputProperty{
//   						CompressionType: jsii.String("compressionType"),
//   						ContentType: jsii.String("contentType"),
//   						DataSource: &DataSourceProperty{
//   							S3DataSource: &S3DataSourceProperty{
//   								S3DataType: jsii.String("s3DataType"),
//   								S3Uri: jsii.String("s3Uri"),
//   							},
//   						},
//   						SplitType: jsii.String("splitType"),
//   					},
//   					TransformOutput: &TransformOutputProperty{
//   						Accept: jsii.String("accept"),
//   						AssembleWith: jsii.String("assembleWith"),
//   						KmsKeyId: jsii.String("kmsKeyId"),
//   						S3OutputPath: jsii.String("s3OutputPath"),
//   					},
//   					TransformResources: &TransformResourcesProperty{
//   						InstanceCount: jsii.Number(123),
//   						InstanceType: jsii.String("instanceType"),
//   						VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   					},
//   				},
//   			},
//   		},
//   		ValidationRole: jsii.String("validationRole"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html
//
type CfnModelPackagePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnModelPackageMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnModelPackagePropsMixin
type jsiiProxy_CfnModelPackagePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnModelPackagePropsMixin) Props() *CfnModelPackageMixinProps {
	var returns *CfnModelPackageMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackagePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::ModelPackage`.
func NewCfnModelPackagePropsMixin(props *CfnModelPackageMixinProps, options *mixins.CfnPropertyMixinOptions) CfnModelPackagePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnModelPackagePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnModelPackagePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::ModelPackage`.
func NewCfnModelPackagePropsMixin_Override(c CfnModelPackagePropsMixin, props *CfnModelPackageMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnModelPackagePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnModelPackagePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModelPackagePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnModelPackagePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModelPackagePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnModelPackagePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

