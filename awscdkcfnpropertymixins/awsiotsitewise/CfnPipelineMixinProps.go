package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPipelinePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnPipelineMixinProps := &CfnPipelineMixinProps{
//   	Computations: []interface{}{
//   		&ComputeNodeProperty{
//   			ComputeNodeName: jsii.String("computeNodeName"),
//   			DependsOn: []*string{
//   				jsii.String("dependsOn"),
//   			},
//   			EnvironmentVariables: map[string]*string{
//   				"environmentVariablesKey": jsii.String("environmentVariables"),
//   			},
//   			TaskName: jsii.String("taskName"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	PipelineName: jsii.String("pipelineName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkspaceName: jsii.String("workspaceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html
//
type CfnPipelineMixinProps struct {
	// The list of compute nodes that form the pipeline DAG.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html#cfn-iotsitewise-pipeline-computations
	//
	Computations interface{} `field:"optional" json:"computations" yaml:"computations"`
	// A description of the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html#cfn-iotsitewise-pipeline-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map of environment variable key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html#cfn-iotsitewise-pipeline-environmentvariables
	//
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The name of the pipeline.
	//
	// Must be unique within the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html#cfn-iotsitewise-pipeline-pipelinename
	//
	PipelineName *string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html#cfn-iotsitewise-pipeline-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html#cfn-iotsitewise-pipeline-workspacename
	//
	WorkspaceName *string `field:"optional" json:"workspaceName" yaml:"workspaceName"`
}

