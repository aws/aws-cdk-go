package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineDefinitionProperty := &pipelineDefinitionProperty{
//   	pipelineDefinitionBody: jsii.String("pipelineDefinitionBody"),
//   	pipelineDefinitionS3Location: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//
//   		// the properties below are optional
//   		eTag: jsii.String("eTag"),
//   		version: jsii.String("version"),
//   	},
//   }
//
type CfnPipeline_PipelineDefinitionProperty struct {
	// `CfnPipeline.PipelineDefinitionProperty.PipelineDefinitionBody`.
	PipelineDefinitionBody *string `field:"optional" json:"pipelineDefinitionBody" yaml:"pipelineDefinitionBody"`
	// `CfnPipeline.PipelineDefinitionProperty.PipelineDefinitionS3Location`.
	PipelineDefinitionS3Location interface{} `field:"optional" json:"pipelineDefinitionS3Location" yaml:"pipelineDefinitionS3Location"`
}

