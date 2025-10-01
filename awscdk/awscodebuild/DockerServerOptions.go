package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The Docker server configuration CodeBuild use to build your Docker image.
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
//   		DockerServer: &DockerServerOptions{
//   			ComputeType: codebuild.DockerServerComputeType_SMALL,
//   			SecurityGroups: []iSecurityGroup{
//   				mySecurityGroup,
//   			},
//   		},
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
//   	SecurityGroups: []*iSecurityGroup{
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
type DockerServerOptions struct {
	// The type of compute to use for the docker server.
	//
	// See the `DockerServerComputeType` enum for the possible values.
	ComputeType DockerServerComputeType `field:"required" json:"computeType" yaml:"computeType"`
	// A list of maximum 5 security groups.
	// Default: - no security group.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

