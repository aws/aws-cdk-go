package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CodeBuild::Project` resource configures how AWS CodeBuild builds your source code.
//
// For example, it tells CodeBuild where to get the source code and which build environment to use.
//
// > To unset or remove a project value via CFN, explicitly provide the attribute with value as empty input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProject := awscdk.Aws_codebuild.NewCfnProject(this, jsii.String("MyCfnProject"), &CfnProjectProps{
//   	Artifacts: &ArtifactsProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		ArtifactIdentifier: jsii.String("artifactIdentifier"),
//   		EncryptionDisabled: jsii.Boolean(false),
//   		Location: jsii.String("location"),
//   		Name: jsii.String("name"),
//   		NamespaceType: jsii.String("namespaceType"),
//   		OverrideArtifactName: jsii.Boolean(false),
//   		Packaging: jsii.String("packaging"),
//   		Path: jsii.String("path"),
//   	},
//   	Environment: &EnvironmentProperty{
//   		ComputeType: jsii.String("computeType"),
//   		Image: jsii.String("image"),
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Certificate: jsii.String("certificate"),
//   		EnvironmentVariables: []interface{}{
//   			&EnvironmentVariableProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//
//   				// the properties below are optional
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		Fleet: &ProjectFleetProperty{
//   			FleetArn: jsii.String("fleetArn"),
//   		},
//   		ImagePullCredentialsType: jsii.String("imagePullCredentialsType"),
//   		PrivilegedMode: jsii.Boolean(false),
//   		RegistryCredential: &RegistryCredentialProperty{
//   			Credential: jsii.String("credential"),
//   			CredentialProvider: jsii.String("credentialProvider"),
//   		},
//   	},
//   	ServiceRole: jsii.String("serviceRole"),
//   	Source: &SourceProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Auth: &SourceAuthProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Resource: jsii.String("resource"),
//   		},
//   		BuildSpec: jsii.String("buildSpec"),
//   		BuildStatusConfig: &BuildStatusConfigProperty{
//   			Context: jsii.String("context"),
//   			TargetUrl: jsii.String("targetUrl"),
//   		},
//   		GitCloneDepth: jsii.Number(123),
//   		GitSubmodulesConfig: &GitSubmodulesConfigProperty{
//   			FetchSubmodules: jsii.Boolean(false),
//   		},
//   		InsecureSsl: jsii.Boolean(false),
//   		Location: jsii.String("location"),
//   		ReportBuildStatus: jsii.Boolean(false),
//   		SourceIdentifier: jsii.String("sourceIdentifier"),
//   	},
//
//   	// the properties below are optional
//   	BadgeEnabled: jsii.Boolean(false),
//   	BuildBatchConfig: &ProjectBuildBatchConfigProperty{
//   		BatchReportMode: jsii.String("batchReportMode"),
//   		CombineArtifacts: jsii.Boolean(false),
//   		Restrictions: &BatchRestrictionsProperty{
//   			ComputeTypesAllowed: []*string{
//   				jsii.String("computeTypesAllowed"),
//   			},
//   			MaximumBuildsAllowed: jsii.Number(123),
//   		},
//   		ServiceRole: jsii.String("serviceRole"),
//   		TimeoutInMins: jsii.Number(123),
//   	},
//   	Cache: &ProjectCacheProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Location: jsii.String("location"),
//   		Modes: []*string{
//   			jsii.String("modes"),
//   		},
//   	},
//   	ConcurrentBuildLimit: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	EncryptionKey: jsii.String("encryptionKey"),
//   	FileSystemLocations: []interface{}{
//   		&ProjectFileSystemLocationProperty{
//   			Identifier: jsii.String("identifier"),
//   			Location: jsii.String("location"),
//   			MountPoint: jsii.String("mountPoint"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			MountOptions: jsii.String("mountOptions"),
//   		},
//   	},
//   	LogsConfig: &LogsConfigProperty{
//   		CloudWatchLogs: &CloudWatchLogsConfigProperty{
//   			Status: jsii.String("status"),
//
//   			// the properties below are optional
//   			GroupName: jsii.String("groupName"),
//   			StreamName: jsii.String("streamName"),
//   		},
//   		S3Logs: &S3LogsConfigProperty{
//   			Status: jsii.String("status"),
//
//   			// the properties below are optional
//   			EncryptionDisabled: jsii.Boolean(false),
//   			Location: jsii.String("location"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	QueuedTimeoutInMinutes: jsii.Number(123),
//   	ResourceAccessRole: jsii.String("resourceAccessRole"),
//   	SecondaryArtifacts: []interface{}{
//   		&ArtifactsProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			ArtifactIdentifier: jsii.String("artifactIdentifier"),
//   			EncryptionDisabled: jsii.Boolean(false),
//   			Location: jsii.String("location"),
//   			Name: jsii.String("name"),
//   			NamespaceType: jsii.String("namespaceType"),
//   			OverrideArtifactName: jsii.Boolean(false),
//   			Packaging: jsii.String("packaging"),
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	SecondarySources: []interface{}{
//   		&SourceProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Auth: &SourceAuthProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Resource: jsii.String("resource"),
//   			},
//   			BuildSpec: jsii.String("buildSpec"),
//   			BuildStatusConfig: &BuildStatusConfigProperty{
//   				Context: jsii.String("context"),
//   				TargetUrl: jsii.String("targetUrl"),
//   			},
//   			GitCloneDepth: jsii.Number(123),
//   			GitSubmodulesConfig: &GitSubmodulesConfigProperty{
//   				FetchSubmodules: jsii.Boolean(false),
//   			},
//   			InsecureSsl: jsii.Boolean(false),
//   			Location: jsii.String("location"),
//   			ReportBuildStatus: jsii.Boolean(false),
//   			SourceIdentifier: jsii.String("sourceIdentifier"),
//   		},
//   	},
//   	SecondarySourceVersions: []interface{}{
//   		&ProjectSourceVersionProperty{
//   			SourceIdentifier: jsii.String("sourceIdentifier"),
//
//   			// the properties below are optional
//   			SourceVersion: jsii.String("sourceVersion"),
//   		},
//   	},
//   	SourceVersion: jsii.String("sourceVersion"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeoutInMinutes: jsii.Number(123),
//   	Triggers: &ProjectTriggersProperty{
//   		BuildType: jsii.String("buildType"),
//   		FilterGroups: []interface{}{
//   			[]interface{}{
//   				&WebhookFilterProperty{
//   					Pattern: jsii.String("pattern"),
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					ExcludeMatchedPattern: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		Webhook: jsii.Boolean(false),
//   	},
//   	Visibility: jsii.String("visibility"),
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html
//
type CfnProject interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// `Artifacts` is a property of the [AWS::CodeBuild::Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies output settings for artifacts generated by an AWS CodeBuild build.
	Artifacts() interface{}
	SetArtifacts(val interface{})
	// The ARN of the AWS CodeBuild project, such as `arn:aws:codebuild:us-west-2:123456789012:project/myProjectName` .
	AttrArn() *string
	AttrId() *string
	// Indicates whether AWS CodeBuild generates a publicly accessible URL for your project's build badge.
	BadgeEnabled() interface{}
	SetBadgeEnabled(val interface{})
	// A `ProjectBuildBatchConfig` object that defines the batch build options for the project.
	BuildBatchConfig() interface{}
	SetBuildBatchConfig(val interface{})
	// Settings that AWS CodeBuild uses to store and reuse build dependencies.
	Cache() interface{}
	SetCache(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The maximum number of concurrent builds that are allowed for this project.
	ConcurrentBuildLimit() *float64
	SetConcurrentBuildLimit(val *float64)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description that makes the build project easy to identify.
	Description() *string
	SetDescription(val *string)
	// The AWS Key Management Service customer master key (CMK) to be used for encrypting the build output artifacts.
	EncryptionKey() *string
	SetEncryptionKey(val *string)
	// The build environment settings for the project, such as the environment type or the environment variables to use for the build environment.
	Environment() interface{}
	SetEnvironment(val interface{})
	// An array of `ProjectFileSystemLocation` objects for a CodeBuild build project.
	FileSystemLocations() interface{}
	SetFileSystemLocations(val interface{})
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
	// Information about logs for the build project.
	LogsConfig() interface{}
	SetLogsConfig(val interface{})
	// The name of the build project.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The number of minutes a build is allowed to be queued before it times out.
	QueuedTimeoutInMinutes() *float64
	SetQueuedTimeoutInMinutes(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds.
	ResourceAccessRole() *string
	SetResourceAccessRole(val *string)
	// A list of `Artifacts` objects.
	SecondaryArtifacts() interface{}
	SetSecondaryArtifacts(val interface{})
	// An array of `ProjectSource` objects.
	SecondarySources() interface{}
	SetSecondarySources(val interface{})
	// An array of `ProjectSourceVersion` objects.
	SecondarySourceVersions() interface{}
	SetSecondarySourceVersions(val interface{})
	// The ARN of the IAM role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
	ServiceRole() *string
	SetServiceRole(val *string)
	// The source code settings for the project, such as the source code's repository type and location.
	Source() interface{}
	SetSource(val interface{})
	// A version of the build input to be built for this project.
	SourceVersion() *string
	SetSourceVersion(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// An arbitrary set of tags (key-value pairs) for the AWS CodeBuild project.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
	// How long, in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait before timing out any related build that did not get marked as completed.
	TimeoutInMinutes() *float64
	SetTimeoutInMinutes(val *float64)
	// For an existing AWS CodeBuild build project that has its source code stored in a GitHub repository, enables AWS CodeBuild to begin automatically rebuilding the source code every time a code change is pushed to the repository.
	Triggers() interface{}
	SetTriggers(val interface{})
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
	// Specifies the visibility of the project's builds.
	//
	// Possible values are:.
	Visibility() *string
	SetVisibility(val *string)
	// `VpcConfig` specifies settings that enable AWS CodeBuild to access resources in an Amazon VPC.
	VpcConfig() interface{}
	SetVpcConfig(val interface{})
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

// The jsii proxy struct for CfnProject
type jsiiProxy_CfnProject struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnProject) Artifacts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"artifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) BadgeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"badgeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) BuildBatchConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"buildBatchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Cache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) ConcurrentBuildLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"concurrentBuildLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) EncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Environment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) FileSystemLocations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileSystemLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) LogsConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) QueuedTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queuedTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) ResourceAccessRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceAccessRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) SecondaryArtifacts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) SecondarySources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondarySources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) SecondarySourceVersions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondarySourceVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Source() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) SourceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) TimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Triggers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) VpcConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}


func NewCfnProject(scope constructs.Construct, id *string, props *CfnProjectProps) CfnProject {
	_init_.Initialize()

	if err := validateNewCfnProjectParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProject{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.CfnProject",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnProject_Override(c CfnProject, scope constructs.Construct, id *string, props *CfnProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.CfnProject",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnProject)SetArtifacts(val interface{}) {
	if err := j.validateSetArtifactsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"artifacts",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetBadgeEnabled(val interface{}) {
	if err := j.validateSetBadgeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"badgeEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetBuildBatchConfig(val interface{}) {
	if err := j.validateSetBuildBatchConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildBatchConfig",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetCache(val interface{}) {
	if err := j.validateSetCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cache",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetConcurrentBuildLimit(val *float64) {
	_jsii_.Set(
		j,
		"concurrentBuildLimit",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetEncryptionKey(val *string) {
	_jsii_.Set(
		j,
		"encryptionKey",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetEnvironment(val interface{}) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetFileSystemLocations(val interface{}) {
	if err := j.validateSetFileSystemLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemLocations",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetLogsConfig(val interface{}) {
	if err := j.validateSetLogsConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logsConfig",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetQueuedTimeoutInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"queuedTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetResourceAccessRole(val *string) {
	_jsii_.Set(
		j,
		"resourceAccessRole",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetSecondaryArtifacts(val interface{}) {
	if err := j.validateSetSecondaryArtifactsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryArtifacts",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetSecondarySources(val interface{}) {
	if err := j.validateSetSecondarySourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondarySources",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetSecondarySourceVersions(val interface{}) {
	if err := j.validateSetSecondarySourceVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondarySourceVersions",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetServiceRole(val *string) {
	if err := j.validateSetServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetSource(val interface{}) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetSourceVersion(val *string) {
	_jsii_.Set(
		j,
		"sourceVersion",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetTimeoutInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"timeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetTriggers(val interface{}) {
	if err := j.validateSetTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetVisibility(val *string) {
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

func (j *jsiiProxy_CfnProject)SetVpcConfig(val interface{}) {
	if err := j.validateSetVpcConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnProject_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProject_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.CfnProject",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnProject_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProject_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.CfnProject",
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
func CfnProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.CfnProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProject_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.CfnProject",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProject) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnProject) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnProject) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnProject) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnProject) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnProject) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnProject) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnProject) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnProject) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnProject) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnProject) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnProject) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnProject) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnProject) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnProject) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnProject) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

