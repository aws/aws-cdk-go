package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Options for customizing a single CodeBuild project.
//
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
type CodeBuildOptions struct {
	// Partial build environment, will be combined with other build environments that apply.
	// Experimental.
	BuildEnvironment *awscodebuild.BuildEnvironment `field:"optional" json:"buildEnvironment" yaml:"buildEnvironment"`
	// Partial buildspec, will be combined with other buildspecs that apply.
	//
	// The BuildSpec must be available inline--it cannot reference a file
	// on disk.
	// Experimental.
	PartialBuildSpec awscodebuild.BuildSpec `field:"optional" json:"partialBuildSpec" yaml:"partialBuildSpec"`
	// Policy statements to add to role.
	// Experimental.
	RolePolicy *[]awsiam.PolicyStatement `field:"optional" json:"rolePolicy" yaml:"rolePolicy"`
	// Which security group(s) to associate with the project network interfaces.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The VPC where to create the CodeBuild network interfaces in.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

