package awsdatapipeline


// A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions.
//
// For more information, see [Controlling Access to Pipelines and Resources](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html) in the *AWS Data Pipeline Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineTagProperty := &pipelineTagProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnPipeline_PipelineTagProperty struct {
	// The key name of a tag.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value to associate with the key name.
	Value *string `field:"required" json:"value" yaml:"value"`
}

