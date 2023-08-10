package awscodebuild


// Example:
//   var vpc vpc
//   var mySecurityGroup securityGroup
//
//   pipelines.NewCodeBuildStep(jsii.String("Synth"), &CodeBuildStepProps{
//   	// ...standard ShellStep props...
//   	Commands: []*string{
//   	},
//   	Env: map[string]interface{}{
//   	},
//
//   	// If you are using a CodeBuildStep explicitly, set the 'cdk.out' directory
//   	// to be the synth step's output.
//   	PrimaryOutputDirectory: jsii.String("cdk.out"),
//
//   	// Control the name of the project
//   	ProjectName: jsii.String("MyProject"),
//
//   	// Control parts of the BuildSpec other than the regular 'build' and 'install' commands
//   	PartialBuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   	}),
//
//   	// Control the build environment
//   	BuildEnvironment: &BuildEnvironment{
//   		ComputeType: codebuild.ComputeType_LARGE,
//   		Privileged: jsii.Boolean(true),
//   	},
//   	Timeout: awscdk.Duration_Minutes(jsii.Number(90)),
//   	FileSystemLocations: []iFileSystemLocation{
//   		codebuild.FileSystemLocation_Efs(&EfsFileSystemLocationProps{
//   			Identifier: jsii.String("myidentifier2"),
//   			Location: jsii.String("myclodation.mydnsroot.com:/loc"),
//   			MountPoint: jsii.String("/media"),
//   			MountOptions: jsii.String("opts"),
//   		}),
//   	},
//
//   	// Control Elastic Network Interface creation
//   	Vpc: vpc,
//   	SubnetSelection: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	},
//   	SecurityGroups: []iSecurityGroup{
//   		mySecurityGroup,
//   	},
//
//   	// Control caching
//   	Cache: codebuild.Cache_Bucket(s3.NewBucket(this, jsii.String("Cache"))),
//
//   	// Additional policy statements for the execution role
//   	RolePolicyStatements: []policyStatement{
//   		iam.NewPolicyStatement(&PolicyStatementProps{
//   		}),
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
	// See the `ComputeType` enum for the possible values.
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

