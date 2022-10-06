package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::CodeBuild::Project`.
//
// The `AWS::CodeBuild::Project` resource configures how AWS CodeBuild builds your source code. For example, it tells CodeBuild where to get the source code and which build environment to use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProject := awscdk.Aws_codebuild.NewCfnProject(this, jsii.String("MyCfnProject"), &cfnProjectProps{
//   	artifacts: &artifactsProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		artifactIdentifier: jsii.String("artifactIdentifier"),
//   		encryptionDisabled: jsii.Boolean(false),
//   		location: jsii.String("location"),
//   		name: jsii.String("name"),
//   		namespaceType: jsii.String("namespaceType"),
//   		overrideArtifactName: jsii.Boolean(false),
//   		packaging: jsii.String("packaging"),
//   		path: jsii.String("path"),
//   	},
//   	environment: &environmentProperty{
//   		computeType: jsii.String("computeType"),
//   		image: jsii.String("image"),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		certificate: jsii.String("certificate"),
//   		environmentVariables: []interface{}{
//   			&environmentVariableProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//
//   				// the properties below are optional
//   				type: jsii.String("type"),
//   			},
//   		},
//   		imagePullCredentialsType: jsii.String("imagePullCredentialsType"),
//   		privilegedMode: jsii.Boolean(false),
//   		registryCredential: &registryCredentialProperty{
//   			credential: jsii.String("credential"),
//   			credentialProvider: jsii.String("credentialProvider"),
//   		},
//   	},
//   	serviceRole: jsii.String("serviceRole"),
//   	source: &sourceProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		auth: &sourceAuthProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			resource: jsii.String("resource"),
//   		},
//   		buildSpec: jsii.String("buildSpec"),
//   		buildStatusConfig: &buildStatusConfigProperty{
//   			context: jsii.String("context"),
//   			targetUrl: jsii.String("targetUrl"),
//   		},
//   		gitCloneDepth: jsii.Number(123),
//   		gitSubmodulesConfig: &gitSubmodulesConfigProperty{
//   			fetchSubmodules: jsii.Boolean(false),
//   		},
//   		insecureSsl: jsii.Boolean(false),
//   		location: jsii.String("location"),
//   		reportBuildStatus: jsii.Boolean(false),
//   		sourceIdentifier: jsii.String("sourceIdentifier"),
//   	},
//
//   	// the properties below are optional
//   	badgeEnabled: jsii.Boolean(false),
//   	buildBatchConfig: &projectBuildBatchConfigProperty{
//   		batchReportMode: jsii.String("batchReportMode"),
//   		combineArtifacts: jsii.Boolean(false),
//   		restrictions: &batchRestrictionsProperty{
//   			computeTypesAllowed: []*string{
//   				jsii.String("computeTypesAllowed"),
//   			},
//   			maximumBuildsAllowed: jsii.Number(123),
//   		},
//   		serviceRole: jsii.String("serviceRole"),
//   		timeoutInMins: jsii.Number(123),
//   	},
//   	cache: &projectCacheProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		location: jsii.String("location"),
//   		modes: []*string{
//   			jsii.String("modes"),
//   		},
//   	},
//   	concurrentBuildLimit: jsii.Number(123),
//   	description: jsii.String("description"),
//   	encryptionKey: jsii.String("encryptionKey"),
//   	fileSystemLocations: []interface{}{
//   		&projectFileSystemLocationProperty{
//   			identifier: jsii.String("identifier"),
//   			location: jsii.String("location"),
//   			mountPoint: jsii.String("mountPoint"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			mountOptions: jsii.String("mountOptions"),
//   		},
//   	},
//   	logsConfig: &logsConfigProperty{
//   		cloudWatchLogs: &cloudWatchLogsConfigProperty{
//   			status: jsii.String("status"),
//
//   			// the properties below are optional
//   			groupName: jsii.String("groupName"),
//   			streamName: jsii.String("streamName"),
//   		},
//   		s3Logs: &s3LogsConfigProperty{
//   			status: jsii.String("status"),
//
//   			// the properties below are optional
//   			encryptionDisabled: jsii.Boolean(false),
//   			location: jsii.String("location"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	queuedTimeoutInMinutes: jsii.Number(123),
//   	resourceAccessRole: jsii.String("resourceAccessRole"),
//   	secondaryArtifacts: []interface{}{
//   		&artifactsProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			artifactIdentifier: jsii.String("artifactIdentifier"),
//   			encryptionDisabled: jsii.Boolean(false),
//   			location: jsii.String("location"),
//   			name: jsii.String("name"),
//   			namespaceType: jsii.String("namespaceType"),
//   			overrideArtifactName: jsii.Boolean(false),
//   			packaging: jsii.String("packaging"),
//   			path: jsii.String("path"),
//   		},
//   	},
//   	secondarySources: []interface{}{
//   		&sourceProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			auth: &sourceAuthProperty{
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				resource: jsii.String("resource"),
//   			},
//   			buildSpec: jsii.String("buildSpec"),
//   			buildStatusConfig: &buildStatusConfigProperty{
//   				context: jsii.String("context"),
//   				targetUrl: jsii.String("targetUrl"),
//   			},
//   			gitCloneDepth: jsii.Number(123),
//   			gitSubmodulesConfig: &gitSubmodulesConfigProperty{
//   				fetchSubmodules: jsii.Boolean(false),
//   			},
//   			insecureSsl: jsii.Boolean(false),
//   			location: jsii.String("location"),
//   			reportBuildStatus: jsii.Boolean(false),
//   			sourceIdentifier: jsii.String("sourceIdentifier"),
//   		},
//   	},
//   	secondarySourceVersions: []interface{}{
//   		&projectSourceVersionProperty{
//   			sourceIdentifier: jsii.String("sourceIdentifier"),
//
//   			// the properties below are optional
//   			sourceVersion: jsii.String("sourceVersion"),
//   		},
//   	},
//   	sourceVersion: jsii.String("sourceVersion"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeoutInMinutes: jsii.Number(123),
//   	triggers: &projectTriggersProperty{
//   		buildType: jsii.String("buildType"),
//   		filterGroups: []interface{}{
//   			[]interface{}{
//   				&webhookFilterProperty{
//   					pattern: jsii.String("pattern"),
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					excludeMatchedPattern: jsii.Boolean(false),
//   				},
//   			},
//   		},
//   		webhook: jsii.Boolean(false),
//   	},
//   	visibility: jsii.String("visibility"),
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   		vpcId: jsii.String("vpcId"),
//   	},
//   })
//
type CfnProject interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `Artifacts` is a property of the [AWS::CodeBuild::Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies output settings for artifacts generated by an AWS CodeBuild build.
	Artifacts() interface{}
	SetArtifacts(val interface{})
	// The ARN of the AWS CodeBuild project, such as `arn:aws:codebuild:us-west-2:123456789012:project/myProjectName` .
	AttrArn() *string
	// Indicates whether AWS CodeBuild generates a publicly accessible URL for your project's build badge.
	//
	// For more information, see [Build Badges Sample](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-badges.html) in the *AWS CodeBuild User Guide* .
	//
	// > Including build badges with your project is currently not supported if the source type is CodePipeline. If you specify `CODEPIPELINE` for the `Source` property, do not specify the `BadgeEnabled` property.
	BadgeEnabled() interface{}
	SetBadgeEnabled(val interface{})
	// A `ProjectBuildBatchConfig` object that defines the batch build options for the project.
	BuildBatchConfig() interface{}
	SetBuildBatchConfig(val interface{})
	// Settings that AWS CodeBuild uses to store and reuse build dependencies.
	Cache() interface{}
	SetCache(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The maximum number of concurrent builds that are allowed for this project.
	//
	// New builds are only started if the current number of builds is less than or equal to this limit. If the current build count meets this limit, new builds are throttled and are not run.
	ConcurrentBuildLimit() *float64
	SetConcurrentBuildLimit(val *float64)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// A description that makes the build project easy to identify.
	Description() *string
	SetDescription(val *string)
	// The AWS Key Management Service customer master key (CMK) to be used for encrypting the build output artifacts.
	//
	// > You can use a cross-account KMS key to encrypt the build output artifacts if your service role has permission to that key.
	//
	// You can specify either the Amazon Resource Name (ARN) of the CMK or, if available, the CMK's alias (using the format `alias/<alias-name>` ). If you don't specify a value, CodeBuild uses the managed CMK for Amazon Simple Storage Service (Amazon S3).
	EncryptionKey() *string
	SetEncryptionKey(val *string)
	// The build environment settings for the project, such as the environment type or the environment variables to use for the build environment.
	Environment() interface{}
	SetEnvironment(val interface{})
	// An array of `ProjectFileSystemLocation` objects for a CodeBuild build project.
	//
	// A `ProjectFileSystemLocation` object specifies the `identifier` , `location` , `mountOptions` , `mountPoint` , and `type` of a file system created using Amazon Elastic File System.
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
	// Experimental.
	LogicalId() *string
	// Information about logs for the build project.
	//
	// A project can create logs in CloudWatch Logs, an S3 bucket, or both.
	LogsConfig() interface{}
	SetLogsConfig(val interface{})
	// The name of the build project.
	//
	// The name must be unique across all of the projects in your AWS account .
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The number of minutes a build is allowed to be queued before it times out.
	QueuedTimeoutInMinutes() *float64
	SetQueuedTimeoutInMinutes(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds.
	ResourceAccessRole() *string
	SetResourceAccessRole(val *string)
	// A list of `Artifacts` objects.
	//
	// Each artifacts object specifies output settings that the project generates during a build.
	SecondaryArtifacts() interface{}
	SetSecondaryArtifacts(val interface{})
	// An array of `ProjectSource` objects.
	SecondarySources() interface{}
	SetSecondarySources(val interface{})
	// An array of `ProjectSourceVersion` objects.
	//
	// If `secondarySourceVersions` is specified at the build level, then they take over these `secondarySourceVersions` (at the project level).
	SecondarySourceVersions() interface{}
	SetSecondarySourceVersions(val interface{})
	// The ARN of the IAM role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
	ServiceRole() *string
	SetServiceRole(val *string)
	// The source code settings for the project, such as the source code's repository type and location.
	Source() interface{}
	SetSource(val interface{})
	// A version of the build input to be built for this project.
	//
	// If not specified, the latest version is used. If specified, it must be one of:
	//
	// - For CodeCommit: the commit ID, branch, or Git tag to use.
	// - For GitHub: the commit ID, pull request ID, branch name, or tag name that corresponds to the version of the source code you want to build. If a pull request ID is specified, it must use the format `pr/pull-request-ID` (for example `pr/25` ). If a branch name is specified, the branch's HEAD commit ID is used. If not specified, the default branch's HEAD commit ID is used.
	// - For Bitbucket: the commit ID, branch name, or tag name that corresponds to the version of the source code you want to build. If a branch name is specified, the branch's HEAD commit ID is used. If not specified, the default branch's HEAD commit ID is used.
	// - For Amazon S3: the version ID of the object that represents the build input ZIP file to use.
	//
	// If `sourceVersion` is specified at the build level, then that version takes precedence over this `sourceVersion` (at the project level).
	//
	// For more information, see [Source Version Sample with CodeBuild](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-source-version.html) in the *AWS CodeBuild User Guide* .
	SourceVersion() *string
	SetSourceVersion(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// An arbitrary set of tags (key-value pairs) for the AWS CodeBuild project.
	//
	// These tags are available for use by AWS services that support AWS CodeBuild build project tags.
	Tags() awscdk.TagManager
	// How long, in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait before timing out any related build that did not get marked as completed.
	//
	// The default is 60 minutes.
	TimeoutInMinutes() *float64
	SetTimeoutInMinutes(val *float64)
	// For an existing AWS CodeBuild build project that has its source code stored in a GitHub repository, enables AWS CodeBuild to begin automatically rebuilding the source code every time a code change is pushed to the repository.
	Triggers() interface{}
	SetTriggers(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Specifies the visibility of the project's builds. Possible values are:.
	//
	// - **PUBLIC_READ** - The project builds are visible to the public.
	// - **PRIVATE** - The project builds are not visible to the public.
	Visibility() *string
	SetVisibility(val *string)
	// `VpcConfig` specifies settings that enable AWS CodeBuild to access resources in an Amazon VPC.
	//
	// For more information, see [Use AWS CodeBuild with Amazon Virtual Private Cloud](https://docs.aws.amazon.com/codebuild/latest/userguide/vpc-support.html) in the *AWS CodeBuild User Guide* .
	VpcConfig() interface{}
	SetVpcConfig(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnProject
type jsiiProxy_CfnProject struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
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

func (j *jsiiProxy_CfnProject) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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


// Create a new `AWS::CodeBuild::Project`.
func NewCfnProject(scope awscdk.Construct, id *string, props *CfnProjectProps) CfnProject {
	_init_.Initialize()

	if err := validateNewCfnProjectParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProject{}

	_jsii_.Create(
		"monocdk.aws_codebuild.CfnProject",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodeBuild::Project`.
func NewCfnProject_Override(c CfnProject, scope awscdk.Construct, id *string, props *CfnProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codebuild.CfnProject",
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
// Experimental.
func CfnProject_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProject_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.CfnProject",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnProject_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnProject_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.CfnProject",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.CfnProject",
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
		"monocdk.aws_codebuild.CfnProject",
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

func (c *jsiiProxy_CfnProject) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
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

func (c *jsiiProxy_CfnProject) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnProject) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnProject) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
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

func (c *jsiiProxy_CfnProject) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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

func (c *jsiiProxy_CfnProject) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnProject) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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

