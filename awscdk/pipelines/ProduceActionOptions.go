package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/constructs-go/constructs/v10"
)

// Options for the `CodePipelineActionFactory.produce()` method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var artifact Artifact
//   var artifactMap ArtifactMap
//   var bucket Bucket
//   var buildImage IBuildImage
//   var buildSpec BuildSpec
//   var cache Cache
//   var codePipeline CodePipeline
//   var construct Construct
//   var fileSystemLocation IFileSystemLocation
//   var fleet Fleet
//   var logGroup LogGroup
//   var policyStatement PolicyStatement
//   var securityGroup SecurityGroup
//   var stackOutputsMap StackOutputsMap
//   var subnet Subnet
//   var subnetFilter SubnetFilter
//   var value interface{}
//   var vpc Vpc
//
//   produceActionOptions := &ProduceActionOptions{
//   	ActionName: jsii.String("actionName"),
//   	Artifacts: artifactMap,
//   	Pipeline: codePipeline,
//   	RunOrder: jsii.Number(123),
//   	Scope: construct,
//   	StackOutputsMap: stackOutputsMap,
//
//   	// the properties below are optional
//   	BeforeSelfMutation: jsii.Boolean(false),
//   	CodeBuildDefaults: &CodeBuildOptions{
//   		BuildEnvironment: &BuildEnvironment{
//   			BuildImage: buildImage,
//   			Certificate: &BuildEnvironmentCertificate{
//   				Bucket: bucket,
//   				ObjectKey: jsii.String("objectKey"),
//   			},
//   			ComputeType: awscdk.Aws_codebuild.ComputeType_SMALL,
//   			DockerServer: &DockerServerOptions{
//   				ComputeType: awscdk.*Aws_codebuild.DockerServerComputeType_SMALL,
//
//   				// the properties below are optional
//   				SecurityGroups: []ISecurityGroup{
//   					securityGroup,
//   				},
//   			},
//   			EnvironmentVariables: map[string]BuildEnvironmentVariable{
//   				"environmentVariablesKey": &BuildEnvironmentVariable{
//   					"value": value,
//
//   					// the properties below are optional
//   					"type": awscdk.*Aws_codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   				},
//   			},
//   			Fleet: fleet,
//   			Privileged: jsii.Boolean(false),
//   		},
//   		Cache: cache,
//   		FileSystemLocations: []IFileSystemLocation{
//   			fileSystemLocation,
//   		},
//   		Logging: &LoggingOptions{
//   			CloudWatch: &CloudWatchLoggingOptions{
//   				Enabled: jsii.Boolean(false),
//   				LogGroup: logGroup,
//   				Prefix: jsii.String("prefix"),
//   			},
//   			S3: &S3LoggingOptions{
//   				Bucket: bucket,
//
//   				// the properties below are optional
//   				Enabled: jsii.Boolean(false),
//   				Encrypted: jsii.Boolean(false),
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   		PartialBuildSpec: buildSpec,
//   		RolePolicy: []PolicyStatement{
//   			policyStatement,
//   		},
//   		SecurityGroups: []ISecurityGroup{
//   			securityGroup,
//   		},
//   		SubnetSelection: &SubnetSelection{
//   			AvailabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			OnePerAz: jsii.Boolean(false),
//   			SubnetFilters: []SubnetFilter{
//   				subnetFilter,
//   			},
//   			SubnetGroupName: jsii.String("subnetGroupName"),
//   			Subnets: []ISubnet{
//   				subnet,
//   			},
//   			SubnetType: awscdk.Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   		},
//   		Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   		Vpc: vpc,
//   	},
//   	FallbackArtifact: artifact,
//   	VariablesNamespace: jsii.String("variablesNamespace"),
//   }
//
type ProduceActionOptions struct {
	// Name the action should get.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// Helper object to translate FileSets to CodePipeline Artifacts.
	Artifacts ArtifactMap `field:"required" json:"artifacts" yaml:"artifacts"`
	// The pipeline the action is being generated for.
	Pipeline CodePipeline `field:"required" json:"pipeline" yaml:"pipeline"`
	// RunOrder the action should get.
	RunOrder *float64 `field:"required" json:"runOrder" yaml:"runOrder"`
	// Scope in which to create constructs.
	Scope constructs.Construct `field:"required" json:"scope" yaml:"scope"`
	// Helper object to produce variables exported from stack deployments.
	//
	// If your step references outputs from a stack deployment, use
	// this to map the output references to Codepipeline variable names.
	//
	// Note - Codepipeline variables can only be referenced in action
	// configurations.
	StackOutputsMap StackOutputsMap `field:"required" json:"stackOutputsMap" yaml:"stackOutputsMap"`
	// Whether or not this action is inserted before self mutation.
	//
	// If it is, the action should take care to reflect some part of
	// its own definition in the pipeline action definition, to
	// trigger a restart after self-mutation (if necessary).
	// Default: false.
	//
	BeforeSelfMutation *bool `field:"optional" json:"beforeSelfMutation" yaml:"beforeSelfMutation"`
	// If this action factory creates a CodeBuild step, default options to inherit.
	// Default: - No CodeBuild project defaults.
	//
	CodeBuildDefaults *CodeBuildOptions `field:"optional" json:"codeBuildDefaults" yaml:"codeBuildDefaults"`
	// An input artifact that CodeBuild projects that don't actually need an input artifact can use.
	//
	// CodeBuild Projects MUST have an input artifact in order to be added to the Pipeline. If
	// the Project doesn't actually care about its input (it can be anything), it can use the
	// Artifact passed here.
	// Default: - A fallback artifact does not exist.
	//
	FallbackArtifact awscodepipeline.Artifact `field:"optional" json:"fallbackArtifact" yaml:"fallbackArtifact"`
	// If this step is producing outputs, the variables namespace assigned to it.
	//
	// Pass this on to the Action you are creating.
	// Default: - Step doesn't produce any outputs.
	//
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
}

