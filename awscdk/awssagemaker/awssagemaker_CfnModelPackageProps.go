package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnModelPackage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var environment interface{}
//   var modelInput interface{}
//
//   cfnModelPackageProps := &cfnModelPackageProps{
//   	additionalInferenceSpecificationDefinition: &additionalInferenceSpecificationDefinitionProperty{
//   		containers: []interface{}{
//   			&modelPackageContainerDefinitionProperty{
//   				image: jsii.String("image"),
//
//   				// the properties below are optional
//   				containerHostname: jsii.String("containerHostname"),
//   				environment: environment,
//   				framework: jsii.String("framework"),
//   				frameworkVersion: jsii.String("frameworkVersion"),
//   				imageDigest: jsii.String("imageDigest"),
//   				modelDataUrl: jsii.String("modelDataUrl"),
//   				modelInput: modelInput,
//   				nearestModelName: jsii.String("nearestModelName"),
//   				productId: jsii.String("productId"),
//   			},
//   		},
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		supportedContentTypes: []*string{
//   			jsii.String("supportedContentTypes"),
//   		},
//   		supportedRealtimeInferenceInstanceTypes: []*string{
//   			jsii.String("supportedRealtimeInferenceInstanceTypes"),
//   		},
//   		supportedResponseMimeTypes: []*string{
//   			jsii.String("supportedResponseMimeTypes"),
//   		},
//   		supportedTransformInstanceTypes: []*string{
//   			jsii.String("supportedTransformInstanceTypes"),
//   		},
//   	},
//   	additionalInferenceSpecifications: []interface{}{
//   		&additionalInferenceSpecificationDefinitionProperty{
//   			containers: []interface{}{
//   				&modelPackageContainerDefinitionProperty{
//   					image: jsii.String("image"),
//
//   					// the properties below are optional
//   					containerHostname: jsii.String("containerHostname"),
//   					environment: environment,
//   					framework: jsii.String("framework"),
//   					frameworkVersion: jsii.String("frameworkVersion"),
//   					imageDigest: jsii.String("imageDigest"),
//   					modelDataUrl: jsii.String("modelDataUrl"),
//   					modelInput: modelInput,
//   					nearestModelName: jsii.String("nearestModelName"),
//   					productId: jsii.String("productId"),
//   				},
//   			},
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   			supportedContentTypes: []*string{
//   				jsii.String("supportedContentTypes"),
//   			},
//   			supportedRealtimeInferenceInstanceTypes: []*string{
//   				jsii.String("supportedRealtimeInferenceInstanceTypes"),
//   			},
//   			supportedResponseMimeTypes: []*string{
//   				jsii.String("supportedResponseMimeTypes"),
//   			},
//   			supportedTransformInstanceTypes: []*string{
//   				jsii.String("supportedTransformInstanceTypes"),
//   			},
//   		},
//   	},
//   	additionalInferenceSpecificationsToAdd: []interface{}{
//   		&additionalInferenceSpecificationDefinitionProperty{
//   			containers: []interface{}{
//   				&modelPackageContainerDefinitionProperty{
//   					image: jsii.String("image"),
//
//   					// the properties below are optional
//   					containerHostname: jsii.String("containerHostname"),
//   					environment: environment,
//   					framework: jsii.String("framework"),
//   					frameworkVersion: jsii.String("frameworkVersion"),
//   					imageDigest: jsii.String("imageDigest"),
//   					modelDataUrl: jsii.String("modelDataUrl"),
//   					modelInput: modelInput,
//   					nearestModelName: jsii.String("nearestModelName"),
//   					productId: jsii.String("productId"),
//   				},
//   			},
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   			supportedContentTypes: []*string{
//   				jsii.String("supportedContentTypes"),
//   			},
//   			supportedRealtimeInferenceInstanceTypes: []*string{
//   				jsii.String("supportedRealtimeInferenceInstanceTypes"),
//   			},
//   			supportedResponseMimeTypes: []*string{
//   				jsii.String("supportedResponseMimeTypes"),
//   			},
//   			supportedTransformInstanceTypes: []*string{
//   				jsii.String("supportedTransformInstanceTypes"),
//   			},
//   		},
//   	},
//   	approvalDescription: jsii.String("approvalDescription"),
//   	certifyForMarketplace: jsii.Boolean(false),
//   	clientToken: jsii.String("clientToken"),
//   	createdBy: &userContextProperty{
//   		domainId: jsii.String("domainId"),
//   		userProfileArn: jsii.String("userProfileArn"),
//   		userProfileName: jsii.String("userProfileName"),
//   	},
//   	customerMetadataProperties: map[string]*string{
//   		"customerMetadataPropertiesKey": jsii.String("customerMetadataProperties"),
//   	},
//   	domain: jsii.String("domain"),
//   	driftCheckBaselines: &driftCheckBaselinesProperty{
//   		bias: &driftCheckBiasProperty{
//   			configFile: &fileSourceProperty{
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   				contentType: jsii.String("contentType"),
//   			},
//   			postTrainingConstraints: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   			preTrainingConstraints: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   		explainability: &driftCheckExplainabilityProperty{
//   			configFile: &fileSourceProperty{
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   				contentType: jsii.String("contentType"),
//   			},
//   			constraints: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   		modelDataQuality: &driftCheckModelDataQualityProperty{
//   			constraints: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   			statistics: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   		modelQuality: &driftCheckModelQualityProperty{
//   			constraints: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   			statistics: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	inferenceSpecification: &inferenceSpecificationProperty{
//   		containers: []interface{}{
//   			&modelPackageContainerDefinitionProperty{
//   				image: jsii.String("image"),
//
//   				// the properties below are optional
//   				containerHostname: jsii.String("containerHostname"),
//   				environment: environment,
//   				framework: jsii.String("framework"),
//   				frameworkVersion: jsii.String("frameworkVersion"),
//   				imageDigest: jsii.String("imageDigest"),
//   				modelDataUrl: jsii.String("modelDataUrl"),
//   				modelInput: modelInput,
//   				nearestModelName: jsii.String("nearestModelName"),
//   				productId: jsii.String("productId"),
//   			},
//   		},
//   		supportedContentTypes: []*string{
//   			jsii.String("supportedContentTypes"),
//   		},
//   		supportedResponseMimeTypes: []*string{
//   			jsii.String("supportedResponseMimeTypes"),
//   		},
//
//   		// the properties below are optional
//   		supportedRealtimeInferenceInstanceTypes: []*string{
//   			jsii.String("supportedRealtimeInferenceInstanceTypes"),
//   		},
//   		supportedTransformInstanceTypes: []*string{
//   			jsii.String("supportedTransformInstanceTypes"),
//   		},
//   	},
//   	lastModifiedBy: &userContextProperty{
//   		domainId: jsii.String("domainId"),
//   		userProfileArn: jsii.String("userProfileArn"),
//   		userProfileName: jsii.String("userProfileName"),
//   	},
//   	lastModifiedTime: jsii.String("lastModifiedTime"),
//   	metadataProperties: &metadataPropertiesProperty{
//   		commitId: jsii.String("commitId"),
//   		generatedBy: jsii.String("generatedBy"),
//   		projectId: jsii.String("projectId"),
//   		repository: jsii.String("repository"),
//   	},
//   	modelApprovalStatus: jsii.String("modelApprovalStatus"),
//   	modelMetrics: &modelMetricsProperty{
//   		bias: &biasProperty{
//   			postTrainingReport: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   			preTrainingReport: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   			report: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   		explainability: &explainabilityProperty{
//   			report: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   		modelDataQuality: &modelDataQualityProperty{
//   			constraints: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   			statistics: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   		modelQuality: &modelQualityProperty{
//   			constraints: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   			statistics: &metricsSourceProperty{
//   				contentType: jsii.String("contentType"),
//   				s3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				contentDigest: jsii.String("contentDigest"),
//   			},
//   		},
//   	},
//   	modelPackageDescription: jsii.String("modelPackageDescription"),
//   	modelPackageGroupName: jsii.String("modelPackageGroupName"),
//   	modelPackageName: jsii.String("modelPackageName"),
//   	modelPackageStatusDetails: &modelPackageStatusDetailsProperty{
//   		validationStatuses: []interface{}{
//   			&modelPackageStatusItemProperty{
//   				name: jsii.String("name"),
//   				status: jsii.String("status"),
//
//   				// the properties below are optional
//   				failureReason: jsii.String("failureReason"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		imageScanStatuses: []interface{}{
//   			&modelPackageStatusItemProperty{
//   				name: jsii.String("name"),
//   				status: jsii.String("status"),
//
//   				// the properties below are optional
//   				failureReason: jsii.String("failureReason"),
//   			},
//   		},
//   	},
//   	modelPackageStatusItem: &modelPackageStatusItemProperty{
//   		name: jsii.String("name"),
//   		status: jsii.String("status"),
//
//   		// the properties below are optional
//   		failureReason: jsii.String("failureReason"),
//   	},
//   	modelPackageVersion: jsii.Number(123),
//   	samplePayloadUrl: jsii.String("samplePayloadUrl"),
//   	sourceAlgorithmSpecification: &sourceAlgorithmSpecificationProperty{
//   		sourceAlgorithms: []interface{}{
//   			&sourceAlgorithmProperty{
//   				algorithmName: jsii.String("algorithmName"),
//
//   				// the properties below are optional
//   				modelDataUrl: jsii.String("modelDataUrl"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	task: jsii.String("task"),
//   	validationSpecification: &validationSpecificationProperty{
//   		validationProfiles: []interface{}{
//   			&validationProfileProperty{
//   				profileName: jsii.String("profileName"),
//   				transformJobDefinition: &transformJobDefinitionProperty{
//   					transformInput: &transformInputProperty{
//   						dataSource: &dataSourceProperty{
//   							s3DataSource: &s3DataSourceProperty{
//   								s3DataType: jsii.String("s3DataType"),
//   								s3Uri: jsii.String("s3Uri"),
//   							},
//   						},
//
//   						// the properties below are optional
//   						compressionType: jsii.String("compressionType"),
//   						contentType: jsii.String("contentType"),
//   						splitType: jsii.String("splitType"),
//   					},
//   					transformOutput: &transformOutputProperty{
//   						s3OutputPath: jsii.String("s3OutputPath"),
//
//   						// the properties below are optional
//   						accept: jsii.String("accept"),
//   						assembleWith: jsii.String("assembleWith"),
//   						kmsKeyId: jsii.String("kmsKeyId"),
//   					},
//   					transformResources: &transformResourcesProperty{
//   						instanceCount: jsii.Number(123),
//   						instanceType: jsii.String("instanceType"),
//
//   						// the properties below are optional
//   						volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   					},
//
//   					// the properties below are optional
//   					batchStrategy: jsii.String("batchStrategy"),
//   					environment: environment,
//   					maxConcurrentTransforms: jsii.Number(123),
//   					maxPayloadInMb: jsii.Number(123),
//   				},
//   			},
//   		},
//   		validationRole: jsii.String("validationRole"),
//   	},
//   }
//
type CfnModelPackageProps struct {
	// `AWS::SageMaker::ModelPackage.AdditionalInferenceSpecificationDefinition`.
	AdditionalInferenceSpecificationDefinition interface{} `field:"optional" json:"additionalInferenceSpecificationDefinition" yaml:"additionalInferenceSpecificationDefinition"`
	// An array of additional Inference Specification objects.
	AdditionalInferenceSpecifications interface{} `field:"optional" json:"additionalInferenceSpecifications" yaml:"additionalInferenceSpecifications"`
	// `AWS::SageMaker::ModelPackage.AdditionalInferenceSpecificationsToAdd`.
	AdditionalInferenceSpecificationsToAdd interface{} `field:"optional" json:"additionalInferenceSpecificationsToAdd" yaml:"additionalInferenceSpecificationsToAdd"`
	// A description provided when the model approval is set.
	ApprovalDescription *string `field:"optional" json:"approvalDescription" yaml:"approvalDescription"`
	// Whether the model package is to be certified to be listed on AWS Marketplace.
	//
	// For information about listing model packages on AWS Marketplace, see [List Your Algorithm or Model Package on AWS Marketplace](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-mkt-list.html) .
	CertifyForMarketplace interface{} `field:"optional" json:"certifyForMarketplace" yaml:"certifyForMarketplace"`
	// `AWS::SageMaker::ModelPackage.ClientToken`.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// `AWS::SageMaker::ModelPackage.CreatedBy`.
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
	// `AWS::SageMaker::ModelPackage.InferenceSpecification`.
	InferenceSpecification interface{} `field:"optional" json:"inferenceSpecification" yaml:"inferenceSpecification"`
	// `AWS::SageMaker::ModelPackage.LastModifiedBy`.
	LastModifiedBy interface{} `field:"optional" json:"lastModifiedBy" yaml:"lastModifiedBy"`
	// The last time the model package was modified.
	LastModifiedTime *string `field:"optional" json:"lastModifiedTime" yaml:"lastModifiedTime"`
	// `AWS::SageMaker::ModelPackage.MetadataProperties`.
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
	// `AWS::SageMaker::ModelPackage.ModelPackageStatusDetails`.
	ModelPackageStatusDetails interface{} `field:"optional" json:"modelPackageStatusDetails" yaml:"modelPackageStatusDetails"`
	// `AWS::SageMaker::ModelPackage.ModelPackageStatusItem`.
	ModelPackageStatusItem interface{} `field:"optional" json:"modelPackageStatusItem" yaml:"modelPackageStatusItem"`
	// The version number of a versioned model.
	ModelPackageVersion *float64 `field:"optional" json:"modelPackageVersion" yaml:"modelPackageVersion"`
	// The Amazon Simple Storage Service path where the sample payload are stored.
	//
	// This path must point to a single gzip compressed tar archive (.tar.gz suffix).
	SamplePayloadUrl *string `field:"optional" json:"samplePayloadUrl" yaml:"samplePayloadUrl"`
	// `AWS::SageMaker::ModelPackage.SourceAlgorithmSpecification`.
	SourceAlgorithmSpecification interface{} `field:"optional" json:"sourceAlgorithmSpecification" yaml:"sourceAlgorithmSpecification"`
	// A list of the tags associated with the model package.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The machine learning task your model package accomplishes.
	//
	// Common machine learning tasks include object detection and image classification.
	Task *string `field:"optional" json:"task" yaml:"task"`
	// `AWS::SageMaker::ModelPackage.ValidationSpecification`.
	ValidationSpecification interface{} `field:"optional" json:"validationSpecification" yaml:"validationSpecification"`
}

