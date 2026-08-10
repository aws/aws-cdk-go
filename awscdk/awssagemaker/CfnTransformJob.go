package awssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource type definition for AWS::SageMaker::TransformJob.
//
// A transform job uses a trained model to get inferences on a dataset and saves these results to an Amazon S3 location that you specify.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransformJob := awscdk.Aws_sagemaker.NewCfnTransformJob(this, jsii.String("MyCfnTransformJob"), &CfnTransformJobProps{
//   	ModelName: jsii.String("modelName"),
//   	TransformInput: &TransformInputProperty{
//   		DataSource: &DataSourceProperty{
//   			S3DataSource: &S3DataSourceProperty{
//   				S3DataType: jsii.String("s3DataType"),
//   				S3Uri: jsii.String("s3Uri"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		CompressionType: jsii.String("compressionType"),
//   		ContentType: jsii.String("contentType"),
//   		SplitType: jsii.String("splitType"),
//   	},
//   	TransformOutput: &TransformOutputProperty{
//   		S3OutputPath: jsii.String("s3OutputPath"),
//
//   		// the properties below are optional
//   		Accept: jsii.String("accept"),
//   		AssembleWith: jsii.String("assembleWith"),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	TransformResources: &TransformResourcesProperty{
//   		InstanceCount: jsii.Number(123),
//   		InstanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   	},
//
//   	// the properties below are optional
//   	BatchStrategy: jsii.String("batchStrategy"),
//   	DataCaptureConfig: &DataCaptureConfigProperty{
//   		DestinationS3Uri: jsii.String("destinationS3Uri"),
//
//   		// the properties below are optional
//   		GenerateInferenceId: jsii.Boolean(false),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	DataProcessing: &DataProcessingProperty{
//   		InputFilter: jsii.String("inputFilter"),
//   		JoinSource: jsii.String("joinSource"),
//   		OutputFilter: jsii.String("outputFilter"),
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	ExperimentConfig: &ExperimentConfigProperty{
//   		ExperimentName: jsii.String("experimentName"),
//   		TrialComponentDisplayName: jsii.String("trialComponentDisplayName"),
//   		TrialName: jsii.String("trialName"),
//   	},
//   	MaxConcurrentTransforms: jsii.Number(123),
//   	MaxPayloadInMb: jsii.Number(123),
//   	ModelClientConfig: &ModelClientConfigProperty{
//   		InvocationsMaxRetries: jsii.Number(123),
//   		InvocationsTimeoutInSeconds: jsii.Number(123),
//   	},
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-transformjob.html
//
type CfnTransformJob interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawssagemaker.ITransformJobRef
	awscdk.ITaggableV2
	// A timestamp that shows when the transform job was created.
	AttrCreationTime() *string
	// Indicates when the transform job has been completed, or has stopped or failed.
	AttrTransformEndTime() *string
	// The Amazon Resource Name (ARN) of the transform job.
	AttrTransformJobArn() *string
	// The name of the transform job.
	//
	// The name must be unique within an AWS Region in an AWS account.
	AttrTransformJobName() *string
	// The status of the transform job.
	AttrTransformJobStatus() *string
	// Indicates when the transform job starts on ML instances.
	AttrTransformStartTime() *string
	// Specifies the number of records to include in a mini-batch for an HTTP inference request.
	BatchStrategy() *string
	SetBatchStrategy(val *string)
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnPropertyNames() *map[string]*string
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Configuration to control how SageMaker captures inference data.
	DataCaptureConfig() interface{}
	SetDataCaptureConfig(val interface{})
	// The data structure used to specify the data to be used for inference in a batch transform job.
	DataProcessing() interface{}
	SetDataProcessing(val interface{})
	Env() *interfaces.ResourceEnvironment
	// The environment variables to set in the Docker container.
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
	// The maximum number of parallel requests that can be sent to each instance in a transform job.
	MaxConcurrentTransforms() *float64
	SetMaxConcurrentTransforms(val *float64)
	// The maximum allowed size of the payload, in MB.
	MaxPayloadInMb() *float64
	SetMaxPayloadInMb(val *float64)
	// Configures the timeout and maximum number of retries for processing a transform job invocation.
	ModelClientConfig() interface{}
	SetModelClientConfig(val interface{})
	// The name of the model that you want to use for the transform job.
	ModelName() *string
	SetModelName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs.
	Tags() *[]*CfnTransformJob_TagsItemsProperty
	SetTags(val *[]*CfnTransformJob_TagsItemsProperty)
	// Describes the input source and the way the transform job consumes it.
	TransformInput() interface{}
	SetTransformInput(val interface{})
	// A reference to a TransformJob resource.
	TransformJobRef() *interfacesawssagemaker.TransformJobReference
	// Describes the results of the transform job.
	TransformOutput() interface{}
	SetTransformOutput(val interface{})
	// Describes the resources, including ML instance types and ML instance count, to use for the transform job.
	TransformResources() interface{}
	SetTransformResources(val interface{})
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
	// This method has been renamed to `addResourceDependency` to more clearly
	// set it apart from `construct.node.addDependency`. See the documentation
	// of that function for more details.
	// Deprecated: Use `addResourceDependency` instead.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This method has been renamed to `addResourceDependency`, which makes it
	// more clear that this method operates at a different level from the
	// construct-level `construct.node.addDependency()` mechanism.
	// Deprecated: Use `addResourceDependency` instead.
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
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method only adds dependencies between L1 resources. If you are
	// looking for a generic construct-to-construct dependency mechanism that works
	// for all constructs including L2s, use `construct.node.addDependency` instead.
	AddResourceDependency(target awscdk.CfnResource, reason *string)
	// Sets the cross-stack reference strength for this resource.
	//
	// When set, any cross-stack reference to this resource will use the specified
	// strength instead of the global default from the consuming stack's context.
	ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength)
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
	CfnPropertyName(cdkPropertyName *string) *string
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	// Deprecated: Use `removeResourceDependency` instead.
	RemoveDependency(target awscdk.CfnResource)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveResourceDependency(target awscdk.CfnResource)
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnTransformJob
type jsiiProxy_CfnTransformJob struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawssagemakerITransformJobRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnTransformJob) AttrCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) AttrTransformEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTransformEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) AttrTransformJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTransformJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) AttrTransformJobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTransformJobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) AttrTransformJobStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTransformJobStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) AttrTransformStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTransformStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) BatchStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) CfnPropertyNames() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"cfnPropertyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) DataCaptureConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCaptureConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) DataProcessing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataProcessing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) Environment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) ExperimentConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"experimentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) MaxConcurrentTransforms() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentTransforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) MaxPayloadInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPayloadInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) ModelClientConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelClientConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) ModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) Tags() *[]*CfnTransformJob_TagsItemsProperty {
	var returns *[]*CfnTransformJob_TagsItemsProperty
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) TransformInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) TransformJobRef() *interfacesawssagemaker.TransformJobReference {
	var returns *interfacesawssagemaker.TransformJobReference
	_jsii_.Get(
		j,
		"transformJobRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) TransformOutput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) TransformResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransformJob) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::SageMaker::TransformJob`.
func NewCfnTransformJob(scope constructs.Construct, id *string, props *CfnTransformJobProps) CfnTransformJob {
	_init_.Initialize()

	if err := validateNewCfnTransformJobParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransformJob{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sagemaker.CfnTransformJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SageMaker::TransformJob`.
func NewCfnTransformJob_Override(c CfnTransformJob, scope constructs.Construct, id *string, props *CfnTransformJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sagemaker.CfnTransformJob",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTransformJob)SetBatchStrategy(val *string) {
	_jsii_.Set(
		j,
		"batchStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnTransformJob)SetDataCaptureConfig(val interface{}) {
	if err := j.validateSetDataCaptureConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCaptureConfig",
		val,
	)
}

func (j *jsiiProxy_CfnTransformJob)SetDataProcessing(val interface{}) {
	if err := j.validateSetDataProcessingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataProcessing",
		val,
	)
}

func (j *jsiiProxy_CfnTransformJob)SetEnvironment(val interface{}) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_CfnTransformJob)SetExperimentConfig(val interface{}) {
	if err := j.validateSetExperimentConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"experimentConfig",
		val,
	)
}

func (j *jsiiProxy_CfnTransformJob)SetMaxConcurrentTransforms(val *float64) {
	_jsii_.Set(
		j,
		"maxConcurrentTransforms",
		val,
	)
}

func (j *jsiiProxy_CfnTransformJob)SetMaxPayloadInMb(val *float64) {
	_jsii_.Set(
		j,
		"maxPayloadInMb",
		val,
	)
}

func (j *jsiiProxy_CfnTransformJob)SetModelClientConfig(val interface{}) {
	if err := j.validateSetModelClientConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelClientConfig",
		val,
	)
}

func (j *jsiiProxy_CfnTransformJob)SetModelName(val *string) {
	if err := j.validateSetModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelName",
		val,
	)
}

func (j *jsiiProxy_CfnTransformJob)SetTags(val *[]*CfnTransformJob_TagsItemsProperty) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnTransformJob)SetTransformInput(val interface{}) {
	if err := j.validateSetTransformInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformInput",
		val,
	)
}

func (j *jsiiProxy_CfnTransformJob)SetTransformOutput(val interface{}) {
	if err := j.validateSetTransformOutputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformOutput",
		val,
	)
}

func (j *jsiiProxy_CfnTransformJob)SetTransformResources(val interface{}) {
	if err := j.validateSetTransformResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformResources",
		val,
	)
}

func CfnTransformJob_ArnForTransformJob(resource interfacesawssagemaker.ITransformJobRef) *string {
	_init_.Initialize()

	if err := validateCfnTransformJob_ArnForTransformJobParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnTransformJob",
		"arnForTransformJob",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTransformJob_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransformJob_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnTransformJob",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnTransformJob_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransformJob_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnTransformJob",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnTransformJob.
func CfnTransformJob_IsCfnTransformJob(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransformJob_IsCfnTransformJobParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnTransformJob",
		"isCfnTransformJob",
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
func CfnTransformJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransformJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sagemaker.CfnTransformJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransformJob_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sagemaker.CfnTransformJob",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransformJob) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTransformJob) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTransformJob) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTransformJob) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTransformJob) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTransformJob) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTransformJob) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTransformJob) AddResourceDependency(target awscdk.CfnResource, reason *string) {
	if err := c.validateAddResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addResourceDependency",
		[]interface{}{target, reason},
	)
}

func (c *jsiiProxy_CfnTransformJob) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := c.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (c *jsiiProxy_CfnTransformJob) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTransformJob) CfnPropertyName(cdkPropertyName *string) *string {
	if err := c.validateCfnPropertyNameParameters(cdkPropertyName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"cfnPropertyName",
		[]interface{}{cdkPropertyName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTransformJob) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnTransformJob) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnTransformJob) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTransformJob) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTransformJob) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTransformJob) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTransformJob) RemoveResourceDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeResourceDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTransformJob) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnTransformJob) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnTransformJob) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTransformJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTransformJob) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnTransformJob) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

