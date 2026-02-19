package awsemrserverless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsemrserverless/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsemrserverless"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::EMRServerless::Application` resource specifies an EMR Serverless application.
//
// An application uses open source analytics frameworks to run jobs that process data. To create an application, you must specify the release version for the open source framework version you want to use and the type of application you want, such as Apache Spark or Apache Hive. After you create an application, you can submit data processing jobs or interactive requests to it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationObjectProperty_ ConfigurationObjectProperty
//
//   cfnApplication := awscdk.Aws_emrserverless.NewCfnApplication(this, jsii.String("MyCfnApplication"), &CfnApplicationProps{
//   	ReleaseLabel: jsii.String("releaseLabel"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Architecture: jsii.String("architecture"),
//   	AutoStartConfiguration: &AutoStartConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	AutoStopConfiguration: &AutoStopConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		IdleTimeoutMinutes: jsii.Number(123),
//   	},
//   	IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   		IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   	},
//   	ImageConfiguration: &ImageConfigurationInputProperty{
//   		ImageUri: jsii.String("imageUri"),
//   	},
//   	InitialCapacity: []interface{}{
//   		&InitialCapacityConfigKeyValuePairProperty{
//   			Key: jsii.String("key"),
//   			Value: &InitialCapacityConfigProperty{
//   				WorkerConfiguration: &WorkerConfigurationProperty{
//   					Cpu: jsii.String("cpu"),
//   					Memory: jsii.String("memory"),
//
//   					// the properties below are optional
//   					Disk: jsii.String("disk"),
//   					DiskType: jsii.String("diskType"),
//   				},
//   				WorkerCount: jsii.Number(123),
//   			},
//   		},
//   	},
//   	InteractiveConfiguration: &InteractiveConfigurationProperty{
//   		LivyEndpointEnabled: jsii.Boolean(false),
//   		StudioEnabled: jsii.Boolean(false),
//   	},
//   	MaximumCapacity: &MaximumAllowedResourcesProperty{
//   		Cpu: jsii.String("cpu"),
//   		Memory: jsii.String("memory"),
//
//   		// the properties below are optional
//   		Disk: jsii.String("disk"),
//   	},
//   	MonitoringConfiguration: &MonitoringConfigurationProperty{
//   		CloudWatchLoggingConfiguration: &CloudWatchLoggingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   			LogTypeMap: []interface{}{
//   				&LogTypeMapKeyValuePairProperty{
//   					Key: jsii.String("key"),
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		ManagedPersistenceMonitoringConfiguration: &ManagedPersistenceMonitoringConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   		},
//   		PrometheusMonitoringConfiguration: &PrometheusMonitoringConfigurationProperty{
//   			RemoteWriteUrl: jsii.String("remoteWriteUrl"),
//   		},
//   		S3MonitoringConfiguration: &S3MonitoringConfigurationProperty{
//   			EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   			LogUri: jsii.String("logUri"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	RuntimeConfiguration: []interface{}{
//   		&ConfigurationObjectProperty{
//   			Classification: jsii.String("classification"),
//
//   			// the properties below are optional
//   			Configurations: []interface{}{
//   				configurationObjectProperty_,
//   			},
//   			Properties: map[string]*string{
//   				"propertiesKey": jsii.String("properties"),
//   			},
//   		},
//   	},
//   	SchedulerConfiguration: &SchedulerConfigurationProperty{
//   		MaxConcurrentRuns: jsii.Number(123),
//   		QueueTimeoutMinutes: jsii.Number(123),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkerTypeSpecifications: map[string]interface{}{
//   		"workerTypeSpecificationsKey": &WorkerTypeSpecificationInputProperty{
//   			"imageConfiguration": &ImageConfigurationInputProperty{
//   				"imageUri": jsii.String("imageUri"),
//   			},
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrserverless-application.html
//
type CfnApplication interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsemrserverless.IApplicationRef
	awscdk.ITaggable
	// A reference to a Application resource.
	ApplicationRef() *interfacesawsemrserverless.ApplicationReference
	// The CPU architecture of an application.
	Architecture() *string
	SetArchitecture(val *string)
	// The ID of the application, such as `ab4rp1abcs8xz47n3x0example` .
	AttrApplicationId() *string
	// The Amazon Resource Name (ARN) of the EMR Serverless Application.
	AttrArn() *string
	// The configuration for an application to automatically start on job submission.
	AutoStartConfiguration() interface{}
	SetAutoStartConfiguration(val interface{})
	// The configuration for an application to automatically stop after a certain amount of time being idle.
	AutoStopConfiguration() interface{}
	SetAutoStopConfiguration(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	Env() *interfaces.ResourceEnvironment
	// The IAM Identity Center configuration applied to enable trusted identity propagation.
	IdentityCenterConfiguration() interface{}
	SetIdentityCenterConfiguration(val interface{})
	// The image configuration applied to all worker types.
	ImageConfiguration() interface{}
	SetImageConfiguration(val interface{})
	// The initial capacity of the application.
	InitialCapacity() interface{}
	SetInitialCapacity(val interface{})
	// The interactive configuration object that enables the interactive use cases for an application.
	InteractiveConfiguration() interface{}
	SetInteractiveConfiguration(val interface{})
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
	// The maximum capacity of the application.
	MaximumCapacity() interface{}
	SetMaximumCapacity(val interface{})
	// A configuration specification to be used when provisioning an application.
	MonitoringConfiguration() interface{}
	SetMonitoringConfiguration(val interface{})
	// The name of the application.
	Name() *string
	SetName(val *string)
	// The network configuration for customer VPC connectivity for the application.
	NetworkConfiguration() interface{}
	SetNetworkConfiguration(val interface{})
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The EMR release associated with the application.
	ReleaseLabel() *string
	SetReleaseLabel(val *string)
	// The [Configuration](https://docs.aws.amazon.com/emr-serverless/latest/APIReference/API_Configuration.html) specifications of an application. Each configuration consists of a classification and properties. You use this parameter when creating or updating an application. To see the runtimeConfiguration object of an application, run the [GetApplication](https://docs.aws.amazon.com/emr-serverless/latest/APIReference/API_GetApplication.html) API operation.
	RuntimeConfiguration() interface{}
	SetRuntimeConfiguration(val interface{})
	// The scheduler configuration for batch and streaming jobs running on this application.
	SchedulerConfiguration() interface{}
	SetSchedulerConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The tags assigned to the application.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// The type of application, such as Spark or Hive.
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
	// The specification applied to each worker type.
	WorkerTypeSpecifications() interface{}
	SetWorkerTypeSpecifications(val interface{})
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

// The jsii proxy struct for CfnApplication
type jsiiProxy_CfnApplication struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsemrserverlessIApplicationRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnApplication) ApplicationRef() *interfacesawsemrserverless.ApplicationReference {
	var returns *interfacesawsemrserverless.ApplicationReference
	_jsii_.Get(
		j,
		"applicationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) AttrApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) AutoStartConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoStartConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) AutoStopConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoStopConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) IdentityCenterConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityCenterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) ImageConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) InitialCapacity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) InteractiveConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"interactiveConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) MaximumCapacity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maximumCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) MonitoringConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) NetworkConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) RuntimeConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) SchedulerConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedulerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) WorkerTypeSpecifications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workerTypeSpecifications",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMRServerless::Application`.
func NewCfnApplication(scope constructs.Construct, id *string, props *CfnApplicationProps) CfnApplication {
	_init_.Initialize()

	if err := validateNewCfnApplicationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplication{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emrserverless.CfnApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMRServerless::Application`.
func NewCfnApplication_Override(c CfnApplication, scope constructs.Construct, id *string, props *CfnApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emrserverless.CfnApplication",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApplication)SetArchitecture(val *string) {
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetAutoStartConfiguration(val interface{}) {
	if err := j.validateSetAutoStartConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoStartConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetAutoStopConfiguration(val interface{}) {
	if err := j.validateSetAutoStopConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoStopConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetIdentityCenterConfiguration(val interface{}) {
	if err := j.validateSetIdentityCenterConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityCenterConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetImageConfiguration(val interface{}) {
	if err := j.validateSetImageConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetInitialCapacity(val interface{}) {
	if err := j.validateSetInitialCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetInteractiveConfiguration(val interface{}) {
	if err := j.validateSetInteractiveConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interactiveConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetMaximumCapacity(val interface{}) {
	if err := j.validateSetMaximumCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetMonitoringConfiguration(val interface{}) {
	if err := j.validateSetMonitoringConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetNetworkConfiguration(val interface{}) {
	if err := j.validateSetNetworkConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetReleaseLabel(val *string) {
	if err := j.validateSetReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetRuntimeConfiguration(val interface{}) {
	if err := j.validateSetRuntimeConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetSchedulerConfiguration(val interface{}) {
	if err := j.validateSetSchedulerConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulerConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_CfnApplication)SetWorkerTypeSpecifications(val interface{}) {
	if err := j.validateSetWorkerTypeSpecificationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerTypeSpecifications",
		val,
	)
}

func CfnApplication_ArnForApplication(resource interfacesawsemrserverless.IApplicationRef) *string {
	_init_.Initialize()

	if err := validateCfnApplication_ArnForApplicationParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emrserverless.CfnApplication",
		"arnForApplication",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Creates a new IApplicationRef from an ARN.
func CfnApplication_FromApplicationArn(scope constructs.Construct, id *string, arn *string) interfacesawsemrserverless.IApplicationRef {
	_init_.Initialize()

	if err := validateCfnApplication_FromApplicationArnParameters(scope, id, arn); err != nil {
		panic(err)
	}
	var returns interfacesawsemrserverless.IApplicationRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emrserverless.CfnApplication",
		"fromApplicationArn",
		[]interface{}{scope, id, arn},
		&returns,
	)

	return returns
}

// Creates a new IApplicationRef from a applicationId.
func CfnApplication_FromApplicationId(scope constructs.Construct, id *string, applicationId *string) interfacesawsemrserverless.IApplicationRef {
	_init_.Initialize()

	if err := validateCfnApplication_FromApplicationIdParameters(scope, id, applicationId); err != nil {
		panic(err)
	}
	var returns interfacesawsemrserverless.IApplicationRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emrserverless.CfnApplication",
		"fromApplicationId",
		[]interface{}{scope, id, applicationId},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnApplication.
func CfnApplication_IsCfnApplication(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplication_IsCfnApplicationParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emrserverless.CfnApplication",
		"isCfnApplication",
		[]interface{}{x},
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
func CfnApplication_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplication_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emrserverless.CfnApplication",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnApplication_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplication_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emrserverless.CfnApplication",
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
func CfnApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emrserverless.CfnApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplication_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_emrserverless.CfnApplication",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplication) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApplication) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplication) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplication) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApplication) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApplication) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApplication) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApplication) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnApplication) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnApplication) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApplication) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApplication) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplication) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnApplication) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnApplication) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnApplication) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

