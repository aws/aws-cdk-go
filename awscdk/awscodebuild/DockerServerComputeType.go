package awscodebuild


// Docker server compute type.
//
// Example:
//   var vpc Vpc
//   var mySecurityGroup SecurityGroup
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
//   		DockerServer: &DockerServerOptions{
//   			ComputeType: codebuild.DockerServerComputeType_SMALL,
//   			SecurityGroups: []ISecurityGroup{
//   				mySecurityGroup,
//   			},
//   		},
//   	},
//   	Timeout: awscdk.Duration_Minutes(jsii.Number(90)),
//   	FileSystemLocations: []IFileSystemLocation{
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
//   	SecurityGroups: []ISecurityGroup{
//   		mySecurityGroup,
//   	},
//
//   	// Control caching
//   	Cache: codebuild.Cache_Bucket(s3.NewBucket(this, jsii.String("Cache"))),
//
//   	// Additional policy statements for the execution role
//   	RolePolicyStatements: []PolicyStatement{
//   		iam.NewPolicyStatement(&PolicyStatementProps{
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/codebuild/latest/APIReference/API_DockerServer.html
//
type DockerServerComputeType string

const (
	// BUILD_GENERAL1_SMALL.
	DockerServerComputeType_SMALL DockerServerComputeType = "SMALL"
	// BUILD_GENERAL1_MEDIUM.
	DockerServerComputeType_MEDIUM DockerServerComputeType = "MEDIUM"
	// BUILD_GENERAL1_LARGE.
	DockerServerComputeType_LARGE DockerServerComputeType = "LARGE"
	// BUILD_GENERAL1_XLARGE.
	DockerServerComputeType_X_LARGE DockerServerComputeType = "X_LARGE"
	// BUILD_GENERAL1_2XLARGE.
	DockerServerComputeType_X2_LARGE DockerServerComputeType = "X2_LARGE"
	// ATTRIBUTE_BASED_COMPUTE.
	DockerServerComputeType_ATTRIBUTE_BASED_COMPUTE DockerServerComputeType = "ATTRIBUTE_BASED_COMPUTE"
	// BUILD_LAMBDA_10GB.
	DockerServerComputeType_BUILD_LAMBDA_10GB DockerServerComputeType = "BUILD_LAMBDA_10GB"
	// BUILD_LAMBDA_1GB.
	DockerServerComputeType_BUILD_LAMBDA_1GB DockerServerComputeType = "BUILD_LAMBDA_1GB"
	// BUILD_LAMBDA_2GB.
	DockerServerComputeType_BUILD_LAMBDA_2GB DockerServerComputeType = "BUILD_LAMBDA_2GB"
	// BUILD_LAMBDA_4GB.
	DockerServerComputeType_BUILD_LAMBDA_4GB DockerServerComputeType = "BUILD_LAMBDA_4GB"
	// BUILD_LAMBDA_8GB.
	DockerServerComputeType_BUILD_LAMBDA_8GB DockerServerComputeType = "BUILD_LAMBDA_8GB"
	// CUSTOM_INSTANCE_TYPE.
	DockerServerComputeType_CUSTOM_INSTANCE_TYPE DockerServerComputeType = "CUSTOM_INSTANCE_TYPE"
)

