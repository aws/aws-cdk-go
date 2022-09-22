package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Construction props for SimpleSynthAction.
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
//   var buildSpec buildSpec
//   var policyStatement policyStatement
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var value interface{}
//   var vpc vpc
//
//   simpleSynthActionProps := &simpleSynthActionProps{
//   	cloudAssemblyArtifact: artifact,
//   	sourceArtifact: artifact,
//   	synthCommand: jsii.String("synthCommand"),
//
//   	// the properties below are optional
//   	actionName: jsii.String("actionName"),
//   	additionalArtifacts: []additionalArtifact{
//   		&additionalArtifact{
//   			artifact: artifact,
//   			directory: jsii.String("directory"),
//   		},
//   	},
//   	buildCommand: jsii.String("buildCommand"),
//   	buildCommands: []*string{
//   		jsii.String("buildCommands"),
//   	},
//   	buildSpec: buildSpec,
//   	copyEnvironmentVariables: []*string{
//   		jsii.String("copyEnvironmentVariables"),
//   	},
//   	environment: &buildEnvironment{
//   		buildImage: buildImage,
//   		certificate: &buildEnvironmentCertificate{
//   			bucket: bucket,
//   			objectKey: jsii.String("objectKey"),
//   		},
//   		computeType: awscdk.Aws_codebuild.computeType_SMALL,
//   		environmentVariables: map[string]buildEnvironmentVariable{
//   			"environmentVariablesKey": &buildEnvironmentVariable{
//   				"value": value,
//
//   				// the properties below are optional
//   				"type": awscdk.*Aws_codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   			},
//   		},
//   		privileged: jsii.Boolean(false),
//   	},
//   	environmentVariables: map[string]*buildEnvironmentVariable{
//   		"environmentVariablesKey": &buildEnvironmentVariable{
//   			"value": value,
//
//   			// the properties below are optional
//   			"type": awscdk.*Aws_codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   		},
//   	},
//   	installCommand: jsii.String("installCommand"),
//   	installCommands: []*string{
//   		jsii.String("installCommands"),
//   	},
//   	projectName: jsii.String("projectName"),
//   	rolePolicyStatements: []*policyStatement{
//   		policyStatement,
//   	},
//   	subdirectory: jsii.String("subdirectory"),
//   	subnetSelection: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_ISOLATED,
//   	},
//   	testCommands: []*string{
//   		jsii.String("testCommands"),
//   	},
//   	vpc: vpc,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type SimpleSynthActionProps struct {
	// The artifact where the CloudAssembly should be emitted.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyArtifact awscodepipeline.Artifact `field:"required" json:"cloudAssemblyArtifact" yaml:"cloudAssemblyArtifact"`
	// The source artifact of the CodePipeline.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SourceArtifact awscodepipeline.Artifact `field:"required" json:"sourceArtifact" yaml:"sourceArtifact"`
	// Name of the build action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// Produce additional output artifacts after the build based on the given directories.
	//
	// Can be used to produce additional artifacts during the build step,
	// separate from the cloud assembly, which can be used further on in the
	// pipeline.
	//
	// Directories are evaluated with respect to `subdirectory`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AdditionalArtifacts *[]*AdditionalArtifact `field:"optional" json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// custom BuildSpec that is merged with the generated one.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildSpec awscodebuild.BuildSpec `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// Environment variables to copy over from parent env.
	//
	// These are environment variables that are being used by the build.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CopyEnvironmentVariables *[]*string `field:"optional" json:"copyEnvironmentVariables" yaml:"copyEnvironmentVariables"`
	// Build environment to use for CodeBuild job.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Environment *awscodebuild.BuildEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// Environment variables to send into build.
	//
	// NOTE: You may run into the 1000-character limit for the Action configuration if you have a large
	// number of variables or if their names or values are very long.
	// If you do, pass them to the underlying CodeBuild project directly in `environment` instead.
	// However, you will not be able to use CodePipeline Variables in this case.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Name of the CodeBuild project.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	RolePolicyStatements *[]awsiam.PolicyStatement `field:"optional" json:"rolePolicyStatements" yaml:"rolePolicyStatements"`
	// Directory inside the source where package.json and cdk.json are located.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The VPC where to execute the SimpleSynth.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The synth command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SynthCommand *string `field:"required" json:"synthCommand" yaml:"synthCommand"`
	// The build command.
	//
	// If your programming language requires a compilation step, put the
	// compilation command here.
	// Deprecated: Use `buildCommands` instead.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// The build commands.
	//
	// If your programming language requires a compilation step, put the
	// compilation command here.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildCommands *[]*string `field:"optional" json:"buildCommands" yaml:"buildCommands"`
	// The install command.
	//
	// If not provided by the build image or another dependency
	// management tool, at least install the CDK CLI here using
	// `npm install -g aws-cdk`.
	// Deprecated: Use `installCommands` instead.
	InstallCommand *string `field:"optional" json:"installCommand" yaml:"installCommand"`
	// Install commands.
	//
	// If not provided by the build image or another dependency
	// management tool, at least install the CDK CLI here using
	// `npm install -g aws-cdk`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	InstallCommands *[]*string `field:"optional" json:"installCommands" yaml:"installCommands"`
	// Test commands.
	//
	// These commands are run after the build commands but before the
	// synth command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	TestCommands *[]*string `field:"optional" json:"testCommands" yaml:"testCommands"`
}

