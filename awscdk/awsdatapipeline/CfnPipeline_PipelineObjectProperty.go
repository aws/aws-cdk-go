package awsdatapipeline


// PipelineObject is property of the AWS::DataPipeline::Pipeline resource that contains information about a pipeline object.
//
// This can be a logical, physical, or physical attempt pipeline object. The complete set of components of a pipeline defines the pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineObjectProperty := &PipelineObjectProperty{
//   	Fields: []interface{}{
//   		&FieldProperty{
//   			Key: jsii.String("key"),
//
//   			// the properties below are optional
//   			RefValue: jsii.String("refValue"),
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   }
//
type CfnPipeline_PipelineObjectProperty struct {
	// Key-value pairs that define the properties of the object.
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// The ID of the object.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The name of the object.
	Name *string `field:"required" json:"name" yaml:"name"`
}

