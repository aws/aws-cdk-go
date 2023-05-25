package awsdatapipeline


// Properties for defining a `CfnPipeline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipelineProps := &CfnPipelineProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Activate: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	ParameterObjects: []interface{}{
//   		&ParameterObjectProperty{
//   			Attributes: []interface{}{
//   				&ParameterAttributeProperty{
//   					Key: jsii.String("key"),
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	ParameterValues: []interface{}{
//   		&ParameterValueProperty{
//   			Id: jsii.String("id"),
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	PipelineObjects: []interface{}{
//   		&PipelineObjectProperty{
//   			Fields: []interface{}{
//   				&FieldProperty{
//   					Key: jsii.String("key"),
//
//   					// the properties below are optional
//   					RefValue: jsii.String("refValue"),
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	PipelineTags: []interface{}{
//   		&PipelineTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPipelineProps struct {
	// The name of the pipeline.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicates whether to validate and start the pipeline or stop an active pipeline.
	//
	// By default, the value is set to `true` .
	Activate interface{} `field:"optional" json:"activate" yaml:"activate"`
	// A description of the pipeline.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The parameter objects used with the pipeline.
	ParameterObjects interface{} `field:"optional" json:"parameterObjects" yaml:"parameterObjects"`
	// The parameter values used with the pipeline.
	ParameterValues interface{} `field:"optional" json:"parameterValues" yaml:"parameterValues"`
	// The objects that define the pipeline.
	//
	// These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see [Editing Your Pipeline](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-manage-pipeline-modify-console.html) in the *AWS Data Pipeline Developer Guide* .
	PipelineObjects interface{} `field:"optional" json:"pipelineObjects" yaml:"pipelineObjects"`
	// A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions.
	//
	// For more information, see [Controlling Access to Pipelines and Resources](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html) in the *AWS Data Pipeline Developer Guide* .
	PipelineTags interface{} `field:"optional" json:"pipelineTags" yaml:"pipelineTags"`
}

