package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPipeline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipelineProps := &CfnPipelineProps{
//   	Computations: []interface{}{
//   		&ComputeNodeProperty{
//   			ComputeNodeName: jsii.String("computeNodeName"),
//   			TaskName: jsii.String("taskName"),
//
//   			// the properties below are optional
//   			DependsOn: []*string{
//   				jsii.String("dependsOn"),
//   			},
//   			EnvironmentVariables: map[string]*string{
//   				"environmentVariablesKey": jsii.String("environmentVariables"),
//   			},
//   		},
//   	},
//   	PipelineName: jsii.String("pipelineName"),
//   	WorkspaceName: jsii.String("workspaceName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html
//
type CfnPipelineProps struct {
	// The list of compute nodes that form the pipeline DAG.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html#cfn-iotsitewise-pipeline-computations
	//
	Computations interface{} `field:"required" json:"computations" yaml:"computations"`
	// The name of the pipeline.
	//
	// Must be unique within the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html#cfn-iotsitewise-pipeline-pipelinename
	//
	PipelineName *string `field:"required" json:"pipelineName" yaml:"pipelineName"`
	// The name of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html#cfn-iotsitewise-pipeline-workspacename
	//
	WorkspaceName *string `field:"required" json:"workspaceName" yaml:"workspaceName"`
	// A description of the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html#cfn-iotsitewise-pipeline-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map of environment variable key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html#cfn-iotsitewise-pipeline-environmentvariables
	//
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html#cfn-iotsitewise-pipeline-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

