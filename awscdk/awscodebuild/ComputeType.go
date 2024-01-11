package awscodebuild


// Build machine compute type.
//
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
type ComputeType string

const (
	ComputeType_SMALL ComputeType = "SMALL"
	ComputeType_MEDIUM ComputeType = "MEDIUM"
	ComputeType_LARGE ComputeType = "LARGE"
	ComputeType_X_LARGE ComputeType = "X_LARGE"
	ComputeType_X2_LARGE ComputeType = "X2_LARGE"
	ComputeType_LAMBDA_1GB ComputeType = "LAMBDA_1GB"
	ComputeType_LAMBDA_2GB ComputeType = "LAMBDA_2GB"
	ComputeType_LAMBDA_4GB ComputeType = "LAMBDA_4GB"
	ComputeType_LAMBDA_8GB ComputeType = "LAMBDA_8GB"
	ComputeType_LAMBDA_10GB ComputeType = "LAMBDA_10GB"
)

