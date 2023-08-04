package awscodepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var resource resource
//   var role role
//
//   actionProperties := &ActionProperties{
//   	ActionName: jsii.String("actionName"),
//   	ArtifactBounds: &ActionArtifactBounds{
//   		MaxInputs: jsii.Number(123),
//   		MaxOutputs: jsii.Number(123),
//   		MinInputs: jsii.Number(123),
//   		MinOutputs: jsii.Number(123),
//   	},
//   	Category: awscdk.Aws_codepipeline.ActionCategory_SOURCE,
//   	Provider: jsii.String("provider"),
//
//   	// the properties below are optional
//   	Account: jsii.String("account"),
//   	Inputs: []*artifact{
//   		artifact,
//   	},
//   	Outputs: []*artifact{
//   		artifact,
//   	},
//   	Owner: jsii.String("owner"),
//   	Region: jsii.String("region"),
//   	Resource: resource,
//   	Role: role,
//   	RunOrder: jsii.Number(123),
//   	VariablesNamespace: jsii.String("variablesNamespace"),
//   	Version: jsii.String("version"),
//   }
//
type ActionProperties struct {
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	ArtifactBounds *ActionArtifactBounds `field:"required" json:"artifactBounds" yaml:"artifactBounds"`
	// The category of the action.
	//
	// The category defines which action type the owner
	// (the entity that performs the action) performs.
	Category ActionCategory `field:"required" json:"category" yaml:"category"`
	// The service provider that the action calls.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// The account the Action is supposed to live in.
	//
	// For Actions backed by resources,
	// this is inferred from the Stack `resource` is part of.
	// However, some Actions, like the CloudFormation ones,
	// are not backed by any resource, and they still might want to be cross-account.
	// In general, a concrete Action class should specify either `resource`,
	// or `account` - but not both.
	Account *string `field:"optional" json:"account" yaml:"account"`
	Inputs *[]Artifact `field:"optional" json:"inputs" yaml:"inputs"`
	Outputs *[]Artifact `field:"optional" json:"outputs" yaml:"outputs"`
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// The AWS region the given Action resides in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the `PipelineProps#crossRegionReplicationBuckets` property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	// Default: the Action resides in the same region as the Pipeline.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The optional resource that is backing this Action.
	//
	// This is used for automatically handling Actions backed by
	// resources from a different account and/or region.
	Resource awscdk.IResource `field:"optional" json:"resource" yaml:"resource"`
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The order in which AWS CodePipeline runs this action. For more information, see the AWS CodePipeline User Guide.
	//
	// https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Default: - a name will be generated, based on the stage and action names.
	//
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	Version *string `field:"optional" json:"version" yaml:"version"`
}

