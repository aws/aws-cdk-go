package previewawsdatapipelinemixins


// PipelineObject is property of the AWS::DataPipeline::Pipeline resource that contains information about a pipeline object.
//
// This can be a logical, physical, or physical attempt pipeline object. The complete set of components of a pipeline defines the pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pipelineObjectProperty := &PipelineObjectProperty{
//   	Fields: []interface{}{
//   		&FieldProperty{
//   			Key: jsii.String("key"),
//   			RefValue: jsii.String("refValue"),
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobject.html
//
type CfnPipelinePropsMixin_PipelineObjectProperty struct {
	// Key-value pairs that define the properties of the object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobject.html#cfn-datapipeline-pipeline-pipelineobject-fields
	//
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// The ID of the object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobject.html#cfn-datapipeline-pipeline-pipelineobject-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobject.html#cfn-datapipeline-pipeline-pipelineobject-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

