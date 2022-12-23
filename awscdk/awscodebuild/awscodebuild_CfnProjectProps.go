package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectProps := &cfnProjectProps{
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
//   }
//
type CfnProjectProps struct {
	// `Artifacts` is a property of the [AWS::CodeBuild::Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies output settings for artifacts generated by an AWS CodeBuild build.
	Artifacts interface{} `field:"required" json:"artifacts" yaml:"artifacts"`
	// The build environment settings for the project, such as the environment type or the environment variables to use for the build environment.
	Environment interface{} `field:"required" json:"environment" yaml:"environment"`
	// The ARN of the IAM role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// The source code settings for the project, such as the source code's repository type and location.
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// Indicates whether AWS CodeBuild generates a publicly accessible URL for your project's build badge.
	//
	// For more information, see [Build Badges Sample](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-badges.html) in the *AWS CodeBuild User Guide* .
	//
	// > Including build badges with your project is currently not supported if the source type is CodePipeline. If you specify `CODEPIPELINE` for the `Source` property, do not specify the `BadgeEnabled` property.
	BadgeEnabled interface{} `field:"optional" json:"badgeEnabled" yaml:"badgeEnabled"`
	// A `ProjectBuildBatchConfig` object that defines the batch build options for the project.
	BuildBatchConfig interface{} `field:"optional" json:"buildBatchConfig" yaml:"buildBatchConfig"`
	// Settings that AWS CodeBuild uses to store and reuse build dependencies.
	Cache interface{} `field:"optional" json:"cache" yaml:"cache"`
	// The maximum number of concurrent builds that are allowed for this project.
	//
	// New builds are only started if the current number of builds is less than or equal to this limit. If the current build count meets this limit, new builds are throttled and are not run.
	ConcurrentBuildLimit *float64 `field:"optional" json:"concurrentBuildLimit" yaml:"concurrentBuildLimit"`
	// A description that makes the build project easy to identify.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS Key Management Service customer master key (CMK) to be used for encrypting the build output artifacts.
	//
	// > You can use a cross-account KMS key to encrypt the build output artifacts if your service role has permission to that key.
	//
	// You can specify either the Amazon Resource Name (ARN) of the CMK or, if available, the CMK's alias (using the format `alias/<alias-name>` ). If you don't specify a value, CodeBuild uses the managed CMK for Amazon Simple Storage Service (Amazon S3).
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// An array of `ProjectFileSystemLocation` objects for a CodeBuild build project.
	//
	// A `ProjectFileSystemLocation` object specifies the `identifier` , `location` , `mountOptions` , `mountPoint` , and `type` of a file system created using Amazon Elastic File System.
	FileSystemLocations interface{} `field:"optional" json:"fileSystemLocations" yaml:"fileSystemLocations"`
	// Information about logs for the build project.
	//
	// A project can create logs in CloudWatch Logs, an S3 bucket, or both.
	LogsConfig interface{} `field:"optional" json:"logsConfig" yaml:"logsConfig"`
	// The name of the build project.
	//
	// The name must be unique across all of the projects in your AWS account .
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The number of minutes a build is allowed to be queued before it times out.
	QueuedTimeoutInMinutes *float64 `field:"optional" json:"queuedTimeoutInMinutes" yaml:"queuedTimeoutInMinutes"`
	// The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds.
	ResourceAccessRole *string `field:"optional" json:"resourceAccessRole" yaml:"resourceAccessRole"`
	// A list of `Artifacts` objects.
	//
	// Each artifacts object specifies output settings that the project generates during a build.
	SecondaryArtifacts interface{} `field:"optional" json:"secondaryArtifacts" yaml:"secondaryArtifacts"`
	// An array of `ProjectSource` objects.
	SecondarySources interface{} `field:"optional" json:"secondarySources" yaml:"secondarySources"`
	// An array of `ProjectSourceVersion` objects.
	//
	// If `secondarySourceVersions` is specified at the build level, then they take over these `secondarySourceVersions` (at the project level).
	SecondarySourceVersions interface{} `field:"optional" json:"secondarySourceVersions" yaml:"secondarySourceVersions"`
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
	SourceVersion *string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
	// An arbitrary set of tags (key-value pairs) for the AWS CodeBuild project.
	//
	// These tags are available for use by AWS services that support AWS CodeBuild build project tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// How long, in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait before timing out any related build that did not get marked as completed.
	//
	// The default is 60 minutes.
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
	// For an existing AWS CodeBuild build project that has its source code stored in a GitHub repository, enables AWS CodeBuild to begin automatically rebuilding the source code every time a code change is pushed to the repository.
	Triggers interface{} `field:"optional" json:"triggers" yaml:"triggers"`
	// Specifies the visibility of the project's builds. Possible values are:.
	//
	// - **PUBLIC_READ** - The project builds are visible to the public.
	// - **PRIVATE** - The project builds are not visible to the public.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// `VpcConfig` specifies settings that enable AWS CodeBuild to access resources in an Amazon VPC.
	//
	// For more information, see [Use AWS CodeBuild with Amazon Virtual Private Cloud](https://docs.aws.amazon.com/codebuild/latest/userguide/vpc-support.html) in the *AWS CodeBuild User Guide* .
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

