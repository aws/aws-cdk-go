package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Properties for a CdkPipeline.
//
// Example:
//   sourceArtifact := codepipeline.NewArtifact()
//   cloudAssemblyArtifact := codepipeline.NewArtifact()
//   pipeline := pipelines.NewCdkPipeline(this, jsii.String("MyPipeline"), &cdkPipelineProps{
//   	cloudAssemblyArtifact: cloudAssemblyArtifact,
//   	synthAction: pipelines.simpleSynthAction.standardNpmSynth(&standardNpmSynthOptions{
//   		sourceArtifact: sourceArtifact,
//   		cloudAssemblyArtifact: cloudAssemblyArtifact,
//   		environment: &buildEnvironment{
//   			privileged: jsii.Boolean(true),
//   		},
//   	}),
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type CdkPipelineProps struct {
	// The artifact you have defined to be the artifact to hold the cloudAssemblyArtifact for the synth action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyArtifact awscodepipeline.Artifact `field:"required" json:"cloudAssemblyArtifact" yaml:"cloudAssemblyArtifact"`
	// Custom BuildSpec that is merged with generated one (for asset publishing actions).
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetBuildSpec awscodebuild.BuildSpec `field:"optional" json:"assetBuildSpec" yaml:"assetBuildSpec"`
	// Additional commands to run before installing cdk-assets during the asset publishing step Use this to setup proxies or npm mirrors.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetPreInstallCommands *[]*string `field:"optional" json:"assetPreInstallCommands" yaml:"assetPreInstallCommands"`
	// CDK CLI version to use in pipeline.
	//
	// Some Actions in the pipeline will download and run a version of the CDK
	// CLI. Specify the version here.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CdkCliVersion *string `field:"optional" json:"cdkCliVersion" yaml:"cdkCliVersion"`
	// Existing CodePipeline to add deployment stages to.
	//
	// Use this if you want more control over the CodePipeline that gets created.
	// You can choose to not pass this value, in which case a new CodePipeline is
	// created with default settings.
	//
	// If you pass an existing CodePipeline, it should have been created
	// with `restartExecutionOnUpdate: true`.
	//
	// [disable-awslint:ref-via-interface].
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CodePipeline awscodepipeline.Pipeline `field:"optional" json:"codePipeline" yaml:"codePipeline"`
	// Create KMS keys for cross-account deployments.
	//
	// This controls whether the pipeline is enabled for cross-account deployments.
	//
	// Can only be set if `codePipeline` is not set.
	//
	// By default cross-account deployments are enabled, but this feature requires
	// that KMS Customer Master Keys are created which have a cost of $1/month.
	//
	// If you do not need cross-account deployments, you can set this to `false` to
	// not create those keys and save on that cost (the artifact bucket will be
	// encrypted with an AWS-managed key). However, cross-account deployments will
	// no longer be possible.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CrossAccountKeys *bool `field:"optional" json:"crossAccountKeys" yaml:"crossAccountKeys"`
	// A list of credentials used to authenticate to Docker registries.
	//
	// Specify any credentials necessary within the pipeline to build, synth, update, or publish assets.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	DockerCredentials *[]DockerCredential `field:"optional" json:"dockerCredentials" yaml:"dockerCredentials"`
	// Enables KMS key rotation for cross-account keys.
	//
	// Cannot be set if `crossAccountKeys` was set to `false`.
	//
	// Key rotation costs $1/month when enabled.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	EnableKeyRotation *bool `field:"optional" json:"enableKeyRotation" yaml:"enableKeyRotation"`
	// Name of the pipeline.
	//
	// Can only be set if `codePipeline` is not set.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PipelineName *string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
	// Whether the pipeline will update itself.
	//
	// This needs to be set to `true` to allow the pipeline to reconfigure
	// itself when assets or stages are being added to it, and `true` is the
	// recommended setting.
	//
	// You can temporarily set this to `false` while you are iterating
	// on the pipeline itself and prefer to deploy changes using `cdk deploy`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SelfMutating *bool `field:"optional" json:"selfMutating" yaml:"selfMutating"`
	// Custom BuildSpec that is merged with generated one (for self-mutation stage).
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SelfMutationBuildSpec awscodebuild.BuildSpec `field:"optional" json:"selfMutationBuildSpec" yaml:"selfMutationBuildSpec"`
	// Whether this pipeline creates one asset upload action per asset type or one asset upload per asset.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SinglePublisherPerType *bool `field:"optional" json:"singlePublisherPerType" yaml:"singlePublisherPerType"`
	// The CodePipeline action used to retrieve the CDK app's source.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SourceAction awscodepipeline.IAction `field:"optional" json:"sourceAction" yaml:"sourceAction"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// Whether the pipeline needs to build Docker images in the UpdatePipeline stage.
	//
	// If the UpdatePipeline stage tries to build a Docker image and this flag is not
	// set to `true`, the build step will run in non-privileged mode and consequently
	// will fail with a message like:
	//
	// > Cannot connect to the Docker daemon at unix:///var/run/docker.sock.
	// > Is the docker daemon running?
	//
	// This flag has an effect only if `selfMutating` is also `true`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SupportDockerAssets *bool `field:"optional" json:"supportDockerAssets" yaml:"supportDockerAssets"`
	// The CodePipeline action build and synthesis step of the CDK app.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SynthAction awscodepipeline.IAction `field:"optional" json:"synthAction" yaml:"synthAction"`
	// The VPC where to execute the CdkPipeline actions.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

