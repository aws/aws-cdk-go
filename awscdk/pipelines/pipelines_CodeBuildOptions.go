package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options for customizing a single CodeBuild project.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var source iFileSetProducer // the repository source
//   var synthCommands []*string // Commands to synthesize your app
//   var installCommands []*string
//   // Commands to install your toolchain
//
//   pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &codePipelineProps{
//   	// Standard CodePipeline properties...
//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
//   		input: source,
//   		commands: synthCommands,
//   	}),
//
//   	// Configure CodeBuild to use a drop-in Docker replacement.
//   	codeBuildDefaults: &codeBuildOptions{
//   		buildEnvironment: &buildEnvironment{
//   			partialBuildSpec: codebuild_BuildSpec_FromObject(map[string]map[string]map[string][]*string{
//   				"phases": map[string]map[string][]*string{
//   					"install": map[string][]*string{
//   						// Add the shell commands to install your drop-in Docker
//   						// replacement to the CodeBuild enviromment.
//   						"commands": installCommands,
//   					},
//   				},
//   			}),
//   			environmentVariables: map[string]buildEnvironmentVariable{
//   				// Instruct the AWS CDK to use `drop-in-replacement` instead of
//   				// `docker` when building / publishing docker images.
//   				// e.g., `drop-in-replacement build . -f path/to/Dockerfile`
//   				"CDK_DOCKER": jsii.String("drop-in-replacement"),
//   			},
//   		},
//   	},
//   })
//
type CodeBuildOptions struct {
	// Partial build environment, will be combined with other build environments that apply.
	BuildEnvironment *awscodebuild.BuildEnvironment `field:"optional" json:"buildEnvironment" yaml:"buildEnvironment"`
	// Caching strategy to use.
	Cache awscodebuild.Cache `field:"optional" json:"cache" yaml:"cache"`
	// Partial buildspec, will be combined with other buildspecs that apply.
	//
	// The BuildSpec must be available inline--it cannot reference a file
	// on disk.
	PartialBuildSpec awscodebuild.BuildSpec `field:"optional" json:"partialBuildSpec" yaml:"partialBuildSpec"`
	// Policy statements to add to role.
	RolePolicy *[]awsiam.PolicyStatement `field:"optional" json:"rolePolicy" yaml:"rolePolicy"`
	// Which security group(s) to associate with the project network interfaces.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The VPC where to create the CodeBuild network interfaces in.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

