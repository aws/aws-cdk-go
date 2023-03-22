package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPipeline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parallelismConfiguration interface{}
//   var pipelineDefinition interface{}
//
//   cfnPipelineProps := &cfnPipelineProps{
//   	pipelineDefinition: pipelineDefinition,
//   	pipelineName: jsii.String("pipelineName"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	parallelismConfiguration: parallelismConfiguration,
//   	pipelineDescription: jsii.String("pipelineDescription"),
//   	pipelineDisplayName: jsii.String("pipelineDisplayName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPipelineProps struct {
	// The definition of the pipeline.
	//
	// This can be either a JSON string or an Amazon S3 location.
	PipelineDefinition interface{} `field:"required" json:"pipelineDefinition" yaml:"pipelineDefinition"`
	// The name of the pipeline.
	PipelineName *string `field:"required" json:"pipelineName" yaml:"pipelineName"`
	// The Amazon Resource Name (ARN) of the IAM role used to execute the pipeline.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// `AWS::SageMaker::Pipeline.ParallelismConfiguration`.
	ParallelismConfiguration interface{} `field:"optional" json:"parallelismConfiguration" yaml:"parallelismConfiguration"`
	// The description of the pipeline.
	PipelineDescription *string `field:"optional" json:"pipelineDescription" yaml:"pipelineDescription"`
	// The display name of the pipeline.
	PipelineDisplayName *string `field:"optional" json:"pipelineDisplayName" yaml:"pipelineDisplayName"`
	// The tags of the pipeline.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

