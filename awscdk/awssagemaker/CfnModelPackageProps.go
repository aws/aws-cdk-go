package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnModelPackage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var modelInput interface{}
//
//   cfnModelPackageProps := &CfnModelPackageProps{
//   	AdditionalInferenceSpecificationDefinition: &AdditionalInferenceSpecificationDefinitionProperty{
//   		Containers: []interface{}{
//   			&ModelPackageContainerDefinitionProperty{
//   				Image: jsii.String("image"),
//
//   				// the properties below are optional
//   				ContainerHostname: jsii.String("containerHostname"),
//   				Environment: map[string]*string{
//   					"environmentKey": jsii.String("environment"),
//   				},
//   				Framework: jsii.String("framework"),
//   				FrameworkVersion: jsii.String("frameworkVersion"),
//   				ImageDigest: jsii.String("imageDigest"),
//   				ModelDataUrl: jsii.String("modelDataUrl"),
//   				ModelInput: modelInput,
//   				NearestModelName: jsii.String("nearestModelName"),
//   				ProductId: jsii.String("productId"),
//   			},
//   		},
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
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
//   	AdditionalInferenceSpecifications: []interface{}{
//   		&AdditionalInferenceSpecificationDefinitionProperty{
//   			Containers: []interface{}{
//   				&ModelPackageContainerDefinitionProperty{
//   					Image: jsii.String("image"),
//
//   					// the properties below are optional
//   					ContainerHostname: jsii.String("containerHostname"),
//   					Environment: map[string]*string{
//   						"environmentKey": jsii.String("environment"),
//   					},
//   					Framework: jsii.String("framework"),
//   					FrameworkVersion: jsii.String("frameworkVersion"),
//   					ImageDigest: jsii.String("imageDigest"),
//   					ModelDataUrl: jsii.String("modelDataUrl"),
//   					ModelInput: modelInput,
//   					NearestModelName: jsii.String("nearestModelName"),
//   					ProductId: jsii.String("productId"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
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
//   					Image: jsii.String("image"),
//
//   					// the properties below are optional
//   					ContainerHostname: jsii.String("containerHostname"),
//   					Environment: map[string]*string{
//   						"environmentKey": jsii.String("environment"),
//   					},
//   					Framework: jsii.String("framework"),
//   					FrameworkVersion: jsii.String("frameworkVersion"),
//   					ImageDigest: jsii.String("imageDigest"),
//   					ModelDataUrl: jsii.String("modelDataUrl"),
//   					ModelInput: modelInput,
//   					NearestModelName: jsii.String("nearestModelName"),
//   					ProductId: jsii.String("productId"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
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
//   	CreatedBy: &UserContextProperty{
//   		DomainId: jsii.String("domainId"),
//   		UserProfileArn: jsii.String("userProfileArn"),
//   		UserProfileName: jsii.String("userProfileName"),
//   	},
//   	CustomerMetadataProperties: map[string]*string{
//   		"customerMetadataPropertiesKey": jsii.String("customerMetadataProperties"),
//   	},
//   	Domain: jsii.String("domain"),
//   	DriftCheckBaselines: &DriftCheckBaselinesProperty{
//   		Bias: &DriftCheckBiasProperty{
//   			ConfigFile: &FileSourceProperty{
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   			},
//   			PostTrainingConstraints: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   			PreTrainingConstraints: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   		Explainability: &DriftCheckExplainabilityProperty{
//   			ConfigFile: &FileSourceProperty{
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   				ContentType: jsii.String("contentType"),
//   			},
//   			Constraints: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   		ModelDataQuality: &DriftCheckModelDataQualityProperty{
//   			Constraints: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   			Statistics: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   		ModelQuality: &DriftCheckModelQualityProperty{
//   			Constraints: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   			Statistics: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	InferenceSpecification: &InferenceSpecificationProperty{
//   		Containers: []interface{}{
//   			&ModelPackageContainerDefinitionProperty{
//   				Image: jsii.String("image"),
//
//   				// the properties below are optional
//   				ContainerHostname: jsii.String("containerHostname"),
//   				Environment: map[string]*string{
//   					"environmentKey": jsii.String("environment"),
//   				},
//   				Framework: jsii.String("framework"),
//   				FrameworkVersion: jsii.String("frameworkVersion"),
//   				ImageDigest: jsii.String("imageDigest"),
//   				ModelDataUrl: jsii.String("modelDataUrl"),
//   				ModelInput: modelInput,
//   				NearestModelName: jsii.String("nearestModelName"),
//   				ProductId: jsii.String("productId"),
//   			},
//   		},
//   		SupportedContentTypes: []*string{
//   			jsii.String("supportedContentTypes"),
//   		},
//   		SupportedResponseMimeTypes: []*string{
//   			jsii.String("supportedResponseMimeTypes"),
//   		},
//
//   		// the properties below are optional
//   		SupportedRealtimeInferenceInstanceTypes: []*string{
//   			jsii.String("supportedRealtimeInferenceInstanceTypes"),
//   		},
//   		SupportedTransformInstanceTypes: []*string{
//   			jsii.String("supportedTransformInstanceTypes"),
//   		},
//   	},
//   	LastModifiedBy: &UserContextProperty{
//   		DomainId: jsii.String("domainId"),
//   		UserProfileArn: jsii.String("userProfileArn"),
//   		UserProfileName: jsii.String("userProfileName"),
//   	},
//   	LastModifiedTime: jsii.String("lastModifiedTime"),
//   	MetadataProperties: &MetadataPropertiesProperty{
//   		CommitId: jsii.String("commitId"),
//   		GeneratedBy: jsii.String("generatedBy"),
//   		ProjectId: jsii.String("projectId"),
//   		Repository: jsii.String("repository"),
//   	},
//   	ModelApprovalStatus: jsii.String("modelApprovalStatus"),
//   	ModelMetrics: &ModelMetricsProperty{
//   		Bias: &BiasProperty{
//   			PostTrainingReport: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   			PreTrainingReport: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   			Report: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   		Explainability: &ExplainabilityProperty{
//   			Report: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   		ModelDataQuality: &ModelDataQualityProperty{
//   			Constraints: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   			Statistics: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   		ModelQuality: &ModelQualityProperty{
//   			Constraints: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   			Statistics: &MetricsSourceProperty{
//   				ContentType: jsii.String("contentType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				ContentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   	},
//   	ModelPackageDescription: jsii.String("modelPackageDescription"),
//   	ModelPackageGroupName: jsii.String("modelPackageGroupName"),
//   	ModelPackageName: jsii.String("modelPackageName"),
//   	ModelPackageStatusDetails: &ModelPackageStatusDetailsProperty{
//   		ValidationStatuses: []interface{}{
//   			&ModelPackageStatusItemProperty{
//   				Name: jsii.String("name"),
//   				Status: jsii.String("status"),
//
//   				// the properties below are optional
//   				FailureReason: jsii.String("failureReason"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		ImageScanStatuses: []interface{}{
//   			&ModelPackageStatusItemProperty{
//   				Name: jsii.String("name"),
//   				Status: jsii.String("status"),
//
//   				// the properties below are optional
//   				FailureReason: jsii.String("failureReason"),
//   			},
//   		},
//   	},
//   	ModelPackageStatusItem: &ModelPackageStatusItemProperty{
//   		Name: jsii.String("name"),
//   		Status: jsii.String("status"),
//
//   		// the properties below are optional
//   		FailureReason: jsii.String("failureReason"),
//   	},
//   	ModelPackageVersion: jsii.Number(123),
//   	SamplePayloadUrl: jsii.String("samplePayloadUrl"),
//   	SourceAlgorithmSpecification: &SourceAlgorithmSpecificationProperty{
//   		SourceAlgorithms: []interface{}{
//   			&SourceAlgorithmProperty{
//   				AlgorithmName: jsii.String("algorithmName"),
//
//   				// the properties below are optional
//   				ModelDataUrl: jsii.String("modelDataUrl"),
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
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
//   					TransformInput: &TransformInputProperty{
//   						DataSource: &DataSourceProperty{
//   							S3DataSource: &S3DataSourceProperty{
//   								S3DataType: jsii.String("s3DataType"),
//   								S3Uri: jsii.String("s3Uri"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						CompressionType: jsii.String("compressionType"),
//   						ContentType: jsii.String("contentType"),
//   						SplitType: jsii.String("splitType"),
//   					},
//   					TransformOutput: &TransformOutputProperty{
//   						S3OutputPath: jsii.String("s3OutputPath"),
//
//   						// the properties below are optional
//   						Accept: jsii.String("accept"),
//   						AssembleWith: jsii.String("assembleWith"),
//   						KmsKeyId: jsii.String("kmsKeyId"),
//   					},
//   					TransformResources: &TransformResourcesProperty{
//   						InstanceCount: jsii.Number(123),
//   						InstanceType: jsii.String("instanceType"),
//
//   						// the properties below are optional
//   						VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   					},
//
//   					// the properties below are optional
//   					BatchStrategy: jsii.String("batchStrategy"),
//   					Environment: map[string]*string{
//   						"environmentKey": jsii.String("environment"),
//   					},
//   					MaxConcurrentTransforms: jsii.Number(123),
//   					MaxPayloadInMb: jsii.Number(123),
//   				},
//   			},
//   		},
//   		ValidationRole: jsii.String("validationRole"),
//   	},
//   }
//
type CfnModelPackageProps struct {
	// A structure of additional Inference Specification.
	//
	// Additional Inference Specification specifies details about inference jobs that can be run with models based on this model package.
	AdditionalInferenceSpecificationDefinition interface{} `field:"optional" json:"additionalInferenceSpecificationDefinition" yaml:"additionalInferenceSpecificationDefinition"`
	// An array of additional Inference Specification objects.
	AdditionalInferenceSpecifications interface{} `field:"optional" json:"additionalInferenceSpecifications" yaml:"additionalInferenceSpecifications"`
	// An array of additional Inference Specification objects to be added to the existing array.
	//
	// The total number of additional Inference Specification objects cannot exceed 15. Each additional Inference Specification object specifies artifacts based on this model package that can be used on inference endpoints. Generally used with SageMaker Neo to store the compiled artifacts.
	AdditionalInferenceSpecificationsToAdd interface{} `field:"optional" json:"additionalInferenceSpecificationsToAdd" yaml:"additionalInferenceSpecificationsToAdd"`
	// A description provided when the model approval is set.
	ApprovalDescription *string `field:"optional" json:"approvalDescription" yaml:"approvalDescription"`
	// Whether the model package is to be certified to be listed on AWS Marketplace.
	//
	// For information about listing model packages on AWS Marketplace, see [List Your Algorithm or Model Package on AWS Marketplace](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-mkt-list.html) .
	CertifyForMarketplace interface{} `field:"optional" json:"certifyForMarketplace" yaml:"certifyForMarketplace"`
	// A unique token that guarantees that the call to this API is idempotent.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// Information about the user who created or modified an experiment, trial, trial component, lineage group, or project.
	CreatedBy interface{} `field:"optional" json:"createdBy" yaml:"createdBy"`
	// The metadata properties for the model package.
	CustomerMetadataProperties interface{} `field:"optional" json:"customerMetadataProperties" yaml:"customerMetadataProperties"`
	// The machine learning domain of your model package and its components.
	//
	// Common machine learning domains include computer vision and natural language processing.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Represents the drift check baselines that can be used when the model monitor is set using the model package.
	DriftCheckBaselines interface{} `field:"optional" json:"driftCheckBaselines" yaml:"driftCheckBaselines"`
	// The environment variables to set in the Docker container.
	//
	// Each key and value in the `Environment` string to string map can have length of up to 1024. We support up to 16 entries in the map.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Defines how to perform inference generation after a training job is run.
	InferenceSpecification interface{} `field:"optional" json:"inferenceSpecification" yaml:"inferenceSpecification"`
	// Information about the user who created or modified an experiment, trial, trial component, lineage group, or project.
	LastModifiedBy interface{} `field:"optional" json:"lastModifiedBy" yaml:"lastModifiedBy"`
	// The last time the model package was modified.
	LastModifiedTime *string `field:"optional" json:"lastModifiedTime" yaml:"lastModifiedTime"`
	// Metadata properties of the tracking entity, trial, or trial component.
	MetadataProperties interface{} `field:"optional" json:"metadataProperties" yaml:"metadataProperties"`
	// The approval status of the model. This can be one of the following values.
	//
	// - `APPROVED` - The model is approved
	// - `REJECTED` - The model is rejected.
	// - `PENDING_MANUAL_APPROVAL` - The model is waiting for manual approval.
	ModelApprovalStatus *string `field:"optional" json:"modelApprovalStatus" yaml:"modelApprovalStatus"`
	// Metrics for the model.
	ModelMetrics interface{} `field:"optional" json:"modelMetrics" yaml:"modelMetrics"`
	// The description of the model package.
	ModelPackageDescription *string `field:"optional" json:"modelPackageDescription" yaml:"modelPackageDescription"`
	// The model group to which the model belongs.
	ModelPackageGroupName *string `field:"optional" json:"modelPackageGroupName" yaml:"modelPackageGroupName"`
	// The name of the model.
	ModelPackageName *string `field:"optional" json:"modelPackageName" yaml:"modelPackageName"`
	// Specifies the validation and image scan statuses of the model package.
	ModelPackageStatusDetails interface{} `field:"optional" json:"modelPackageStatusDetails" yaml:"modelPackageStatusDetails"`
	// Represents the overall status of a model package.
	ModelPackageStatusItem interface{} `field:"optional" json:"modelPackageStatusItem" yaml:"modelPackageStatusItem"`
	// The version number of a versioned model.
	ModelPackageVersion *float64 `field:"optional" json:"modelPackageVersion" yaml:"modelPackageVersion"`
	// The Amazon Simple Storage Service path where the sample payload are stored.
	//
	// This path must point to a single gzip compressed tar archive (.tar.gz suffix).
	SamplePayloadUrl *string `field:"optional" json:"samplePayloadUrl" yaml:"samplePayloadUrl"`
	// A list of algorithms that were used to create a model package.
	SourceAlgorithmSpecification interface{} `field:"optional" json:"sourceAlgorithmSpecification" yaml:"sourceAlgorithmSpecification"`
	// A list of the tags associated with the model package.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The machine learning task your model package accomplishes.
	//
	// Common machine learning tasks include object detection and image classification.
	Task *string `field:"optional" json:"task" yaml:"task"`
	// Specifies batch transform jobs that SageMaker runs to validate your model package.
	ValidationSpecification interface{} `field:"optional" json:"validationSpecification" yaml:"validationSpecification"`
}

