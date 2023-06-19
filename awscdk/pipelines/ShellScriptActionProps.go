package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties for ShellScriptAction.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var bucket bucket
//   var buildImage iBuildImage
//   var policyStatement policyStatement
//   var securityGroup securityGroup
//   var stackOutput stackOutput
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var value interface{}
//   var vpc vpc
//
//   shellScriptActionProps := &ShellScriptActionProps{
//   	ActionName: jsii.String("actionName"),
//   	Commands: []*string{
//   		jsii.String("commands"),
//   	},
//
//   	// the properties below are optional
//   	AdditionalArtifacts: []*artifact{
//   		artifact,
//   	},
//   	BashOptions: jsii.String("bashOptions"),
//   	Environment: &BuildEnvironment{
//   		BuildImage: buildImage,
//   		Certificate: &BuildEnvironmentCertificate{
//   			Bucket: bucket,
//   			ObjectKey: jsii.String("objectKey"),
//   		},
//   		ComputeType: awscdk.Aws_codebuild.ComputeType_SMALL,
//   		EnvironmentVariables: map[string]buildEnvironmentVariable{
//   			"environmentVariablesKey": &buildEnvironmentVariable{
//   				"value": value,
//
//   				// the properties below are optional
//   				"type": awscdk.*Aws_codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   			},
//   		},
//   		Privileged: jsii.Boolean(false),
//   	},
//   	EnvironmentVariables: map[string]*buildEnvironmentVariable{
//   		"environmentVariablesKey": &buildEnvironmentVariable{
//   			"value": value,
//
//   			// the properties below are optional
//   			"type": awscdk.*Aws_codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   		},
//   	},
//   	RolePolicyStatements: []*policyStatement{
//   		policyStatement,
//   	},
//   	RunOrder: jsii.Number(123),
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	SubnetSelection: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		OnePerAz: jsii.Boolean(false),
//   		SubnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		SubnetGroupName: jsii.String("subnetGroupName"),
//   		SubnetName: jsii.String("subnetName"),
//   		Subnets: []iSubnet{
//   			subnet,
//   		},
//   		SubnetType: awscdk.Aws_ec2.SubnetType_ISOLATED,
//   	},
//   	UseOutputs: map[string]*stackOutput{
//   		"useOutputsKey": stackOutput,
//   	},
//   	Vpc: vpc,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type ShellScriptActionProps struct {
	// Name of the validation action in the pipeline.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// Commands to run.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Commands *[]*string `field:"required" json:"commands" yaml:"commands"`
	// Additional artifacts to use as input for the CodeBuild project.
	//
	// You can use these files to load more complex test sets into the
	// shellscript build environment.
	//
	// The files artifact given here will be unpacked into the current
	// working directory, the other ones will be unpacked into directories
	// which are available through the environment variables
	// $CODEBUILD_SRC_DIR_<artifactName>.
	//
	// The CodeBuild job must have at least one input artifact, so you
	// must provide either at least one additional artifact here or one
	// stack output using `useOutput`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AdditionalArtifacts *[]awscodepipeline.Artifact `field:"optional" json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// Bash options to set at the start of the script.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BashOptions *string `field:"optional" json:"bashOptions" yaml:"bashOptions"`
	// The CodeBuild environment where scripts are executed.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Environment *awscodebuild.BuildEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// Environment variables to send into build.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Additional policy statements to add to the execution role.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	RolePolicyStatements *[]awsiam.PolicyStatement `field:"optional" json:"rolePolicyStatements" yaml:"rolePolicyStatements"`
	// RunOrder for this action.
	//
	// Use this to sequence the shell script after the deployments.
	//
	// The default value is 100 so you don't have to supply the value if you just
	// want to run this after the application stacks have been deployed, and you
	// don't have more than 100 stacks.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// Which security group to associate with the script's project network interfaces.
	//
	// If no security group is identified, one will be created automatically.
	//
	// Only used if 'vpc' is supplied.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// Stack outputs to make available as environment variables.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	UseOutputs *map[string]StackOutput `field:"optional" json:"useOutputs" yaml:"useOutputs"`
	// The VPC where to execute the specified script.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

