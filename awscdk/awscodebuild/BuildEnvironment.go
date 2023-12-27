package awscodebuild


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
type BuildEnvironment struct {
	// The image used for the builds.
	// Default: LinuxBuildImage.STANDARD_1_0
	//
	BuildImage IBuildImage `field:"optional" json:"buildImage" yaml:"buildImage"`
	// The location of the PEM-encoded certificate for the build project.
	// Default: - No external certificate is added to the project.
	//
	Certificate *BuildEnvironmentCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// The type of compute to use for this build.
	//
	// See the `ComputeType` enum for the possible values.
	// Default: taken from `#buildImage#defaultComputeType`.
	//
	ComputeType ComputeType `field:"optional" json:"computeType" yaml:"computeType"`
	// The environment variables that your builds can use.
	EnvironmentVariables *map[string]*BuildEnvironmentVariable `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Indicates how the project builds Docker images.
	//
	// Specify true to enable
	// running the Docker daemon inside a Docker container. This value must be
	// set to true only if this build project will be used to build Docker
	// images, and the specified build environment image is not one provided by
	// AWS CodeBuild with Docker support. Otherwise, all associated builds that
	// attempt to interact with the Docker daemon will fail.
	// Default: false.
	//
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
}

