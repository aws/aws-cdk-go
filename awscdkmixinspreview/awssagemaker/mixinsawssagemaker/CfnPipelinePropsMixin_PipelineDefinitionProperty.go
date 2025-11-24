package mixinsawssagemaker


// The definition of the pipeline.
//
// This can be either a JSON string or an Amazon S3 location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pipelineDefinitionProperty := &PipelineDefinitionProperty{
//   	PipelineDefinitionBody: jsii.String("pipelineDefinitionBody"),
//   	PipelineDefinitionS3Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		ETag: jsii.String("eTag"),
//   		Key: jsii.String("key"),
//   		Version: jsii.String("version"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-pipelinedefinition.html
//
type CfnPipelinePropsMixin_PipelineDefinitionProperty struct {
	// The [JSON pipeline definition](https://docs.aws.amazon.com/https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-pipelinedefinition.html#cfn-sagemaker-pipeline-pipelinedefinition-pipelinedefinitionbody
	//
	PipelineDefinitionBody *string `field:"optional" json:"pipelineDefinitionBody" yaml:"pipelineDefinitionBody"`
	// The location of the pipeline definition stored in Amazon S3.
	//
	// If specified, SageMaker retrieves the pipeline definition from this location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-pipelinedefinition.html#cfn-sagemaker-pipeline-pipelinedefinition-pipelinedefinitions3location
	//
	PipelineDefinitionS3Location interface{} `field:"optional" json:"pipelineDefinitionS3Location" yaml:"pipelineDefinitionS3Location"`
}

