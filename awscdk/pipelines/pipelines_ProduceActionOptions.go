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
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var artifact artifact
//   var artifactMap artifactMap
//   var bucket bucket
//   var buildImage iBuildImage
//   var buildSpec buildSpec
//   var cache cache
//   var codePipeline codePipeline
//   var construct construct
//   var policyStatement policyStatement
//   var securityGroup securityGroup
//   var stackOutputsMap stackOutputsMap
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var value interface{}
//   var vpc vpc
//
//   produceActionOptions := &produceActionOptions{
//   	actionName: jsii.String("actionName"),
//   	artifacts: artifactMap,
//   	pipeline: codePipeline,
//   	runOrder: jsii.Number(123),
//   	scope: construct,
//   	stackOutputsMap: stackOutputsMap,
//
//   	// the properties below are optional
//   	beforeSelfMutation: jsii.Boolean(false),
//   	codeBuildDefaults: &codeBuildOptions{
//   		buildEnvironment: &buildEnvironment{
//   			buildImage: buildImage,
//   			certificate: &buildEnvironmentCertificate{
//   				bucket: bucket,
//   				objectKey: jsii.String("objectKey"),
//   			},
//   			computeType: awscdk.Aws_codebuild.computeType_SMALL,
//   			environmentVariables: map[string]buildEnvironmentVariable{
//   				"environmentVariablesKey": &buildEnvironmentVariable{
//   					"value": value,
//
//   					// the properties below are optional
//   					"type": awscdk.*Aws_codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   				},
//   			},
//   			privileged: jsii.Boolean(false),
//   		},
//   		cache: cache,
//   		partialBuildSpec: buildSpec,
//   		rolePolicy: []*policyStatement{
//   			policyStatement,
//   		},
//   		securityGroups: []iSecurityGroup{
//   			securityGroup,
//   		},
//   		subnetSelection: &subnetSelection{
//   			availabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			onePerAz: jsii.Boolean(false),
//   			subnetFilters: []*subnetFilter{
//   				subnetFilter,
//   			},
//   			subnetGroupName: jsii.String("subnetGroupName"),
//   			subnets: []iSubnet{
//   				subnet,
//   			},
//   			subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//   		},
//   		timeout: cdk.duration.minutes(jsii.Number(30)),
//   		vpc: vpc,
//   	},
//   	fallbackArtifact: artifact,
//   	variablesNamespace: jsii.String("variablesNamespace"),
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
	BeforeSelfMutation *bool `field:"optional" json:"beforeSelfMutation" yaml:"beforeSelfMutation"`
	// If this action factory creates a CodeBuild step, default options to inherit.
	CodeBuildDefaults *CodeBuildOptions `field:"optional" json:"codeBuildDefaults" yaml:"codeBuildDefaults"`
	// An input artifact that CodeBuild projects that don't actually need an input artifact can use.
	//
	// CodeBuild Projects MUST have an input artifact in order to be added to the Pipeline. If
	// the Project doesn't actually care about its input (it can be anything), it can use the
	// Artifact passed here.
	FallbackArtifact awscodepipeline.Artifact `field:"optional" json:"fallbackArtifact" yaml:"fallbackArtifact"`
	// If this step is producing outputs, the variables namespace assigned to it.
	//
	// Pass this on to the Action you are creating.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
}

