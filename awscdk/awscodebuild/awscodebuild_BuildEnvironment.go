package awscodebuild


// Example:
//   var vpc vpc
//   var mySecurityGroup securityGroup
//
//   pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &codePipelineProps{
//   	// Standard CodePipeline properties
//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
//   		input: pipelines.codePipelineSource.connection(jsii.String("my-org/my-app"), jsii.String("main"), &connectionSourceOptions{
//   			connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
//
//   	// Defaults for all CodeBuild projects
//   	codeBuildDefaults: &codeBuildOptions{
//   		// Prepend commands and configuration to all projects
//   		partialBuildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   			"version": jsii.String("0.2"),
//   		}),
//
//   		// Control the build environment
//   		buildEnvironment: &buildEnvironment{
//   			computeType: codebuild.computeType_LARGE,
//   		},
//
//   		// Control Elastic Network Interface creation
//   		vpc: vpc,
//   		subnetSelection: &subnetSelection{
//   			subnetType: ec2.subnetType_PRIVATE_WITH_EGRESS,
//   		},
//   		securityGroups: []iSecurityGroup{
//   			mySecurityGroup,
//   		},
//
//   		// Additional policy statements for the execution role
//   		rolePolicy: []policyStatement{
//   			iam.NewPolicyStatement(&policyStatementProps{
//   			}),
//   		},
//   	},
//
//   	synthCodeBuildDefaults: &codeBuildOptions{
//   	},
//   	assetPublishingCodeBuildDefaults: &codeBuildOptions{
//   	},
//   	selfMutationCodeBuildDefaults: &codeBuildOptions{
//   	},
//   })
//
type BuildEnvironment struct {
	// The image used for the builds.
	BuildImage IBuildImage `field:"optional" json:"buildImage" yaml:"buildImage"`
	// The location of the PEM-encoded certificate for the build project.
	Certificate *BuildEnvironmentCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// The type of compute to use for this build.
	//
	// See the {@link ComputeType} enum for the possible values.
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
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
}

