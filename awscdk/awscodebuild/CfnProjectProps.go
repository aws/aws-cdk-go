package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectProps := &CfnProjectProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html
//
type CfnProjectProps struct {
	// `Artifacts` is a property of the [AWS::CodeBuild::Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies output settings for artifacts generated by an AWS CodeBuild build.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-artifacts
	//
	Artifacts interface{} `field:"required" json:"artifacts" yaml:"artifacts"`
	// The build environment settings for the project, such as the environment type or the environment variables to use for the build environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-environment
	//
	Environment interface{} `field:"required" json:"environment" yaml:"environment"`
	// The ARN of the IAM role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-servicerole
	//
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// The source code settings for the project, such as the source code's repository type and location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// Indicates whether AWS CodeBuild generates a publicly accessible URL for your project's build badge.
	//
	// For more information, see [Build Badges Sample](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-badges.html) in the *AWS CodeBuild User Guide* .
	//
	// > Including build badges with your project is currently not supported if the source type is CodePipeline. If you specify `CODEPIPELINE` for the `Source` property, do not specify the `BadgeEnabled` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-badgeenabled
	//
	BadgeEnabled interface{} `field:"optional" json:"badgeEnabled" yaml:"badgeEnabled"`
	// A `ProjectBuildBatchConfig` object that defines the batch build options for the project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-buildbatchconfig
	//
	BuildBatchConfig interface{} `field:"optional" json:"buildBatchConfig" yaml:"buildBatchConfig"`
	// Settings that AWS CodeBuild uses to store and reuse build dependencies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-cache
	//
	Cache interface{} `field:"optional" json:"cache" yaml:"cache"`
	// The maximum number of concurrent builds that are allowed for this project.
	//
	// New builds are only started if the current number of builds is less than or equal to this limit. If the current build count meets this limit, new builds are throttled and are not run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-concurrentbuildlimit
	//
	ConcurrentBuildLimit *float64 `field:"optional" json:"concurrentBuildLimit" yaml:"concurrentBuildLimit"`
	// A description that makes the build project easy to identify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS Key Management Service customer master key (CMK) to be used for encrypting the build output artifacts.
	//
	// > You can use a cross-account KMS key to encrypt the build output artifacts if your service role has permission to that key.
	//
	// You can specify either the Amazon Resource Name (ARN) of the CMK or, if available, the CMK's alias (using the format `alias/<alias-name>` ). If you don't specify a value, CodeBuild uses the managed CMK for Amazon Simple Storage Service (Amazon S3).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-encryptionkey
	//
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// An array of `ProjectFileSystemLocation` objects for a CodeBuild build project.
	//
	// A `ProjectFileSystemLocation` object specifies the `identifier` , `location` , `mountOptions` , `mountPoint` , and `type` of a file system created using Amazon Elastic File System.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-filesystemlocations
	//
	FileSystemLocations interface{} `field:"optional" json:"fileSystemLocations" yaml:"fileSystemLocations"`
	// Information about logs for the build project.
	//
	// A project can create logs in CloudWatch Logs, an S3 bucket, or both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-logsconfig
	//
	LogsConfig interface{} `field:"optional" json:"logsConfig" yaml:"logsConfig"`
	// The name of the build project.
	//
	// The name must be unique across all of the projects in your AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The number of minutes a build is allowed to be queued before it times out.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-queuedtimeoutinminutes
	//
	QueuedTimeoutInMinutes *float64 `field:"optional" json:"queuedTimeoutInMinutes" yaml:"queuedTimeoutInMinutes"`
	// The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-resourceaccessrole
	//
	ResourceAccessRole *string `field:"optional" json:"resourceAccessRole" yaml:"resourceAccessRole"`
	// A list of `Artifacts` objects.
	//
	// Each artifacts object specifies output settings that the project generates during a build.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-secondaryartifacts
	//
	SecondaryArtifacts interface{} `field:"optional" json:"secondaryArtifacts" yaml:"secondaryArtifacts"`
	// An array of `ProjectSource` objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-secondarysources
	//
	SecondarySources interface{} `field:"optional" json:"secondarySources" yaml:"secondarySources"`
	// An array of `ProjectSourceVersion` objects.
	//
	// If `secondarySourceVersions` is specified at the build level, then they take over these `secondarySourceVersions` (at the project level).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-secondarysourceversions
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-sourceversion
	//
	SourceVersion *string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
	// An arbitrary set of tags (key-value pairs) for the AWS CodeBuild project.
	//
	// These tags are available for use by AWS services that support AWS CodeBuild build project tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// How long, in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait before timing out any related build that did not get marked as completed.
	//
	// The default is 60 minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-timeoutinminutes
	//
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
	// For an existing AWS CodeBuild build project that has its source code stored in a GitHub repository, enables AWS CodeBuild to begin automatically rebuilding the source code every time a code change is pushed to the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-triggers
	//
	Triggers interface{} `field:"optional" json:"triggers" yaml:"triggers"`
	// Specifies the visibility of the project's builds. Possible values are:.
	//
	// - **PUBLIC_READ** - The project builds are visible to the public.
	// - **PRIVATE** - The project builds are not visible to the public.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// `VpcConfig` specifies settings that enable AWS CodeBuild to access resources in an Amazon VPC.
	//
	// For more information, see [Use AWS CodeBuild with Amazon Virtual Private Cloud](https://docs.aws.amazon.com/codebuild/latest/userguide/vpc-support.html) in the *AWS CodeBuild User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

