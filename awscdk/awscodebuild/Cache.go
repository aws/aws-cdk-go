package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Cache options for CodeBuild Project.
//
// A cache can store reusable pieces of your build environment and use them across multiple builds.
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
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-caching.html
//
type Cache interface {
}

// The jsii proxy struct for Cache
type jsiiProxy_Cache struct {
	_ byte // padding
}

func NewCache_Override(c Cache) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.Cache",
		nil, // no parameters
		c,
	)
}

// Create an S3 caching strategy.
func Cache_Bucket(bucket awss3.IBucket, options *BucketCacheOptions) Cache {
	_init_.Initialize()

	if err := validateCache_BucketParameters(bucket, options); err != nil {
		panic(err)
	}
	var returns Cache

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Cache",
		"bucket",
		[]interface{}{bucket, options},
		&returns,
	)

	return returns
}

// Create a local caching strategy.
func Cache_Local(modes ...LocalCacheMode) Cache {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range modes {
		args = append(args, a)
	}

	var returns Cache

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Cache",
		"local",
		args,
		&returns,
	)

	return returns
}

func Cache_None() Cache {
	_init_.Initialize()

	var returns Cache

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Cache",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

