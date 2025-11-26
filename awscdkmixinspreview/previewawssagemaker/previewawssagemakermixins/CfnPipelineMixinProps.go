package previewawssagemakermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPipelinePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var parallelismConfiguration interface{}
//   var pipelineDefinition interface{}
//
//   cfnPipelineMixinProps := &CfnPipelineMixinProps{
//   	ParallelismConfiguration: parallelismConfiguration,
//   	PipelineDefinition: pipelineDefinition,
//   	PipelineDescription: jsii.String("pipelineDescription"),
//   	PipelineDisplayName: jsii.String("pipelineDisplayName"),
//   	PipelineName: jsii.String("pipelineName"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-pipeline.html
//
type CfnPipelineMixinProps struct {
	// The parallelism configuration applied to the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-pipeline.html#cfn-sagemaker-pipeline-parallelismconfiguration
	//
	ParallelismConfiguration interface{} `field:"optional" json:"parallelismConfiguration" yaml:"parallelismConfiguration"`
	// The definition of the pipeline.
	//
	// This can be either a JSON string or an Amazon S3 location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-pipeline.html#cfn-sagemaker-pipeline-pipelinedefinition
	//
	PipelineDefinition interface{} `field:"optional" json:"pipelineDefinition" yaml:"pipelineDefinition"`
	// The description of the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-pipeline.html#cfn-sagemaker-pipeline-pipelinedescription
	//
	PipelineDescription *string `field:"optional" json:"pipelineDescription" yaml:"pipelineDescription"`
	// The display name of the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-pipeline.html#cfn-sagemaker-pipeline-pipelinedisplayname
	//
	PipelineDisplayName *string `field:"optional" json:"pipelineDisplayName" yaml:"pipelineDisplayName"`
	// The name of the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-pipeline.html#cfn-sagemaker-pipeline-pipelinename
	//
	PipelineName *string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
	// The Amazon Resource Name (ARN) of the IAM role used to execute the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-pipeline.html#cfn-sagemaker-pipeline-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The tags of the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-pipeline.html#cfn-sagemaker-pipeline-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

