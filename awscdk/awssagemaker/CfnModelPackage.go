package awssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A versioned model that can be deployed for SageMaker inference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var modelInput interface{}
//
//   cfnModelPackage := awscdk.Aws_sagemaker.NewCfnModelPackage(this, jsii.String("MyCfnModelPackage"), &CfnModelPackageProps{
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
//   	},
//   	ModelPackageVersion: jsii.Number(123),
//   	SamplePayloadUrl: jsii.String("samplePayloadUrl"),
//   	SkipModelValidation: jsii.String("skipModelValidation"),
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
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackage.html
//
type CfnModelPackage interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// An array of additional Inference Specification objects.
	AdditionalInferenceSpecifications() interface{}
	SetAdditionalInferenceSpecifications(val interface{})
	// An array of additional Inference Specification objects to be added to the existing array.
	AdditionalInferenceSpecificationsToAdd() interface{}
	SetAdditionalInferenceSpecificationsToAdd(val interface{})
	// A description provided when the model approval is set.
	ApprovalDescription() *string
	SetApprovalDescription(val *string)
	// The time that the model package was created.
	AttrCreationTime() *string
	// The Amazon Resource Name (ARN) of the model package.
	AttrModelPackageArn() *string
	// The status of the model package. This can be one of the following values.
	//
	// - `PENDING` - The model package creation is pending.
	// - `IN_PROGRESS` - The model package is in the process of being created.
	// - `COMPLETED` - The model package was successfully created.
	// - `FAILED` - The model package creation failed.
	// - `DELETING` - The model package is in the process of being deleted.
	AttrModelPackageStatus() *string
	// Whether the model package is to be certified to be listed on AWS Marketplace.
	CertifyForMarketplace() interface{}
	SetCertifyForMarketplace(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A unique token that guarantees that the call to this API is idempotent.
	ClientToken() *string
	SetClientToken(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The metadata properties for the model package.
	CustomerMetadataProperties() interface{}
	SetCustomerMetadataProperties(val interface{})
	// The machine learning domain of your model package and its components.
	Domain() *string
	SetDomain(val *string)
	// Represents the drift check baselines that can be used when the model monitor is set using the model package.
	DriftCheckBaselines() interface{}
	SetDriftCheckBaselines(val interface{})
	// Defines how to perform inference generation after a training job is run.
	InferenceSpecification() interface{}
	SetInferenceSpecification(val interface{})
	// The last time the model package was modified.
	LastModifiedTime() *string
	SetLastModifiedTime(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// Metadata properties of the tracking entity, trial, or trial component.
	MetadataProperties() interface{}
	SetMetadataProperties(val interface{})
	// The approval status of the model.
	//
	// This can be one of the following values.
	ModelApprovalStatus() *string
	SetModelApprovalStatus(val *string)
	// Metrics for the model.
	ModelMetrics() interface{}
	SetModelMetrics(val interface{})
	// The description of the model package.
	ModelPackageDescription() *string
	SetModelPackageDescription(val *string)
	// The model group to which the model belongs.
	ModelPackageGroupName() *string
	SetModelPackageGroupName(val *string)
	// The name of the model.
	ModelPackageName() *string
	SetModelPackageName(val *string)
	// Specifies the validation and image scan statuses of the model package.
	ModelPackageStatusDetails() interface{}
	SetModelPackageStatusDetails(val interface{})
	// The version number of a versioned model.
	ModelPackageVersion() *float64
	SetModelPackageVersion(val *float64)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Simple Storage Service path where the sample payload are stored.
	SamplePayloadUrl() *string
	SetSamplePayloadUrl(val *string)
	// Indicates if you want to skip model validation.
	SkipModelValidation() *string
	SetSkipModelValidation(val *string)
	// A list of algorithms that were used to create a model package.
	SourceAlgorithmSpecification() interface{}
	SetSourceAlgorithmSpecification(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// A list of the tags associated with the model package.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// The machine learning task your model package accomplishes.
	Task() *string
	SetTask(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Specifies batch transform jobs that SageMaker runs to validate your model package.
	ValidationSpecification() interface{}
	SetValidationSpecification(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnModelPackage
type jsiiProxy_CfnModelPackage struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnModelPackage) AdditionalInferenceSpecifications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalInferenceSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) AdditionalInferenceSpecificationsToAdd() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalInferenceSpecificationsToAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) ApprovalDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvalDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) AttrModelPackageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModelPackageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) AttrModelPackageStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrModelPackageStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) CertifyForMarketplace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certifyForMarketplace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) CustomerMetadataProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerMetadataProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) DriftCheckBaselines() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"driftCheckBaselines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) InferenceSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) MetadataProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) ModelApprovalStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelApprovalStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) ModelMetrics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) ModelPackageDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) ModelPackageGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) ModelPackageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) ModelPackageStatusDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelPackageStatusDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) ModelPackageVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelPackageVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) SamplePayloadUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samplePayloadUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) SkipModelValidation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipModelValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) SourceAlgorithmSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceAlgorithmSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) Task() *string {
	var returns *string
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModelPackage) ValidationSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationSpecification",
		&returns,
	)
	return returns
}


func NewCfnModelPackage(scope constructs.Construct, id *string, props *CfnModelPackageProps) CfnModelPackage {
	_init_.Initialize()

	if err := validateNewCfnModelPackageParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnModelPackage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sagemaker.CfnModelPackage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnModelPackage_Override(c CfnModelPackage, scope constructs.Construct, id *string, props *CfnModelPackageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sagemaker.CfnModelPackage",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetAdditionalInferenceSpecifications(val interface{}) {
	if err := j.validateSetAdditionalInferenceSpecificationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalInferenceSpecifications",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetAdditionalInferenceSpecificationsToAdd(val interface{}) {
	if err := j.validateSetAdditionalInferenceSpecificationsToAddParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalInferenceSpecificationsToAdd",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetApprovalDescription(val *string) {
	_jsii_.Set(
		j,
		"approvalDescription",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetCertifyForMarketplace(val interface{}) {
	if err := j.validateSetCertifyForMarketplaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certifyForMarketplace",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetClientToken(val *string) {
	_jsii_.Set(
		j,
		"clientToken",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetCustomerMetadataProperties(val interface{}) {
	if err := j.validateSetCustomerMetadataPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerMetadataProperties",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetDriftCheckBaselines(val interface{}) {
	if err := j.validateSetDriftCheckBaselinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driftCheckBaselines",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetInferenceSpecification(val interface{}) {
	if err := j.validateSetInferenceSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetLastModifiedTime(val *string) {
	_jsii_.Set(
		j,
		"lastModifiedTime",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetMetadataProperties(val interface{}) {
	if err := j.validateSetMetadataPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataProperties",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetModelApprovalStatus(val *string) {
	_jsii_.Set(
		j,
		"modelApprovalStatus",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetModelMetrics(val interface{}) {
	if err := j.validateSetModelMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelMetrics",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetModelPackageDescription(val *string) {
	_jsii_.Set(
		j,
		"modelPackageDescription",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetModelPackageGroupName(val *string) {
	_jsii_.Set(
		j,
		"modelPackageGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetModelPackageName(val *string) {
	_jsii_.Set(
		j,
		"modelPackageName",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetModelPackageStatusDetails(val interface{}) {
	if err := j.validateSetModelPackageStatusDetailsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelPackageStatusDetails",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetModelPackageVersion(val *float64) {
	_jsii_.Set(
		j,
		"modelPackageVersion",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetSamplePayloadUrl(val *string) {
	_jsii_.Set(
		j,
		"samplePayloadUrl",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetSkipModelValidation(val *string) {
	_jsii_.Set(
		j,
		"skipModelValidation",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetSourceAlgorithmSpecification(val interface{}) {
	if err := j.validateSetSourceAlgorithmSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAlgorithmSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetTask(val *string) {
	_jsii_.Set(
		j,
		"task",
		val,
	)
}

func (j *jsiiProxy_CfnModelPackage)SetValidationSpecification(val interface{}) {
	if err := j.validateSetValidationSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationSpecification",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnModelPackage_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnModelPackage_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnModelPackage",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnModelPackage_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnModelPackage_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnModelPackage",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnModelPackage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnModelPackage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnModelPackage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModelPackage_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sagemaker.CfnModelPackage",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnModelPackage) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnModelPackage) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnModelPackage) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnModelPackage) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnModelPackage) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnModelPackage) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnModelPackage) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnModelPackage) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnModelPackage) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelPackage) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelPackage) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnModelPackage) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelPackage) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelPackage) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnModelPackage) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnModelPackage) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelPackage) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnModelPackage) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelPackage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnModelPackage) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

