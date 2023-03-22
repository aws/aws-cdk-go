package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineDefinitionProperty := &PipelineDefinitionProperty{
//   	PipelineDefinitionBody: jsii.String("pipelineDefinitionBody"),
//   	PipelineDefinitionS3Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//
//   		// the properties below are optional
//   		ETag: jsii.String("eTag"),
//   		Version: jsii.String("version"),
//   	},
//   }
//
type CfnPipeline_PipelineDefinitionProperty struct {
	// `CfnPipeline.PipelineDefinitionProperty.PipelineDefinitionBody`.
	PipelineDefinitionBody *string `field:"optional" json:"pipelineDefinitionBody" yaml:"pipelineDefinitionBody"`
	// `CfnPipeline.PipelineDefinitionProperty.PipelineDefinitionS3Location`.
	PipelineDefinitionS3Location interface{} `field:"optional" json:"pipelineDefinitionS3Location" yaml:"pipelineDefinitionS3Location"`
}

