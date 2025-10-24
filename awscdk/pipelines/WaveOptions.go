package pipelines


// Options to pass to `addWave`.
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
//   	Post: []Step{
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
type WaveOptions struct {
	// Additional steps to run after all of the stages in the wave.
	// Default: - No additional steps.
	//
	Post *[]Step `field:"optional" json:"post" yaml:"post"`
	// Additional steps to run before any of the stages in the wave.
	// Default: - No additional steps.
	//
	Pre *[]Step `field:"optional" json:"pre" yaml:"pre"`
}

