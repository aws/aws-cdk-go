package previewawssagemakermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnModelPackagePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var modelInput interface{}
//
//   cfnModelPackageMixinProps := &CfnModelPackageMixinProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html
//
type CfnModelPackageMixinProps struct {
	// An array of additional Inference Specification objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-additionalinferencespecifications
	//
	AdditionalInferenceSpecifications interface{} `field:"optional" json:"additionalInferenceSpecifications" yaml:"additionalInferenceSpecifications"`
	// An array of additional Inference Specification objects to be added to the existing array.
	//
	// The total number of additional Inference Specification objects cannot exceed 15. Each additional Inference Specification object specifies artifacts based on this model package that can be used on inference endpoints. Generally used with SageMaker Neo to store the compiled artifacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-additionalinferencespecificationstoadd
	//
	AdditionalInferenceSpecificationsToAdd interface{} `field:"optional" json:"additionalInferenceSpecificationsToAdd" yaml:"additionalInferenceSpecificationsToAdd"`
	// A description provided when the model approval is set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-approvaldescription
	//
	ApprovalDescription *string `field:"optional" json:"approvalDescription" yaml:"approvalDescription"`
	// Whether the model package is to be certified to be listed on AWS Marketplace.
	//
	// For information about listing model packages on AWS Marketplace, see [List Your Algorithm or Model Package on AWS Marketplace](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-mkt-list.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-certifyformarketplace
	//
	CertifyForMarketplace interface{} `field:"optional" json:"certifyForMarketplace" yaml:"certifyForMarketplace"`
	// A unique token that guarantees that the call to this API is idempotent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// The metadata properties for the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-customermetadataproperties
	//
	CustomerMetadataProperties interface{} `field:"optional" json:"customerMetadataProperties" yaml:"customerMetadataProperties"`
	// The machine learning domain of your model package and its components.
	//
	// Common machine learning domains include computer vision and natural language processing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Represents the drift check baselines that can be used when the model monitor is set using the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-driftcheckbaselines
	//
	DriftCheckBaselines interface{} `field:"optional" json:"driftCheckBaselines" yaml:"driftCheckBaselines"`
	// Defines how to perform inference generation after a training job is run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-inferencespecification
	//
	InferenceSpecification interface{} `field:"optional" json:"inferenceSpecification" yaml:"inferenceSpecification"`
	// The last time the model package was modified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-lastmodifiedtime
	//
	LastModifiedTime *string `field:"optional" json:"lastModifiedTime" yaml:"lastModifiedTime"`
	// Metadata properties of the tracking entity, trial, or trial component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-metadataproperties
	//
	MetadataProperties interface{} `field:"optional" json:"metadataProperties" yaml:"metadataProperties"`
	// The approval status of the model. This can be one of the following values.
	//
	// - `APPROVED` - The model is approved
	// - `REJECTED` - The model is rejected.
	// - `PENDING_MANUAL_APPROVAL` - The model is waiting for manual approval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-modelapprovalstatus
	//
	ModelApprovalStatus *string `field:"optional" json:"modelApprovalStatus" yaml:"modelApprovalStatus"`
	// An Amazon SageMaker Model Card.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-modelcard
	//
	ModelCard interface{} `field:"optional" json:"modelCard" yaml:"modelCard"`
	// Metrics for the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-modelmetrics
	//
	ModelMetrics interface{} `field:"optional" json:"modelMetrics" yaml:"modelMetrics"`
	// The description of the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-modelpackagedescription
	//
	ModelPackageDescription *string `field:"optional" json:"modelPackageDescription" yaml:"modelPackageDescription"`
	// The model group to which the model belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-modelpackagegroupname
	//
	ModelPackageGroupName *string `field:"optional" json:"modelPackageGroupName" yaml:"modelPackageGroupName"`
	// The name of the model package. The name can be as follows:.
	//
	// - For a versioned model, the name is automatically generated by SageMaker Model Registry and follows the format ' `ModelPackageGroupName/ModelPackageVersion` '.
	// - For an unversioned model, you must provide the name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-modelpackagename
	//
	ModelPackageName *string `field:"optional" json:"modelPackageName" yaml:"modelPackageName"`
	// Specifies the validation and image scan statuses of the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-modelpackagestatusdetails
	//
	ModelPackageStatusDetails interface{} `field:"optional" json:"modelPackageStatusDetails" yaml:"modelPackageStatusDetails"`
	// The version number of a versioned model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-modelpackageversion
	//
	ModelPackageVersion *float64 `field:"optional" json:"modelPackageVersion" yaml:"modelPackageVersion"`
	// The Amazon Simple Storage Service path where the sample payload are stored.
	//
	// This path must point to a single gzip compressed tar archive (.tar.gz suffix).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-samplepayloadurl
	//
	SamplePayloadUrl *string `field:"optional" json:"samplePayloadUrl" yaml:"samplePayloadUrl"`
	// An optional AWS Key Management Service key to encrypt, decrypt, and re-encrypt model package information for regulated workloads with highly sensitive data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-securityconfig
	//
	SecurityConfig interface{} `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// Indicates if you want to skip model validation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-skipmodelvalidation
	//
	SkipModelValidation *string `field:"optional" json:"skipModelValidation" yaml:"skipModelValidation"`
	// A list of algorithms that were used to create a model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-sourcealgorithmspecification
	//
	SourceAlgorithmSpecification interface{} `field:"optional" json:"sourceAlgorithmSpecification" yaml:"sourceAlgorithmSpecification"`
	// The URI of the source for the model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-sourceuri
	//
	SourceUri *string `field:"optional" json:"sourceUri" yaml:"sourceUri"`
	// A list of the tags associated with the model package.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The machine learning task your model package accomplishes.
	//
	// Common machine learning tasks include object detection and image classification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-task
	//
	Task *string `field:"optional" json:"task" yaml:"task"`
	// Specifies batch transform jobs that SageMaker runs to validate your model package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html#cfn-sagemaker-modelpackage-validationspecification
	//
	ValidationSpecification interface{} `field:"optional" json:"validationSpecification" yaml:"validationSpecification"`
}

