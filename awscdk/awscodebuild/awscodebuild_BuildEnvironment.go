package awscodebuild


// Example:
//   var vpc vpc
//   var mySecurityGroup securityGroup
//
//   pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
//   	// Standard CodePipeline properties
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
//   	// Defaults for all CodeBuild projects
//   	CodeBuildDefaults: &CodeBuildOptions{
//   		// Prepend commands and configuration to all projects
//   		PartialBuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
//   			"version": jsii.String("0.2"),
//   		}),
//
//   		// Control the build environment
//   		BuildEnvironment: &BuildEnvironment{
//   			ComputeType: codebuild.ComputeType_LARGE,
//   		},
//
//   		// Control Elastic Network Interface creation
//   		Vpc: vpc,
//   		SubnetSelection: &SubnetSelection{
//   			SubnetType: ec2.SubnetType_PRIVATE_WITH_NAT,
//   		},
//   		SecurityGroups: []iSecurityGroup{
//   			mySecurityGroup,
//   		},
//
//   		// Additional policy statements for the execution role
//   		RolePolicy: []policyStatement{
//   			iam.NewPolicyStatement(&PolicyStatementProps{
//   			}),
//   		},
//   	},
//
//   	SynthCodeBuildDefaults: &CodeBuildOptions{
//   	},
//   	AssetPublishingCodeBuildDefaults: &CodeBuildOptions{
//   	},
//   	SelfMutationCodeBuildDefaults: &CodeBuildOptions{
//   	},
//   })
//
// Experimental.
type BuildEnvironment struct {
	// The image used for the builds.
	// Experimental.
	BuildImage IBuildImage `field:"optional" json:"buildImage" yaml:"buildImage"`
	// The location of the PEM-encoded certificate for the build project.
	// Experimental.
	Certificate *BuildEnvironmentCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// The type of compute to use for this build.
	//
	// See the {@link ComputeType} enum for the possible values.
	// Experimental.
	ComputeType ComputeType `field:"optional" json:"computeType" yaml:"computeType"`
	// The environment variables that your builds can use.
	// Experimental.
	EnvironmentVariables *map[string]*BuildEnvironmentVariable `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Indicates how the project builds Docker images.
	//
	// Specify true to enable
	// running the Docker daemon inside a Docker container. This value must be
	// set to true only if this build project will be used to build Docker
	// images, and the specified build environment image is not one provided by
	// AWS CodeBuild with Docker support. Otherwise, all associated builds that
	// attempt to interact with the Docker daemon will fail.
	// Experimental.
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
}

