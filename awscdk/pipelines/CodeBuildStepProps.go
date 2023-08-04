package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction props for a CodeBuildStep.
//
// Example:
//   pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
//   	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
//   		Input: pipelines.CodePipelineSource_Connection(jsii.String("my-org/my-app"), jsii.String("main"), &ConnectionSourceOptions{
//   			ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		Commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
//
//   	// Turn this on because the pipeline uses Docker image assets
//   	DockerEnabledForSelfMutation: jsii.Boolean(true),
//   })
//
//   pipeline.AddWave(jsii.String("MyWave"), &WaveOptions{
//   	Post: []step{
//   		pipelines.NewCodeBuildStep(jsii.String("RunApproval"), &CodeBuildStepProps{
//   			Commands: []*string{
//   				jsii.String("command-from-image"),
//   			},
//   			BuildEnvironment: &BuildEnvironment{
//   				// The user of a Docker image asset in the pipeline requires turning on
//   				// 'dockerEnabledForSelfMutation'.
//   				BuildImage: codebuild.LinuxBuildImage_FromAsset(this, jsii.String("Image"), &DockerImageAssetProps{
//   					Directory: jsii.String("./docker-image"),
//   				}),
//   			},
//   		}),
//   	},
//   })
//
type CodeBuildStepProps struct {
	// Commands to run.
	Commands *[]*string `field:"required" json:"commands" yaml:"commands"`
	// Additional FileSets to put in other directories.
	//
	// Specifies a mapping from directory name to FileSets. During the
	// script execution, the FileSets will be available in the directories
	// indicated.
	//
	// The directory names may be relative. For example, you can put
	// the main input and an additional input side-by-side with the
	// following configuration:
	//
	// ```ts
	// const script = new pipelines.ShellStep('MainScript', {
	//   commands: ['npm ci','npm run build','npx cdk synth'],
	//   input: pipelines.CodePipelineSource.gitHub('org/source1', 'main'),
	//   additionalInputs: {
	//     '../siblingdir': pipelines.CodePipelineSource.gitHub('org/source2', 'main'),
	//   }
	// });
	// ```.
	// Default: - No additional inputs.
	//
	AdditionalInputs *map[string]IFileSetProducer `field:"optional" json:"additionalInputs" yaml:"additionalInputs"`
	// Environment variables to set.
	// Default: - No environment variables.
	//
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// Set environment variables based on Stack Outputs.
	//
	// `ShellStep`s following stack or stage deployments may
	// access the `CfnOutput`s of those stacks to get access to
	// --for example--automatically generated resource names or
	// endpoint URLs.
	// Default: - No environment variables created from stack outputs.
	//
	EnvFromCfnOutputs *map[string]awscdk.CfnOutput `field:"optional" json:"envFromCfnOutputs" yaml:"envFromCfnOutputs"`
	// FileSet to run these scripts on.
	//
	// The files in the FileSet will be placed in the working directory when
	// the script is executed. Use `additionalInputs` to download file sets
	// to other directories as well.
	// Default: - No input specified.
	//
	Input IFileSetProducer `field:"optional" json:"input" yaml:"input"`
	// Installation commands to run before the regular commands.
	//
	// For deployment engines that support it, install commands will be classified
	// differently in the job history from the regular `commands`.
	// Default: - No installation commands.
	//
	InstallCommands *[]*string `field:"optional" json:"installCommands" yaml:"installCommands"`
	// The directory that will contain the primary output fileset.
	//
	// After running the script, the contents of the given directory
	// will be treated as the primary output of this Step.
	// Default: - No primary output.
	//
	PrimaryOutputDirectory *string `field:"optional" json:"primaryOutputDirectory" yaml:"primaryOutputDirectory"`
	// Custom execution role to be used for the Code Build Action.
	// Default: - A role is automatically created.
	//
	ActionRole awsiam.IRole `field:"optional" json:"actionRole" yaml:"actionRole"`
	// Changes to environment.
	//
	// This environment will be combined with the pipeline's default
	// environment.
	// Default: - Use the pipeline's default build environment.
	//
	BuildEnvironment *awscodebuild.BuildEnvironment `field:"optional" json:"buildEnvironment" yaml:"buildEnvironment"`
	// Caching strategy to use.
	// Default: - No cache.
	//
	Cache awscodebuild.Cache `field:"optional" json:"cache" yaml:"cache"`
	// ProjectFileSystemLocation objects for CodeBuild build projects.
	//
	// A ProjectFileSystemLocation object specifies the identifier, location, mountOptions, mountPoint,
	// and type of a file system created using Amazon Elastic File System.
	// Default: - no file system locations.
	//
	FileSystemLocations *[]awscodebuild.IFileSystemLocation `field:"optional" json:"fileSystemLocations" yaml:"fileSystemLocations"`
	// Information about logs for CodeBuild projects.
	//
	// A CodeBuild project can create logs in Amazon CloudWatch Logs, an S3 bucket, or both.
	// Default: - no log configuration is set.
	//
	Logging *awscodebuild.LoggingOptions `field:"optional" json:"logging" yaml:"logging"`
	// Additional configuration that can only be configured via BuildSpec.
	//
	// You should not use this to specify output artifacts; those
	// should be supplied via the other properties of this class, otherwise
	// CDK Pipelines won't be able to inspect the artifacts.
	//
	// Set the `commands` to an empty array if you want to fully specify
	// the BuildSpec using this field.
	//
	// The BuildSpec must be available inline--it cannot reference a file
	// on disk.
	// Default: - BuildSpec completely derived from other properties.
	//
	PartialBuildSpec awscodebuild.BuildSpec `field:"optional" json:"partialBuildSpec" yaml:"partialBuildSpec"`
	// Name for the generated CodeBuild project.
	// Default: - Automatically generated.
	//
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// Custom execution role to be used for the CodeBuild project.
	// Default: - A role is automatically created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	// Default: - No policy statements added to CodeBuild Project Role.
	//
	RolePolicyStatements *[]awsiam.PolicyStatement `field:"optional" json:"rolePolicyStatements" yaml:"rolePolicyStatements"`
	// Which security group to associate with the script's project network interfaces.
	//
	// If no security group is identified, one will be created automatically.
	//
	// Only used if 'vpc' is supplied.
	// Default: - Security group will be automatically created.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Default: - All private subnets.
	//
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	// Default: Duration.hours(1)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The VPC where to execute the SimpleSynth.
	// Default: - No VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

