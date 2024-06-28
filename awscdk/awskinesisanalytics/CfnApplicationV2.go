package awskinesisanalytics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon Kinesis Data Analytics application.
//
// For information about creating a Kinesis Data Analytics application, see [Creating an Application](https://docs.aws.amazon.com/managed-flink/latest/java/getting-started.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationV2 := awscdk.Aws_kinesisanalytics.NewCfnApplicationV2(this, jsii.String("MyCfnApplicationV2"), &CfnApplicationV2Props{
//   	RuntimeEnvironment: jsii.String("runtimeEnvironment"),
//   	ServiceExecutionRole: jsii.String("serviceExecutionRole"),
//
//   	// the properties below are optional
//   	ApplicationConfiguration: &ApplicationConfigurationProperty{
//   		ApplicationCodeConfiguration: &ApplicationCodeConfigurationProperty{
//   			CodeContent: &CodeContentProperty{
//   				S3ContentLocation: &S3ContentLocationProperty{
//   					BucketArn: jsii.String("bucketArn"),
//   					FileKey: jsii.String("fileKey"),
//
//   					// the properties below are optional
//   					ObjectVersion: jsii.String("objectVersion"),
//   				},
//   				TextContent: jsii.String("textContent"),
//   				ZipFileContent: jsii.String("zipFileContent"),
//   			},
//   			CodeContentType: jsii.String("codeContentType"),
//   		},
//   		ApplicationSnapshotConfiguration: &ApplicationSnapshotConfigurationProperty{
//   			SnapshotsEnabled: jsii.Boolean(false),
//   		},
//   		EnvironmentProperties: &EnvironmentPropertiesProperty{
//   			PropertyGroups: []interface{}{
//   				&PropertyGroupProperty{
//   					PropertyGroupId: jsii.String("propertyGroupId"),
//   					PropertyMap: map[string]*string{
//   						"propertyMapKey": jsii.String("propertyMap"),
//   					},
//   				},
//   			},
//   		},
//   		FlinkApplicationConfiguration: &FlinkApplicationConfigurationProperty{
//   			CheckpointConfiguration: &CheckpointConfigurationProperty{
//   				ConfigurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				CheckpointingEnabled: jsii.Boolean(false),
//   				CheckpointInterval: jsii.Number(123),
//   				MinPauseBetweenCheckpoints: jsii.Number(123),
//   			},
//   			MonitoringConfiguration: &MonitoringConfigurationProperty{
//   				ConfigurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				LogLevel: jsii.String("logLevel"),
//   				MetricsLevel: jsii.String("metricsLevel"),
//   			},
//   			ParallelismConfiguration: &ParallelismConfigurationProperty{
//   				ConfigurationType: jsii.String("configurationType"),
//
//   				// the properties below are optional
//   				AutoScalingEnabled: jsii.Boolean(false),
//   				Parallelism: jsii.Number(123),
//   				ParallelismPerKpu: jsii.Number(123),
//   			},
//   		},
//   		SqlApplicationConfiguration: &SqlApplicationConfigurationProperty{
//   			Inputs: []interface{}{
//   				&InputProperty{
//   					InputSchema: &InputSchemaProperty{
//   						RecordColumns: []interface{}{
//   							&RecordColumnProperty{
//   								Name: jsii.String("name"),
//   								SqlType: jsii.String("sqlType"),
//
//   								// the properties below are optional
//   								Mapping: jsii.String("mapping"),
//   							},
//   						},
//   						RecordFormat: &RecordFormatProperty{
//   							RecordFormatType: jsii.String("recordFormatType"),
//
//   							// the properties below are optional
//   							MappingParameters: &MappingParametersProperty{
//   								CsvMappingParameters: &CSVMappingParametersProperty{
//   									RecordColumnDelimiter: jsii.String("recordColumnDelimiter"),
//   									RecordRowDelimiter: jsii.String("recordRowDelimiter"),
//   								},
//   								JsonMappingParameters: &JSONMappingParametersProperty{
//   									RecordRowPath: jsii.String("recordRowPath"),
//   								},
//   							},
//   						},
//
//   						// the properties below are optional
//   						RecordEncoding: jsii.String("recordEncoding"),
//   					},
//   					NamePrefix: jsii.String("namePrefix"),
//
//   					// the properties below are optional
//   					InputParallelism: &InputParallelismProperty{
//   						Count: jsii.Number(123),
//   					},
//   					InputProcessingConfiguration: &InputProcessingConfigurationProperty{
//   						InputLambdaProcessor: &InputLambdaProcessorProperty{
//   							ResourceArn: jsii.String("resourceArn"),
//   						},
//   					},
//   					KinesisFirehoseInput: &KinesisFirehoseInputProperty{
//   						ResourceArn: jsii.String("resourceArn"),
//   					},
//   					KinesisStreamsInput: &KinesisStreamsInputProperty{
//   						ResourceArn: jsii.String("resourceArn"),
//   					},
//   				},
//   			},
//   		},
//   		VpcConfigurations: []interface{}{
//   			&VpcConfigurationProperty{
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				SubnetIds: []*string{
//   					jsii.String("subnetIds"),
//   				},
//   			},
//   		},
//   		ZeppelinApplicationConfiguration: &ZeppelinApplicationConfigurationProperty{
//   			CatalogConfiguration: &CatalogConfigurationProperty{
//   				GlueDataCatalogConfiguration: &GlueDataCatalogConfigurationProperty{
//   					DatabaseArn: jsii.String("databaseArn"),
//   				},
//   			},
//   			CustomArtifactsConfiguration: []interface{}{
//   				&CustomArtifactConfigurationProperty{
//   					ArtifactType: jsii.String("artifactType"),
//
//   					// the properties below are optional
//   					MavenReference: &MavenReferenceProperty{
//   						ArtifactId: jsii.String("artifactId"),
//   						GroupId: jsii.String("groupId"),
//   						Version: jsii.String("version"),
//   					},
//   					S3ContentLocation: &S3ContentLocationProperty{
//   						BucketArn: jsii.String("bucketArn"),
//   						FileKey: jsii.String("fileKey"),
//
//   						// the properties below are optional
//   						ObjectVersion: jsii.String("objectVersion"),
//   					},
//   				},
//   			},
//   			DeployAsApplicationConfiguration: &DeployAsApplicationConfigurationProperty{
//   				S3ContentLocation: &S3ContentBaseLocationProperty{
//   					BucketArn: jsii.String("bucketArn"),
//
//   					// the properties below are optional
//   					BasePath: jsii.String("basePath"),
//   				},
//   			},
//   			MonitoringConfiguration: &ZeppelinMonitoringConfigurationProperty{
//   				LogLevel: jsii.String("logLevel"),
//   			},
//   		},
//   	},
//   	ApplicationDescription: jsii.String("applicationDescription"),
//   	ApplicationMaintenanceConfiguration: &ApplicationMaintenanceConfigurationProperty{
//   		ApplicationMaintenanceWindowStartTime: jsii.String("applicationMaintenanceWindowStartTime"),
//   	},
//   	ApplicationMode: jsii.String("applicationMode"),
//   	ApplicationName: jsii.String("applicationName"),
//   	RunConfiguration: &RunConfigurationProperty{
//   		ApplicationRestoreConfiguration: &ApplicationRestoreConfigurationProperty{
//   			ApplicationRestoreType: jsii.String("applicationRestoreType"),
//
//   			// the properties below are optional
//   			SnapshotName: jsii.String("snapshotName"),
//   		},
//   		FlinkRunConfiguration: &FlinkRunConfigurationProperty{
//   			AllowNonRestoredState: jsii.Boolean(false),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-application.html
//
type CfnApplicationV2 interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// Use this parameter to configure the application.
	ApplicationConfiguration() interface{}
	SetApplicationConfiguration(val interface{})
	// The description of the application.
	ApplicationDescription() *string
	SetApplicationDescription(val *string)
	// Describes the maintenance configuration for the application.
	ApplicationMaintenanceConfiguration() interface{}
	SetApplicationMaintenanceConfiguration(val interface{})
	// To create a Kinesis Data Analytics Studio notebook, you must set the mode to `INTERACTIVE` .
	ApplicationMode() *string
	SetApplicationMode(val *string)
	// The name of the application.
	ApplicationName() *string
	SetApplicationName(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Describes the starting parameters for an Managed Service for Apache Flink application.
	RunConfiguration() interface{}
	SetRunConfiguration(val interface{})
	// The runtime environment for the application.
	RuntimeEnvironment() *string
	SetRuntimeEnvironment(val *string)
	// Specifies the IAM role that the application uses to access external resources.
	ServiceExecutionRole() *string
	SetServiceExecutionRole(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// A list of one or more tags to assign to the application.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
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

// The jsii proxy struct for CfnApplicationV2
type jsiiProxy_CfnApplicationV2 struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnApplicationV2) ApplicationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) ApplicationDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) ApplicationMaintenanceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationMaintenanceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) ApplicationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) RunConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) RuntimeEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) ServiceExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceExecutionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationV2) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnApplicationV2(scope constructs.Construct, id *string, props *CfnApplicationV2Props) CfnApplicationV2 {
	_init_.Initialize()

	if err := validateNewCfnApplicationV2Parameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationV2{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisanalytics.CfnApplicationV2",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnApplicationV2_Override(c CfnApplicationV2, scope constructs.Construct, id *string, props *CfnApplicationV2Props) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisanalytics.CfnApplicationV2",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApplicationV2)SetApplicationConfiguration(val interface{}) {
	if err := j.validateSetApplicationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationV2)SetApplicationDescription(val *string) {
	_jsii_.Set(
		j,
		"applicationDescription",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationV2)SetApplicationMaintenanceConfiguration(val interface{}) {
	if err := j.validateSetApplicationMaintenanceConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationMaintenanceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationV2)SetApplicationMode(val *string) {
	_jsii_.Set(
		j,
		"applicationMode",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationV2)SetApplicationName(val *string) {
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationV2)SetRunConfiguration(val interface{}) {
	if err := j.validateSetRunConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationV2)SetRuntimeEnvironment(val *string) {
	if err := j.validateSetRuntimeEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEnvironment",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationV2)SetServiceExecutionRole(val *string) {
	if err := j.validateSetServiceExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceExecutionRole",
		val,
	)
}

func (j *jsiiProxy_CfnApplicationV2)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnApplicationV2_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationV2_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisanalytics.CfnApplicationV2",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnApplicationV2_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationV2_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisanalytics.CfnApplicationV2",
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
func CfnApplicationV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisanalytics.CfnApplicationV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationV2_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisanalytics.CfnApplicationV2",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationV2) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApplicationV2) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplicationV2) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplicationV2) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApplicationV2) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApplicationV2) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApplicationV2) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApplicationV2) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApplicationV2) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnApplicationV2) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnApplicationV2) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApplicationV2) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplicationV2) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplicationV2) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApplicationV2) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplicationV2) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnApplicationV2) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnApplicationV2) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplicationV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplicationV2) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

