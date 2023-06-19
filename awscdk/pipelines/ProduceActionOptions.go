package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/constructs-go/constructs/v3"
)

// Options for the `CodePipelineActionFactory.produce()` method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import constructs "github.com/aws/constructs-go/constructs"
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var artifactMap artifactMap
//   var bucket bucket
//   var buildImage iBuildImage
//   var buildSpec buildSpec
//   var codePipeline codePipeline
//   var construct construct
//   var duration duration
//   var policyStatement policyStatement
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var value interface{}
//   var vpc vpc
//
//   produceActionOptions := &ProduceActionOptions{
//   	ActionName: jsii.String("actionName"),
//   	Artifacts: artifactMap,
//   	Pipeline: codePipeline,
//   	RunOrder: jsii.Number(123),
//   	Scope: construct,
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
//   			EnvironmentVariables: map[string]buildEnvironmentVariable{
//   				"environmentVariablesKey": &buildEnvironmentVariable{
//   					"value": value,
//
//   					// the properties below are optional
//   					"type": awscdk.*Aws_codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   				},
//   			},
//   			Privileged: jsii.Boolean(false),
//   		},
//   		PartialBuildSpec: buildSpec,
//   		RolePolicy: []*policyStatement{
//   			policyStatement,
//   		},
//   		SecurityGroups: []iSecurityGroup{
//   			securityGroup,
//   		},
//   		SubnetSelection: &SubnetSelection{
//   			AvailabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			OnePerAz: jsii.Boolean(false),
//   			SubnetFilters: []*subnetFilter{
//   				subnetFilter,
//   			},
//   			SubnetGroupName: jsii.String("subnetGroupName"),
//   			SubnetName: jsii.String("subnetName"),
//   			Subnets: []iSubnet{
//   				subnet,
//   			},
//   			SubnetType: awscdk.Aws_ec2.SubnetType_ISOLATED,
//   		},
//   		Timeout: duration,
//   		Vpc: vpc,
//   	},
//   	FallbackArtifact: artifact,
//   	VariablesNamespace: jsii.String("variablesNamespace"),
//   }
//
// Experimental.
type ProduceActionOptions struct {
	// Name the action should get.
	// Experimental.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// Helper object to translate FileSets to CodePipeline Artifacts.
	// Experimental.
	Artifacts ArtifactMap `field:"required" json:"artifacts" yaml:"artifacts"`
	// The pipeline the action is being generated for.
	// Experimental.
	Pipeline CodePipeline `field:"required" json:"pipeline" yaml:"pipeline"`
	// RunOrder the action should get.
	// Experimental.
	RunOrder *float64 `field:"required" json:"runOrder" yaml:"runOrder"`
	// Scope in which to create constructs.
	// Experimental.
	Scope constructs.Construct `field:"required" json:"scope" yaml:"scope"`
	// Whether or not this action is inserted before self mutation.
	//
	// If it is, the action should take care to reflect some part of
	// its own definition in the pipeline action definition, to
	// trigger a restart after self-mutation (if necessary).
	// Experimental.
	BeforeSelfMutation *bool `field:"optional" json:"beforeSelfMutation" yaml:"beforeSelfMutation"`
	// If this action factory creates a CodeBuild step, default options to inherit.
	// Experimental.
	CodeBuildDefaults *CodeBuildOptions `field:"optional" json:"codeBuildDefaults" yaml:"codeBuildDefaults"`
	// An input artifact that CodeBuild projects that don't actually need an input artifact can use.
	//
	// CodeBuild Projects MUST have an input artifact in order to be added to the Pipeline. If
	// the Project doesn't actually care about its input (it can be anything), it can use the
	// Artifact passed here.
	// Experimental.
	FallbackArtifact awscodepipeline.Artifact `field:"optional" json:"fallbackArtifact" yaml:"fallbackArtifact"`
	// If this step is producing outputs, the variables namespace assigned to it.
	//
	// Pass this on to the Action you are creating.
	// Experimental.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
}

