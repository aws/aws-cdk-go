package previewawsdatapipelinemixins


// Properties for CfnPipelinePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPipelineMixinProps := &CfnPipelineMixinProps{
//   	Activate: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
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
//   					RefValue: jsii.String("refValue"),
//   					StringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	PipelineTags: []PipelineTagProperty{
//   		&PipelineTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html
//
type CfnPipelineMixinProps struct {
	// Indicates whether to validate and start the pipeline or stop an active pipeline.
	//
	// By default, the value is set to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-activate
	//
	Activate interface{} `field:"optional" json:"activate" yaml:"activate"`
	// A description of the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameter objects used with the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-parameterobjects
	//
	ParameterObjects interface{} `field:"optional" json:"parameterObjects" yaml:"parameterObjects"`
	// The parameter values used with the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-parametervalues
	//
	ParameterValues interface{} `field:"optional" json:"parameterValues" yaml:"parameterValues"`
	// The objects that define the pipeline.
	//
	// These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see [Editing Your Pipeline](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-manage-pipeline-modify-console.html) in the *AWS Data Pipeline Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-pipelineobjects
	//
	PipelineObjects interface{} `field:"optional" json:"pipelineObjects" yaml:"pipelineObjects"`
	// A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions.
	//
	// For more information, see [Controlling Access to Pipelines and Resources](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html) in the *AWS Data Pipeline Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datapipeline-pipeline.html#cfn-datapipeline-pipeline-pipelinetags
	//
	PipelineTags *[]*CfnPipelinePropsMixin_PipelineTagProperty `field:"optional" json:"pipelineTags" yaml:"pipelineTags"`
}

