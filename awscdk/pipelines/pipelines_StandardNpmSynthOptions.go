package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Options for a convention-based synth using NPM.
//
// Example:
//   sourceArtifact := codepipeline.NewArtifact()
//   cloudAssemblyArtifact := codepipeline.NewArtifact()
//   pipeline := pipelines.NewCdkPipeline(this, jsii.String("MyPipeline"), &CdkPipelineProps{
//   	CloudAssemblyArtifact: CloudAssemblyArtifact,
//   	SynthAction: pipelines.SimpleSynthAction_StandardNpmSynth(&StandardNpmSynthOptions{
//   		SourceArtifact: *SourceArtifact,
//   		CloudAssemblyArtifact: *CloudAssemblyArtifact,
//   		Environment: &BuildEnvironment{
//   			Privileged: jsii.Boolean(true),
//   		},
//   	}),
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type StandardNpmSynthOptions struct {
	// The artifact where the CloudAssembly should be emitted.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyArtifact awscodepipeline.Artifact `field:"required" json:"cloudAssemblyArtifact" yaml:"cloudAssemblyArtifact"`
	// The source artifact of the CodePipeline.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SourceArtifact awscodepipeline.Artifact `field:"required" json:"sourceArtifact" yaml:"sourceArtifact"`
	// Name of the build action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// Produce additional output artifacts after the build based on the given directories.
	//
	// Can be used to produce additional artifacts during the build step,
	// separate from the cloud assembly, which can be used further on in the
	// pipeline.
	//
	// Directories are evaluated with respect to `subdirectory`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AdditionalArtifacts *[]*AdditionalArtifact `field:"optional" json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// custom BuildSpec that is merged with the generated one.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildSpec awscodebuild.BuildSpec `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// Environment variables to copy over from parent env.
	//
	// These are environment variables that are being used by the build.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CopyEnvironmentVariables *[]*string `field:"optional" json:"copyEnvironmentVariables" yaml:"copyEnvironmentVariables"`
	// Build environment to use for CodeBuild job.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Environment *awscodebuild.BuildEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// Environment variables to send into build.
	//
	// NOTE: You may run into the 1000-character limit for the Action configuration if you have a large
	// number of variables or if their names or values are very long.
	// If you do, pass them to the underlying CodeBuild project directly in `environment` instead.
	// However, you will not be able to use CodePipeline Variables in this case.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Name of the CodeBuild project.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	RolePolicyStatements *[]awsiam.PolicyStatement `field:"optional" json:"rolePolicyStatements" yaml:"rolePolicyStatements"`
	// Directory inside the source where package.json and cdk.json are located.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The VPC where to execute the SimpleSynth.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The build command.
	//
	// By default, we assume NPM projects are either written in JavaScript or are
	// using `ts-node`, so don't need a build command.
	//
	// Otherwise, put the build command here, for example `npm run build`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// The install command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	InstallCommand *string `field:"optional" json:"installCommand" yaml:"installCommand"`
	// The synth command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SynthCommand *string `field:"optional" json:"synthCommand" yaml:"synthCommand"`
	// Test commands.
	//
	// These commands are run after the build commands but before the
	// synth command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	TestCommands *[]*string `field:"optional" json:"testCommands" yaml:"testCommands"`
}

