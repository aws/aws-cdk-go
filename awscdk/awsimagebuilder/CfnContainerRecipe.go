package awsimagebuilder

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsimagebuilder"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new container recipe.
//
// Container recipes define how images are configured, tested, and assessed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContainerRecipe := awscdk.Aws_imagebuilder.NewCfnContainerRecipe(this, jsii.String("MyCfnContainerRecipe"), &CfnContainerRecipeProps{
//   	ContainerType: jsii.String("containerType"),
//   	Name: jsii.String("name"),
//   	ParentImage: jsii.String("parentImage"),
//   	TargetRepository: &TargetContainerRepositoryProperty{
//   		RepositoryName: jsii.String("repositoryName"),
//   		Service: jsii.String("service"),
//   	},
//   	Version: jsii.String("version"),
//
//   	// the properties below are optional
//   	Components: []interface{}{
//   		&ComponentConfigurationProperty{
//   			ComponentArn: jsii.String("componentArn"),
//   			Parameters: []interface{}{
//   				&ComponentParameterProperty{
//   					Name: jsii.String("name"),
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DockerfileTemplateData: jsii.String("dockerfileTemplateData"),
//   	DockerfileTemplateUri: jsii.String("dockerfileTemplateUri"),
//   	ImageOsVersionOverride: jsii.String("imageOsVersionOverride"),
//   	InstanceConfiguration: &InstanceConfigurationProperty{
//   		BlockDeviceMappings: []interface{}{
//   			&InstanceBlockDeviceMappingProperty{
//   				DeviceName: jsii.String("deviceName"),
//   				Ebs: &EbsInstanceBlockDeviceSpecificationProperty{
//   					DeleteOnTermination: jsii.Boolean(false),
//   					Encrypted: jsii.Boolean(false),
//   					Iops: jsii.Number(123),
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					SnapshotId: jsii.String("snapshotId"),
//   					Throughput: jsii.Number(123),
//   					VolumeSize: jsii.Number(123),
//   					VolumeType: jsii.String("volumeType"),
//   				},
//   				NoDevice: jsii.String("noDevice"),
//   				VirtualName: jsii.String("virtualName"),
//   			},
//   		},
//   		Image: jsii.String("image"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	PlatformOverride: jsii.String("platformOverride"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	WorkingDirectory: jsii.String("workingDirectory"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-containerrecipe.html
//
type CfnContainerRecipe interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsimagebuilder.IContainerRecipeRef
	awscdk.ITaggable
	// Returns the Amazon Resource Name (ARN) of the container recipe.
	//
	// For example, `arn:aws:imagebuilder:us-east-1:123456789012:container-recipe/mybasicrecipe/2020.12.17` .
	AttrArn() *string
	// The latest version references of the container recipe.
	AttrLatestVersion() awscdk.IResolvable
	// The Amazon Resource Name (ARN) of the container recipe.
	//
	// > Semantic versioning is included in each object's Amazon Resource Name (ARN), at the level that applies to that object as follows:
	// >
	// > - Versionless ARNs and Name ARNs do not include specific values in any of the nodes. The nodes are either left off entirely, or they are specified as wildcards, for example: x.x.x.
	// > - Version ARNs have only the first three nodes: <major>.<minor>.<patch>
	// > - Build version ARNs have all four nodes, and point to a specific build for a specific version of an object.
	AttrLatestVersionArn() *string
	// The latest version ARN of the created container recipe, with the same major version.
	AttrLatestVersionMajor() *string
	// The latest version ARN of the created container recipe, with the same minor version.
	AttrLatestVersionMinor() *string
	// The latest version ARN of the created container recipe, with the same patch version.
	AttrLatestVersionPatch() *string
	// Returns the name of the container recipe.
	AttrName() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Build and test components that are included in the container recipe.
	Components() interface{}
	SetComponents(val interface{})
	// A reference to a ContainerRecipe resource.
	ContainerRecipeRef() *interfacesawsimagebuilder.ContainerRecipeReference
	// Specifies the type of container, such as Docker.
	ContainerType() *string
	SetContainerType(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the container recipe.
	Description() *string
	SetDescription(val *string)
	// Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside.
	DockerfileTemplateData() *string
	SetDockerfileTemplateData(val *string)
	// The S3 URI for the Dockerfile that will be used to build your container image.
	DockerfileTemplateUri() *string
	SetDockerfileTemplateUri(val *string)
	Env() *interfaces.ResourceEnvironment
	// Specifies the operating system version for the base image.
	ImageOsVersionOverride() *string
	SetImageOsVersionOverride(val *string)
	// A group of options that can be used to configure an instance for building and testing container images.
	InstanceConfiguration() interface{}
	SetInstanceConfiguration(val interface{})
	// The Amazon Resource Name (ARN) that uniquely identifies which KMS key is used to encrypt the container image for distribution to the target Region.
	KmsKeyId() *string
	SetKmsKeyId(val *string)
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
	// The name of the container recipe.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The base image for customizations specified in the container recipe.
	ParentImage() *string
	SetParentImage(val *string)
	// Specifies the operating system platform when you use a custom base image.
	PlatformOverride() *string
	SetPlatformOverride(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// Tags that are attached to the container recipe.
	TagsRaw() *map[string]*string
	SetTagsRaw(val *map[string]*string)
	// The destination repository for the container image.
	TargetRepository() interface{}
	SetTargetRepository(val interface{})
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
	// The semantic version of the container recipe.
	Version() *string
	SetVersion(val *string)
	// The working directory for use during build and test workflows.
	WorkingDirectory() *string
	SetWorkingDirectory(val *string)
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

// The jsii proxy struct for CfnContainerRecipe
type jsiiProxy_CfnContainerRecipe struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsimagebuilderIContainerRecipeRef
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnContainerRecipe) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) AttrLatestVersion() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLatestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) AttrLatestVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLatestVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) AttrLatestVersionMajor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLatestVersionMajor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) AttrLatestVersionMinor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLatestVersionMinor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) AttrLatestVersionPatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLatestVersionPatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Components() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) ContainerRecipeRef() *interfacesawsimagebuilder.ContainerRecipeReference {
	var returns *interfacesawsimagebuilder.ContainerRecipeReference
	_jsii_.Get(
		j,
		"containerRecipeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) ContainerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) DockerfileTemplateData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfileTemplateData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) DockerfileTemplateUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfileTemplateUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) ImageOsVersionOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageOsVersionOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) InstanceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) ParentImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) PlatformOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) TagsRaw() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) TargetRepository() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnContainerRecipe) WorkingDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workingDirectory",
		&returns,
	)
	return returns
}


// Create a new `AWS::ImageBuilder::ContainerRecipe`.
func NewCfnContainerRecipe(scope constructs.Construct, id *string, props *CfnContainerRecipeProps) CfnContainerRecipe {
	_init_.Initialize()

	if err := validateNewCfnContainerRecipeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnContainerRecipe{}

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnContainerRecipe",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ImageBuilder::ContainerRecipe`.
func NewCfnContainerRecipe_Override(c CfnContainerRecipe, scope constructs.Construct, id *string, props *CfnContainerRecipeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_imagebuilder.CfnContainerRecipe",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetComponents(val interface{}) {
	if err := j.validateSetComponentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"components",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetContainerType(val *string) {
	if err := j.validateSetContainerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerType",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetDockerfileTemplateData(val *string) {
	_jsii_.Set(
		j,
		"dockerfileTemplateData",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetDockerfileTemplateUri(val *string) {
	_jsii_.Set(
		j,
		"dockerfileTemplateUri",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetImageOsVersionOverride(val *string) {
	_jsii_.Set(
		j,
		"imageOsVersionOverride",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetInstanceConfiguration(val interface{}) {
	if err := j.validateSetInstanceConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetParentImage(val *string) {
	if err := j.validateSetParentImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentImage",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetPlatformOverride(val *string) {
	_jsii_.Set(
		j,
		"platformOverride",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetTagsRaw(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetTargetRepository(val interface{}) {
	if err := j.validateSetTargetRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRepository",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_CfnContainerRecipe)SetWorkingDirectory(val *string) {
	_jsii_.Set(
		j,
		"workingDirectory",
		val,
	)
}

func CfnContainerRecipe_ArnForContainerRecipe(resource interfacesawsimagebuilder.IContainerRecipeRef) *string {
	_init_.Initialize()

	if err := validateCfnContainerRecipe_ArnForContainerRecipeParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnContainerRecipe",
		"arnForContainerRecipe",
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
func CfnContainerRecipe_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerRecipe_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnContainerRecipe",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnContainerRecipe_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerRecipe_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnContainerRecipe",
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
func CfnContainerRecipe_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnContainerRecipe_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_imagebuilder.CfnContainerRecipe",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnContainerRecipe_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_imagebuilder.CfnContainerRecipe",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnContainerRecipe) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnContainerRecipe) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnContainerRecipe) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerRecipe) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerRecipe) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnContainerRecipe) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnContainerRecipe) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerRecipe) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnContainerRecipe) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

