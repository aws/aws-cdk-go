package awsdatabrew

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatabrew/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatabrew"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a new DataBrew job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnJob := awscdk.Aws_databrew.NewCfnJob(this, jsii.String("MyCfnJob"), &CfnJobProps{
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	DatabaseOutputs: []interface{}{
//   		&DatabaseOutputProperty{
//   			DatabaseOptions: &DatabaseTableOutputOptionsProperty{
//   				TableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				TempDirectory: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					BucketOwner: jsii.String("bucketOwner"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			GlueConnectionName: jsii.String("glueConnectionName"),
//
//   			// the properties below are optional
//   			DatabaseOutputMode: jsii.String("databaseOutputMode"),
//   		},
//   	},
//   	DataCatalogOutputs: []interface{}{
//   		&DataCatalogOutputProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			TableName: jsii.String("tableName"),
//
//   			// the properties below are optional
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseOptions: &DatabaseTableOutputOptionsProperty{
//   				TableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				TempDirectory: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					BucketOwner: jsii.String("bucketOwner"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			Overwrite: jsii.Boolean(false),
//   			S3Options: &S3TableOutputOptionsProperty{
//   				Location: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					BucketOwner: jsii.String("bucketOwner"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   		},
//   	},
//   	DatasetName: jsii.String("datasetName"),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	EncryptionMode: jsii.String("encryptionMode"),
//   	JobSample: &JobSampleProperty{
//   		Mode: jsii.String("mode"),
//   		Size: jsii.Number(123),
//   	},
//   	LogSubscription: jsii.String("logSubscription"),
//   	MaxCapacity: jsii.Number(123),
//   	MaxRetries: jsii.Number(123),
//   	OutputLocation: &OutputLocationProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		BucketOwner: jsii.String("bucketOwner"),
//   		Key: jsii.String("key"),
//   	},
//   	Outputs: []interface{}{
//   		&OutputProperty{
//   			Location: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				BucketOwner: jsii.String("bucketOwner"),
//   				Key: jsii.String("key"),
//   			},
//
//   			// the properties below are optional
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			Format: jsii.String("format"),
//   			FormatOptions: &OutputFormatOptionsProperty{
//   				Csv: &CsvOutputOptionsProperty{
//   					Delimiter: jsii.String("delimiter"),
//   				},
//   			},
//   			MaxOutputFiles: jsii.Number(123),
//   			Overwrite: jsii.Boolean(false),
//   			PartitionColumns: []*string{
//   				jsii.String("partitionColumns"),
//   			},
//   		},
//   	},
//   	ProfileConfiguration: &ProfileConfigurationProperty{
//   		ColumnStatisticsConfigurations: []interface{}{
//   			&ColumnStatisticsConfigurationProperty{
//   				Statistics: &StatisticsConfigurationProperty{
//   					IncludedStatistics: []*string{
//   						jsii.String("includedStatistics"),
//   					},
//   					Overrides: []interface{}{
//   						&StatisticOverrideProperty{
//   							Parameters: map[string]*string{
//   								"parametersKey": jsii.String("parameters"),
//   							},
//   							Statistic: jsii.String("statistic"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				Selectors: []interface{}{
//   					&ColumnSelectorProperty{
//   						Name: jsii.String("name"),
//   						Regex: jsii.String("regex"),
//   					},
//   				},
//   			},
//   		},
//   		DatasetStatisticsConfiguration: &StatisticsConfigurationProperty{
//   			IncludedStatistics: []*string{
//   				jsii.String("includedStatistics"),
//   			},
//   			Overrides: []interface{}{
//   				&StatisticOverrideProperty{
//   					Parameters: map[string]*string{
//   						"parametersKey": jsii.String("parameters"),
//   					},
//   					Statistic: jsii.String("statistic"),
//   				},
//   			},
//   		},
//   		EntityDetectorConfiguration: &EntityDetectorConfigurationProperty{
//   			EntityTypes: []*string{
//   				jsii.String("entityTypes"),
//   			},
//
//   			// the properties below are optional
//   			AllowedStatistics: &AllowedStatisticsProperty{
//   				Statistics: []*string{
//   					jsii.String("statistics"),
//   				},
//   			},
//   		},
//   		ProfileColumns: []interface{}{
//   			&ColumnSelectorProperty{
//   				Name: jsii.String("name"),
//   				Regex: jsii.String("regex"),
//   			},
//   		},
//   	},
//   	ProjectName: jsii.String("projectName"),
//   	Recipe: &RecipeProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Version: jsii.String("version"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Timeout: jsii.Number(123),
//   	ValidationConfigurations: []interface{}{
//   		&ValidationConfigurationProperty{
//   			RulesetArn: jsii.String("rulesetArn"),
//
//   			// the properties below are optional
//   			ValidationMode: jsii.String("validationMode"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html
//
type CfnJob interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsdatabrew.IJobRef
	awscdk.ITaggable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Represents a list of JDBC database output objects which defines the output destination for a DataBrew recipe job to write into.
	DatabaseOutputs() interface{}
	SetDatabaseOutputs(val interface{})
	// One or more artifacts that represent the AWS Glue Data Catalog output from running the job.
	DataCatalogOutputs() interface{}
	SetDataCatalogOutputs(val interface{})
	// A dataset that the job is to process.
	DatasetName() *string
	SetDatasetName(val *string)
	// The Amazon Resource Name (ARN) of an encryption key that is used to protect the job output.
	EncryptionKeyArn() *string
	SetEncryptionKeyArn(val *string)
	// The encryption mode for the job, which can be one of the following:.
	EncryptionMode() *string
	SetEncryptionMode(val *string)
	Env() *interfaces.ResourceEnvironment
	// A reference to a Job resource.
	JobRef() *interfacesawsdatabrew.JobReference
	// A sample configuration for profile jobs only, which determines the number of rows on which the profile job is run.
	JobSample() interface{}
	SetJobSample(val interface{})
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
	// The current status of Amazon CloudWatch logging for the job.
	LogSubscription() *string
	SetLogSubscription(val *string)
	// The maximum number of nodes that can be consumed when the job processes data.
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	// The maximum number of times to retry the job after a job run fails.
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	// The unique name of the job.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The location in Amazon S3 where the job writes its output.
	OutputLocation() interface{}
	SetOutputLocation(val interface{})
	// One or more artifacts that represent output from running the job.
	Outputs() interface{}
	SetOutputs(val interface{})
	// Configuration for profile jobs.
	ProfileConfiguration() interface{}
	SetProfileConfiguration(val interface{})
	// The name of the project that the job is associated with.
	ProjectName() *string
	SetProjectName(val *string)
	// A series of data transformation steps that the job runs.
	Recipe() interface{}
	SetRecipe(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of the role to be assumed for this job.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// Metadata tags that have been applied to the job.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// The job's timeout in minutes.
	Timeout() *float64
	SetTimeout(val *float64)
	// The job type of the job, which must be one of the following:.
	Type() *string
	SetType(val *string)
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
	// List of validation configurations that are applied to the profile job.
	ValidationConfigurations() interface{}
	SetValidationConfigurations(val interface{})
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnJob
type jsiiProxy_CfnJob struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsdatabrewIJobRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnJob) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) DatabaseOutputs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseOutputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) DataCatalogOutputs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCatalogOutputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) DatasetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) EncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) EncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) JobRef() *interfacesawsdatabrew.JobReference {
	var returns *interfacesawsdatabrew.JobReference
	_jsii_.Get(
		j,
		"jobRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) JobSample() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobSample",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) LogSubscription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) OutputLocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Outputs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) ProfileConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profileConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Recipe() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recipe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) ValidationConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationConfigurations",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataBrew::Job`.
func NewCfnJob(scope constructs.Construct, id *string, props *CfnJobProps) CfnJob {
	_init_.Initialize()

	if err := validateNewCfnJobParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnJob{}

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataBrew::Job`.
func NewCfnJob_Override(c CfnJob, scope constructs.Construct, id *string, props *CfnJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_databrew.CfnJob",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnJob)SetDatabaseOutputs(val interface{}) {
	if err := j.validateSetDatabaseOutputsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseOutputs",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetDataCatalogOutputs(val interface{}) {
	if err := j.validateSetDataCatalogOutputsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCatalogOutputs",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetDatasetName(val *string) {
	_jsii_.Set(
		j,
		"datasetName",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetEncryptionKeyArn(val *string) {
	_jsii_.Set(
		j,
		"encryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetEncryptionMode(val *string) {
	_jsii_.Set(
		j,
		"encryptionMode",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetJobSample(val interface{}) {
	if err := j.validateSetJobSampleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobSample",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetLogSubscription(val *string) {
	_jsii_.Set(
		j,
		"logSubscription",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetMaxCapacity(val *float64) {
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetOutputLocation(val interface{}) {
	if err := j.validateSetOutputLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputLocation",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetOutputs(val interface{}) {
	if err := j.validateSetOutputsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputs",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetProfileConfiguration(val interface{}) {
	if err := j.validateSetProfileConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetProjectName(val *string) {
	_jsii_.Set(
		j,
		"projectName",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetRecipe(val interface{}) {
	if err := j.validateSetRecipeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recipe",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetValidationConfigurations(val interface{}) {
	if err := j.validateSetValidationConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationConfigurations",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnJob_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJob_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnJob",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnJob.
func CfnJob_IsCfnJob(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJob_IsCfnJobParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnJob",
		"isCfnJob",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnJob_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJob_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnJob",
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
func CfnJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_databrew.CfnJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJob_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_databrew.CfnJob",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnJob) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnJob) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnJob) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnJob) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnJob) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnJob) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnJob) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnJob) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnJob) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnJob) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnJob) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnJob) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJob) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJob) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnJob) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnJob) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnJob) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnJob) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJob) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnJob) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

