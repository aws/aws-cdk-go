package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Example:
//   // Create a Cloudfront Web Distribution
//   import cloudfront "github.com/aws/aws-cdk-go/awscdk"
//   var distribution distribution
//
//
//   // Create the build project that will invalidate the cache
//   invalidateBuildProject := codebuild.NewPipelineProject(this, jsii.String("InvalidateProject"), &PipelineProjectProps{
//   	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string][]*string{
//   			"build": map[string][]*string{
//   				"commands": []*string{
//   					jsii.String("aws cloudfront create-invalidation --distribution-id ${CLOUDFRONT_ID} --paths \"/*\""),
//   				},
//   			},
//   		},
//   	}),
//   	EnvironmentVariables: map[string]buildEnvironmentVariable{
//   		"CLOUDFRONT_ID": &buildEnvironmentVariable{
//   			"value": distribution.distributionId,
//   		},
//   	},
//   })
//
//   // Add Cloudfront invalidation permissions to the project
//   distributionArn := fmt.Sprintf("arn:aws:cloudfront::%v:distribution/%v", this.Account, distribution.DistributionId)
//   invalidateBuildProject.addToRolePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Resources: []*string{
//   		distributionArn,
//   	},
//   	Actions: []*string{
//   		jsii.String("cloudfront:CreateInvalidation"),
//   	},
//   }))
//
//   // Create the pipeline (here only the S3 deploy and Invalidate cache build)
//   deployBucket := s3.NewBucket(this, jsii.String("DeployBucket"))
//   deployInput := codepipeline.NewArtifact()
//   codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
//   	Stages: []stageProps{
//   		&stageProps{
//   			StageName: jsii.String("Deploy"),
//   			Actions: []iAction{
//   				codepipeline_actions.NewS3DeployAction(&S3DeployActionProps{
//   					ActionName: jsii.String("S3Deploy"),
//   					Bucket: deployBucket,
//   					Input: deployInput,
//   					RunOrder: jsii.Number(1),
//   				}),
//   				codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   					ActionName: jsii.String("InvalidateCache"),
//   					Project: invalidateBuildProject,
//   					Input: deployInput,
//   					RunOrder: jsii.Number(2),
//   				}),
//   			},
//   		},
//   	},
//   })
//
type PipelineProjectProps struct {
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
	// Default: BuildEnvironment.LinuxBuildImage.STANDARD_1_0
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
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/vpc-support.html for more details.
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
	// VPC network to place codebuild network interfaces.
	//
	// Specify this if the codebuild project needs to access resources in a VPC.
	// Default: - No VPC is specified.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

