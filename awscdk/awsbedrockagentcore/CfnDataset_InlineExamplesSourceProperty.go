package awsbedrockagentcore


// Inline examples provided directly in the request body.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var examples interface{}
//
//   inlineExamplesSourceProperty := &InlineExamplesSourceProperty{
//   	Examples: []interface{}{
//   		examples,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-dataset-inlineexamplessource.html
//
type CfnDataset_InlineExamplesSourceProperty struct {
	// Examples to add.
	//
	// Each example is a free-form JSON document validated against the declared schemaType.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-dataset-inlineexamplessource.html#cfn-bedrockagentcore-dataset-inlineexamplessource-examples
	//
	Examples interface{} `field:"required" json:"examples" yaml:"examples"`
}

