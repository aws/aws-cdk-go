package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Example:
//   var ecrRepository repository
//
//
//   codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
//   	Environment: &BuildEnvironment{
//   		BuildImage: codebuild.WindowsBuildImage_FromEcrRepository(ecrRepository, jsii.String("v1.0"), codebuild.WindowsImageType_SERVER_2019),
//   		// optional certificate to include in the build image
//   		Certificate: &BuildEnvironmentCertificate{
//   			Bucket: s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("amzn-s3-demo-bucket")),
//   			ObjectKey: jsii.String("path/to/cert.pem"),
//   		},
//   	},
//   })
//
type ProjectProps struct {
	// Whether to allow the CodeBuild to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// CodeBuild project to connect to network targets.
	//
	// Only used if 'vpc' is supplied.
	// Default: true.
	//
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Indicates whether AWS CodeBuild generates a publicly accessible URL for your project's build badge.
	//
	// For more information, see Build Badges Sample
	// in the AWS CodeBuild User Guide.
	// Default: false.
	//
	Badge *bool `field:"optional" json:"badge" yaml:"badge"`
	// Filename or contents of buildspec in JSON format.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec-ref-example
	//
	// Default: - Empty buildspec.
	//
	BuildSpec BuildSpec `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// Caching strategy to use.
	// Default: Cache.none
	//
	Cache Cache `field:"optional" json:"cache" yaml:"cache"`
	// Whether to check for the presence of any secrets in the environment variables of the default type, BuildEnvironmentVariableType.PLAINTEXT. Since using a secret for the value of that kind of variable would result in it being displayed in plain text in the AWS Console, the construct will throw an exception if it detects a secret was passed there. Pass this property as false if you want to skip this validation, and keep using a secret in a plain text environment variable.
	// Default: true.
	//
	CheckSecretsInPlainTextEnvVariables *bool `field:"optional" json:"checkSecretsInPlainTextEnvVariables" yaml:"checkSecretsInPlainTextEnvVariables"`
	// Maximum number of concurrent builds.
	//
	// Minimum value is 1 and maximum is account build limit.
	// Default: - no explicit limit is set.
	//
	ConcurrentBuildLimit *float64 `field:"optional" json:"concurrentBuildLimit" yaml:"concurrentBuildLimit"`
	// A description of the project.
	//
	// Use the description to identify the purpose
	// of the project.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Encryption key to use to read and write artifacts.
	// Default: - The AWS-managed CMK for Amazon Simple Storage Service (Amazon S3) is used.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Build environment to use for the build.
	// Default: BuildEnvironment.LinuxBuildImage.STANDARD_7_0
	//
	Environment *BuildEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// Additional environment variables to add to the build environment.
	// Default: - No additional environment variables are specified.
	//
	EnvironmentVariables *map[string]*BuildEnvironmentVariable `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// An  ProjectFileSystemLocation objects for a CodeBuild build project.
	//
	// A ProjectFileSystemLocation object specifies the identifier, location, mountOptions, mountPoint,
	// and type of a file system created using Amazon Elastic File System.
	// Default: - no file system locations.
	//
	FileSystemLocations *[]IFileSystemLocation `field:"optional" json:"fileSystemLocations" yaml:"fileSystemLocations"`
	// Add permissions to this project's role to create and use test report groups with name starting with the name of this project.
	//
	// That is the standard report group that gets created when a simple name
	// (in contrast to an ARN)
	// is used in the 'reports' section of the buildspec of this project.
	// This is usually harmless, but you can turn these off if you don't plan on using test
	// reports in this project.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/test-report-group-naming.html
	//
	// Default: true.
	//
	GrantReportGroupPermissions *bool `field:"optional" json:"grantReportGroupPermissions" yaml:"grantReportGroupPermissions"`
	// Information about logs for the build project.
	//
	// A project can create logs in Amazon CloudWatch Logs, an S3 bucket, or both.
	// Default: - no log configuration is set.
	//
	Logging *LoggingOptions `field:"optional" json:"logging" yaml:"logging"`
	// The physical, human-readable name of the CodeBuild Project.
	// Default: - Name is automatically generated.
	//
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// The number of minutes after which AWS CodeBuild stops the build if it's still in queue.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	// Default: - no queue timeout is set.
	//
	QueuedTimeout awscdk.Duration `field:"optional" json:"queuedTimeout" yaml:"queuedTimeout"`
	// Service Role to assume while running the build.
	// Default: - A role will be created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// What security group to associate with the codebuild project's network interfaces.
	//
	// If no security group is identified, one will be created automatically.
	//
	// Only used if 'vpc' is supplied.
	// Default: - Security group will be automatically created.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Add the permissions necessary for debugging builds with SSM Session Manager.
	//
	// If the following prerequisites have been met:
	//
	// - The necessary permissions have been added by setting this flag to true.
	// - The build image has the SSM agent installed (true for default CodeBuild images).
	// - The build is started with [debugSessionEnabled](https://docs.aws.amazon.com/codebuild/latest/APIReference/API_StartBuild.html#CodeBuild-StartBuild-request-debugSessionEnabled) set to true.
	//
	// Then the build container can be paused and inspected using Session Manager
	// by invoking the `codebuild-breakpoint` command somewhere during the build.
	//
	// `codebuild-breakpoint` commands will be ignored if the build is not started
	// with `debugSessionEnabled=true`.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/session-manager.html
	//
	// Default: false.
	//
	SsmSessionPermissions *bool `field:"optional" json:"ssmSessionPermissions" yaml:"ssmSessionPermissions"`
	// Where to place the network interfaces within the VPC.
	//
	// To access AWS services, your CodeBuild project needs to be in one of the following types of subnets:
	//
	// 1. Subnets with access to the internet (of type PRIVATE_WITH_EGRESS).
	// 2. Private subnets unconnected to the internet, but with [VPC endpoints](https://docs.aws.amazon.com/codebuild/latest/userguide/use-vpc-endpoints-with-codebuild.html) for the necessary services.
	//
	// If you don't specify a subnet selection, the default behavior is to use PRIVATE_WITH_EGRESS subnets first if they exist,
	// then PRIVATE_WITHOUT_EGRESS, and finally PUBLIC subnets. If your VPC doesn't have PRIVATE_WITH_EGRESS subnets but you need
	// AWS service access, add VPC Endpoints to your private subnets.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/vpc-support.html
	//
	// Default: - private subnets if available else public subnets.
	//
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	// Default: Duration.hours(1)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Specifies the visibility of the project's builds.
	// Default: - no visibility is set.
	//
	Visibility ProjectVisibility `field:"optional" json:"visibility" yaml:"visibility"`
	// VPC network to place codebuild network interfaces.
	//
	// Specify this if the codebuild project needs to access resources in a VPC.
	// Default: - No VPC is specified.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Defines where build artifacts will be stored.
	//
	// Could be: PipelineBuildArtifacts, NoArtifacts and S3Artifacts.
	// Default: NoArtifacts.
	//
	Artifacts IArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	// The secondary artifacts for the Project.
	//
	// Can also be added after the Project has been created by using the `Project#addSecondaryArtifact` method.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
	//
	// Default: - No secondary artifacts.
	//
	SecondaryArtifacts *[]IArtifacts `field:"optional" json:"secondaryArtifacts" yaml:"secondaryArtifacts"`
	// The secondary sources for the Project.
	//
	// Can be also added after the Project has been created by using the `Project#addSecondarySource` method.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
	//
	// Default: - No secondary sources.
	//
	SecondarySources *[]ISource `field:"optional" json:"secondarySources" yaml:"secondarySources"`
	// The source of the build.
	//
	// *Note*: if `NoSource` is given as the source,
	// then you need to provide an explicit `buildSpec`.
	// Default: - NoSource.
	//
	Source ISource `field:"optional" json:"source" yaml:"source"`
}

