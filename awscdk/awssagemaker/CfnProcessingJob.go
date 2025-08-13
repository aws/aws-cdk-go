package awssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Amazon SageMaker processing job that is used to analyze data and evaluate models.
//
// For more information, see [Process Data and Evaluate Models](https://docs.aws.amazon.com/sagemaker/latest/dg/processing-job.html) .
//
// Also, note the following details specific to processing jobs created using CloudFormation stacks:
//
// - When you delete a CloudFormation stack with a processing job resource, the processing job is stopped using the [StopProcessingJob](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_StopProcessingJob.html) API but not deleted. Any tags associated with the processing job are deleted using the [DeleteTags](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DeleteTags.html) API.
// - If any part of your CloudFormation stack deployment fails and a rollback initiates, processing jobs with a specified `ProcessingJobName` value might cause the stack to become stuck in a failed state. This occurs because during a rollback, CloudFormation attempts to recreate the stack resources. Processing job names must be unique, so when CloudFormation attempts to recreate a processing job using the already defined name, this results in an `AlreadyExists` error. To prevent this, we recommend that you don't specify the optional `ProcessingJobName` property, thereby allowing SageMaker to auto-generate a unique name for your processing job. This ensures successful stack rollbacks when necessary. If you must use custom job names, you have to manually modify the `ProcessingJobName` and redeploy the stack to recover from a failed rollback.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProcessingJob := awscdk.Aws_sagemaker.NewCfnProcessingJob(this, jsii.String("MyCfnProcessingJob"), &CfnProcessingJobProps{
//   	AppSpecification: &AppSpecificationProperty{
//   		ImageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		ContainerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		ContainerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   	},
//   	ProcessingResources: &ProcessingResourcesProperty{
//   		ClusterConfig: &ClusterConfigProperty{
//   			InstanceCount: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			VolumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	ExperimentConfig: &ExperimentConfigProperty{
//   		ExperimentName: jsii.String("experimentName"),
//   		RunName: jsii.String("runName"),
//   		TrialComponentDisplayName: jsii.String("trialComponentDisplayName"),
//   		TrialName: jsii.String("trialName"),
//   	},
//   	NetworkConfig: &NetworkConfigProperty{
//   		EnableInterContainerTrafficEncryption: jsii.Boolean(false),
//   		EnableNetworkIsolation: jsii.Boolean(false),
//   		VpcConfig: &VpcConfigProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	ProcessingInputs: []interface{}{
//   		&ProcessingInputsObjectProperty{
//   			InputName: jsii.String("inputName"),
//
//   			// the properties below are optional
//   			AppManaged: jsii.Boolean(false),
//   			DatasetDefinition: &DatasetDefinitionProperty{
//   				AthenaDatasetDefinition: &AthenaDatasetDefinitionProperty{
//   					Catalog: jsii.String("catalog"),
//   					Database: jsii.String("database"),
//   					OutputFormat: jsii.String("outputFormat"),
//   					OutputS3Uri: jsii.String("outputS3Uri"),
//   					QueryString: jsii.String("queryString"),
//
//   					// the properties below are optional
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					OutputCompression: jsii.String("outputCompression"),
//   					WorkGroup: jsii.String("workGroup"),
//   				},
//   				DataDistributionType: jsii.String("dataDistributionType"),
//   				InputMode: jsii.String("inputMode"),
//   				LocalPath: jsii.String("localPath"),
//   				RedshiftDatasetDefinition: &RedshiftDatasetDefinitionProperty{
//   					ClusterId: jsii.String("clusterId"),
//   					ClusterRoleArn: jsii.String("clusterRoleArn"),
//   					Database: jsii.String("database"),
//   					DbUser: jsii.String("dbUser"),
//   					OutputFormat: jsii.String("outputFormat"),
//   					OutputS3Uri: jsii.String("outputS3Uri"),
//   					QueryString: jsii.String("queryString"),
//
//   					// the properties below are optional
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					OutputCompression: jsii.String("outputCompression"),
//   				},
//   			},
//   			S3Input: &S3InputProperty{
//   				S3DataType: jsii.String("s3DataType"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				LocalPath: jsii.String("localPath"),
//   				S3CompressionType: jsii.String("s3CompressionType"),
//   				S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   				S3InputMode: jsii.String("s3InputMode"),
//   			},
//   		},
//   	},
//   	ProcessingJobName: jsii.String("processingJobName"),
//   	ProcessingOutputConfig: &ProcessingOutputConfigProperty{
//   		Outputs: []interface{}{
//   			&ProcessingOutputsObjectProperty{
//   				OutputName: jsii.String("outputName"),
//
//   				// the properties below are optional
//   				AppManaged: jsii.Boolean(false),
//   				FeatureStoreOutput: &FeatureStoreOutputProperty{
//   					FeatureGroupName: jsii.String("featureGroupName"),
//   				},
//   				S3Output: &S3OutputProperty{
//   					S3UploadMode: jsii.String("s3UploadMode"),
//   					S3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					LocalPath: jsii.String("localPath"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	StoppingCondition: &StoppingConditionProperty{
//   		MaxRuntimeInSeconds: jsii.Number(123),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-processingjob.html
//
type CfnProcessingJob interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggableV2
	// Configuration to run a processing job in a specified container image.
	AppSpecification() interface{}
	SetAppSpecification(val interface{})
	// The Amazon Resource Name (ARN) of the AutoML job associated with this processing job.
	AttrAutoMlJobArn() *string
	// The time the processing job was created.
	AttrCreationTime() *string
	// A string, up to one KB in size, that contains metadata from the processing container when the processing job exits.
	AttrExitMessage() *string
	// A string, up to one KB in size, that contains the reason a processing job failed, if it failed.
	AttrFailureReason() *string
	// The time the processing job was last modified.
	AttrLastModifiedTime() *string
	// The ARN of a monitoring schedule for an endpoint associated with this processing job.
	AttrMonitoringScheduleArn() *string
	// The time that the processing job ended.
	AttrProcessingEndTime() *string
	// The ARN of the processing job.
	AttrProcessingJobArn() *string
	// The status of the processing job.
	AttrProcessingJobStatus() *string
	// The time that the processing job started.
	AttrProcessingStartTime() *string
	// The ARN of the training job associated with this processing job.
	AttrTrainingJobArn() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Sets the environment variables in the Docker container.
	Environment() interface{}
	SetEnvironment(val interface{})
	// Associates a SageMaker job as a trial component with an experiment and trial.
	ExperimentConfig() interface{}
	SetExperimentConfig(val interface{})
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
	// Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.
	NetworkConfig() interface{}
	SetNetworkConfig(val interface{})
	// The tree node.
	Node() constructs.Node
	// List of input configurations for the processing job.
	ProcessingInputs() interface{}
	SetProcessingInputs(val interface{})
	// The name of the processing job.
	ProcessingJobName() *string
	SetProcessingJobName(val *string)
	// Contains information about the output location for the compiled model and the target device that the model runs on.
	ProcessingOutputConfig() interface{}
	SetProcessingOutputConfig(val interface{})
	// Identifies the resources, ML compute instances, and ML storage volumes to deploy for a processing job.
	ProcessingResources() interface{}
	SetProcessingResources(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ARN of the role used to create the processing job.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Configures conditions under which the processing job should be stopped, such as how long the processing job has been running.
	StoppingCondition() interface{}
	SetStoppingCondition(val interface{})
	// An array of key-value pairs.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
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

// The jsii proxy struct for CfnProcessingJob
type jsiiProxy_CfnProcessingJob struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnProcessingJob) AppSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) AttrAutoMlJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAutoMlJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) AttrExitMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrExitMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) AttrFailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFailureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) AttrLastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) AttrMonitoringScheduleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMonitoringScheduleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) AttrProcessingEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProcessingEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) AttrProcessingJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProcessingJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) AttrProcessingJobStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProcessingJobStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) AttrProcessingStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProcessingStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) AttrTrainingJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTrainingJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) Environment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) ExperimentConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"experimentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) NetworkConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) ProcessingInputs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processingInputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) ProcessingJobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processingJobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) ProcessingOutputConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processingOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) ProcessingResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processingResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) StoppingCondition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProcessingJob) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnProcessingJob(scope constructs.Construct, id *string, props *CfnProcessingJobProps) CfnProcessingJob {
	_init_.Initialize()

	if err := validateNewCfnProcessingJobParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProcessingJob{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sagemaker.CfnProcessingJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnProcessingJob_Override(c CfnProcessingJob, scope constructs.Construct, id *string, props *CfnProcessingJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sagemaker.CfnProcessingJob",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnProcessingJob)SetAppSpecification(val interface{}) {
	if err := j.validateSetAppSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnProcessingJob)SetEnvironment(val interface{}) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_CfnProcessingJob)SetExperimentConfig(val interface{}) {
	if err := j.validateSetExperimentConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"experimentConfig",
		val,
	)
}

func (j *jsiiProxy_CfnProcessingJob)SetNetworkConfig(val interface{}) {
	if err := j.validateSetNetworkConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConfig",
		val,
	)
}

func (j *jsiiProxy_CfnProcessingJob)SetProcessingInputs(val interface{}) {
	if err := j.validateSetProcessingInputsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processingInputs",
		val,
	)
}

func (j *jsiiProxy_CfnProcessingJob)SetProcessingJobName(val *string) {
	_jsii_.Set(
		j,
		"processingJobName",
		val,
	)
}

func (j *jsiiProxy_CfnProcessingJob)SetProcessingOutputConfig(val interface{}) {
	if err := j.validateSetProcessingOutputConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processingOutputConfig",
		val,
	)
}

func (j *jsiiProxy_CfnProcessingJob)SetProcessingResources(val interface{}) {
	if err := j.validateSetProcessingResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processingResources",
		val,
	)
}

func (j *jsiiProxy_CfnProcessingJob)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnProcessingJob)SetStoppingCondition(val interface{}) {
	if err := j.validateSetStoppingConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stoppingCondition",
		val,
	)
}

func (j *jsiiProxy_CfnProcessingJob)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnProcessingJob_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProcessingJob_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnProcessingJob",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnProcessingJob_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProcessingJob_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnProcessingJob",
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
func CfnProcessingJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProcessingJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnProcessingJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProcessingJob_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sagemaker.CfnProcessingJob",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProcessingJob) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnProcessingJob) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnProcessingJob) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnProcessingJob) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnProcessingJob) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnProcessingJob) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnProcessingJob) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnProcessingJob) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnProcessingJob) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnProcessingJob) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnProcessingJob) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnProcessingJob) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProcessingJob) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProcessingJob) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnProcessingJob) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnProcessingJob) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnProcessingJob) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnProcessingJob) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProcessingJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProcessingJob) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

