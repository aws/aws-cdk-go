package awsdatapipeline


// Properties for defining a `CfnPipeline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipelineProps := &cfnPipelineProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	activate: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	parameterObjects: []interface{}{
//   		&parameterObjectProperty{
//   			attributes: []interface{}{
//   				&parameterAttributeProperty{
//   					key: jsii.String("key"),
//   					stringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			id: jsii.String("id"),
//   		},
//   	},
//   	parameterValues: []interface{}{
//   		&parameterValueProperty{
//   			id: jsii.String("id"),
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	pipelineObjects: []interface{}{
//   		&pipelineObjectProperty{
//   			fields: []interface{}{
//   				&fieldProperty{
//   					key: jsii.String("key"),
//
//   					// the properties below are optional
//   					refValue: jsii.String("refValue"),
//   					stringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			id: jsii.String("id"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	pipelineTags: []interface{}{
//   		&pipelineTagProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

